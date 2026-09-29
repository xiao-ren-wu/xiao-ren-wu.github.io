// Command siteserver serves the built static site that lives in the parent
// directory, so the generated HTML can be previewed locally before it is pushed
// to GitHub Pages.
//
// The site is an Astro + Starlight build, and three of its characteristics are
// not handled correctly by the standard library's file server:
//
//   - Pagefind's search index ships a WebAssembly module under the extension
//     ".pagefind". pagefind.js loads it with WebAssembly.instantiateStreaming,
//     which rejects any Content-Type other than application/wasm and falls back
//     to a noticeably slower code path, so the type has to be pinned.
//   - Every knowledge-base URL is percent-encoded UTF-8 (for example
//     /knowledge/%E7%AE%97%E6%B3%95/...), so the on-disk directory name is only
//     recoverable from the decoded path.
//   - .git lives inside the published root, so dotfile requests have to be
//     refused instead of served.
package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// contentTypes pins the Content-Type of every extension the site ships, so the
// answer never depends on the host's /etc/mime.types.
//
// The .pagefind entry is load-bearing; the .pf_* entries are opaque compressed
// blobs that pagefind decodes through arrayBuffer(), for which octet-stream is
// both correct and harmless.
var contentTypes = map[string]string{
	".html":        "text/html; charset=utf-8",
	".htm":         "text/html; charset=utf-8",
	".css":         "text/css; charset=utf-8",
	".js":          "text/javascript; charset=utf-8",
	".mjs":         "text/javascript; charset=utf-8",
	".json":        "application/json; charset=utf-8",
	".xml":         "application/xml; charset=utf-8",
	".txt":         "text/plain; charset=utf-8",
	".md":          "text/markdown; charset=utf-8",
	".svg":         "image/svg+xml",
	".wasm":        "application/wasm",
	".pagefind":    "application/wasm",
	".pf_index":    "application/octet-stream",
	".pf_fragment": "application/octet-stream",
	".pf_meta":     "application/octet-stream",
	".woff2":       "font/woff2",
	".woff":        "font/woff",
	".ttf":         "font/ttf",
	".otf":         "font/otf",
	".ico":         "image/x-icon",
	".png":         "image/png",
	".jpg":         "image/jpeg",
	".jpeg":        "image/jpeg",
	".gif":         "image/gif",
	".webp":        "image/webp",
	".avif":        "image/avif",
}

// site resolves request URLs to files below root.
type site struct {
	root        string
	notFound    []byte
	notFoundMod time.Time
}

func newSite(root string) (*site, error) {
	s := &site{root: root}

	body, err := os.ReadFile(filepath.Join(root, "404.html"))
	switch {
	case err == nil:
		s.notFound = body
		if info, err := os.Stat(filepath.Join(root, "404.html")); err == nil {
			s.notFoundMod = info.ModTime()
		}
	case errors.Is(err, fs.ErrNotExist):
		log.Printf("warning: no 404.html in %s, falling back to a plain-text error", root)
		s.notFound = []byte("404 not found\n")
	default:
		return nil, fmt.Errorf("read 404.html: %w", err)
	}
	return s, nil
}

// resolve maps a request path to a file under the site root. It returns a
// redirect target when the URL should be canonicalised (a directory requested
// without its trailing slash), the absolute file path to serve, and whether
// either was found.
func (s *site) resolve(urlPath string) (file, redirect string, ok bool) {
	// net/http has already percent-decoded the URL into URL.Path, which is
	// what matches the on-disk UTF-8 names.
	clean := path.Clean("/" + strings.TrimPrefix(urlPath, "/"))

	for _, seg := range strings.Split(clean, "/") {
		// Refuse dotfiles and dot-directories. .git sits inside the published
		// root, so serving it would hand out the entire repository history.
		if seg != "" && strings.HasPrefix(seg, ".") {
			return "", "", false
		}
	}

	fsPath := filepath.Join(s.root, filepath.FromSlash(clean))
	info, err := os.Stat(fsPath)
	switch {
	case err == nil && info.IsDir():
		if clean == "/" || strings.HasSuffix(urlPath, "/") {
			index := filepath.Join(fsPath, "index.html")
			if st, err := os.Stat(index); err == nil && st.Mode().IsRegular() {
				return index, "", true
			}
			return "", "", false
		}
		return "", clean + "/", true

	case err == nil && info.Mode().IsRegular():
		return fsPath, "", true
	}

	// A directory requested without its trailing slash and without the
	// directory existing yet, e.g. /knowledge/算法/二叉树/二叉树的右视图.
	if !strings.HasSuffix(clean, "/") && path.Ext(clean) == "" {
		index := filepath.Join(fsPath, "index.html")
		if st, err := os.Stat(index); err == nil && st.Mode().IsRegular() {
			return "", clean + "/", true
		}
	}
	return "", "", false
}

func (s *site) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, redirect, ok := s.resolve(r.URL.Path)
	if !ok {
		s.serveNotFound(w, r)
		return
	}
	if redirect != "" {
		u := *r.URL
		u.Path = redirect
		// Clearing RawPath makes net/http re-encode the UTF-8 path itself.
		u.RawPath = ""
		http.Redirect(w, r, u.RequestURI(), http.StatusMovedPermanently)
		return
	}

	f, err := os.Open(file)
	if err != nil {
		s.serveNotFound(w, r)
		return
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil || !info.Mode().IsRegular() {
		s.serveNotFound(w, r)
		return
	}

	h := w.Header()
	h.Set("Content-Type", contentType(file))
	h.Set("Cache-Control", cacheControl(r.URL.Path))
	// Setting the ETag up front lets http.ServeContent's precondition check
	// answer If-None-Match with 304, which it does not do on its own.
	h.Set("Etag", etag(info))
	http.ServeContent(w, r, filepath.Base(file), info.ModTime(), f)
}

// serveNotFound replies with the site's own 404 page and a real 404 status.
func (s *site) serveNotFound(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")
	h.Set("Cache-Control", "no-cache")
	h.Set("Content-Length", strconv.Itoa(len(s.notFound)))
	w.WriteHeader(http.StatusNotFound)
	if r.Method == http.MethodHead {
		return
	}
	io.Copy(w, bytes.NewReader(s.notFound))
}

// cacheControl marks the content-hashed bundles under /_astro/ immutable and
// everything else must-revalidate, so a local rebuild shows up on reload.
func cacheControl(urlPath string) string {
	if strings.HasPrefix(urlPath, "/_astro/") {
		return "public, max-age=31536000, immutable"
	}
	return "no-cache"
}

func contentType(file string) string {
	ext := strings.ToLower(filepath.Ext(file))
	if ct, ok := contentTypes[ext]; ok {
		return ct
	}
	if ct := mime.TypeByExtension(ext); ct != "" {
		return ct
	}
	return "application/octet-stream"
}

func etag(info fs.FileInfo) string {
	return fmt.Sprintf(`"%x-%x"`, info.ModTime().UnixNano(), info.Size())
}

// statusRecorder remembers the status code for the access log.
type statusRecorder struct {
	http.ResponseWriter
	code int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.code = code
	r.ResponseWriter.WriteHeader(code)
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, code: http.StatusOK}
		next.ServeHTTP(rec, r)
		log.Printf("%s %s %d %s", r.Method, r.URL.RequestURI(), rec.code, time.Since(start).Round(time.Microsecond))
	})
}

// looksLikeSite reports whether dir holds a build of this project, as opposed
// to some other directory that merely happens to contain an index.html.
func looksLikeSite(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, "index.html")); err != nil {
		return false
	}
	for _, marker := range []string{"404.html", "sitemap-index.xml", "pagefind", "_astro"} {
		if _, err := os.Stat(filepath.Join(dir, marker)); err == nil {
			return true
		}
	}
	return false
}

// resolveRoot determines which directory to serve.
//
// An empty flag value means auto-detect, and detection intentionally does not
// lean on the working directory: resolving ".." against it silently served a
// completely unrelated static site when the binary was run from the repository
// root instead of from server/. The executable's own location is tried first so
// the answer stays the same wherever the binary is invoked from.
func resolveRoot(flagValue string) (string, error) {
	if flagValue != "" {
		return filepath.Abs(flagValue)
	}

	var candidates []string
	if exe, err := os.Executable(); err == nil {
		dir := filepath.Dir(exe)
		// The binary normally lives in server/, so the site is one level up.
		candidates = append(candidates, filepath.Dir(dir), dir)
	}
	if wd, err := os.Getwd(); err == nil {
		candidates = append(candidates, wd)
		// Walk up a few levels to cover `go run .` from inside server/ without
		// ever reaching far enough to latch onto an unrelated site.
		for dir, i := filepath.Dir(wd), 0; i < 3 && dir != filepath.Dir(dir); i++ {
			candidates = append(candidates, dir)
			dir = filepath.Dir(dir)
		}
	}

	for _, dir := range candidates {
		abs, err := filepath.Abs(dir)
		if err != nil {
			continue
		}
		if looksLikeSite(abs) {
			return abs, nil
		}
	}
	if len(candidates) == 0 {
		return "", errors.New("cannot locate the site root; pass -root <dir>")
	}
	return "", fmt.Errorf("no built site found near %s; pass -root <dir>", candidates[0])
}

func main() {
	addr := flag.String("addr", "127.0.0.1:4321", "address to listen on")
	root := flag.String("root", "", "directory holding the built site (default: auto-detect)")
	flag.Parse()

	absRoot, err := resolveRoot(*root)
	if err != nil {
		log.Fatal(err)
	}
	if !looksLikeSite(absRoot) {
		log.Printf("warning: %s does not look like a build of this site (no 404.html, sitemap-index.xml, pagefind/ or _astro/)", absRoot)
	}
	s, err := newSite(absRoot)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:              *addr,
		Handler:           logRequests(s),
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		log.Print("shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(shutdownCtx)
	}()

	log.Printf("serving %s on http://%s", absRoot, *addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
