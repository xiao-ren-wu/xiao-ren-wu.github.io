package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// fixture mirrors the shape of the real build output: a root index.html, a
// 404 page, percent-encoded UTF-8 knowledge pages stored as directory/index.html,
// hashed bundles under /_astro/ and a .git directory inside the published root.
func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()

	files := map[string]string{
		"index.html":                            "<html>home</html>",
		"404.html":                              "<html>not found</html>",
		"favicon.svg":                           "<svg/>",
		"_astro/page.B1D-nYk3.js":               "export default 1",
		"pagefind/wasm.unknown.pagefind":        "\x00asm\x01\x00\x00\x00",
		"pagefind/index/zh-cn_1618cc9.pf_index": "compressed",
		"knowledge/算法/index.html":               "<html>算法</html>",
		"knowledge/算法/二叉树/二叉树的右视图/index.html": "<html>右视图</html>",
		".git/config":       "[core]",
		"noindex/readme.md": "# no index",
	}
	for name, body := range files {
		full := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func TestContentType(t *testing.T) {
	tests := []struct {
		file string
		want string
	}{
		// The one that actually matters: pagefind's WebAssembly module.
		{"pagefind/wasm.unknown.pagefind", "application/wasm"},
		{"knowledge/算法/index.html", "text/html; charset=utf-8"},
		{"_astro/page.B1D-nYk3.js", "text/javascript; charset=utf-8"},
		{"pagefind/index/zh-cn_1618cc9.pf_index", "application/octet-stream"},
		{"favicon.svg", "image/svg+xml"},
		{"sitemap-0.xml", "application/xml; charset=utf-8"},
		{"pagefind/pagefind-entry.json", "application/json; charset=utf-8"},
		{"no-such-extension.xyz", "application/octet-stream"},
	}
	for _, tt := range tests {
		if got := contentType(tt.file); got != tt.want {
			t.Errorf("contentType(%q) = %q, want %q", tt.file, got, tt.want)
		}
	}
}

func TestResolve(t *testing.T) {
	s, err := newSite(fixture(t))
	if err != nil {
		t.Fatal(err)
	}
	index := filepath.FromSlash("knowledge/算法/二叉树/二叉树的右视图/index.html")

	tests := []struct {
		name      string
		url       string
		wantFile  string
		wantRedir string
		wantOK    bool
	}{
		{name: "root serves index", url: "/", wantFile: "index.html", wantOK: true},
		{name: "explicit file", url: "/index.html", wantFile: "index.html", wantOK: true},
		{name: "utf8 directory", url: "/knowledge/算法/", wantFile: "knowledge/算法/index.html", wantOK: true},
		{name: "utf8 page", url: "/knowledge/算法/二叉树/二叉树的右视图/", wantFile: index, wantOK: true},
		{name: "utf8 page without slash redirects", url: "/knowledge/算法/二叉树/二叉树的右视图", wantRedir: "/knowledge/算法/二叉树/二叉树的右视图/", wantOK: true},
		{name: "hashed bundle", url: "/_astro/page.B1D-nYk3.js", wantFile: "_astro/page.B1D-nYk3.js", wantOK: true},
		{name: "pagefind wasm", url: "/pagefind/wasm.unknown.pagefind", wantFile: "pagefind/wasm.unknown.pagefind", wantOK: true},
		{name: "dotfile refused", url: "/.git/config"},
		{name: "dotdir refused", url: "/.git/"},
		{name: "traversal contained", url: "/../../etc/passwd"},
		{name: "double slash normalised", url: "//index.html", wantFile: "index.html", wantOK: true},
		{name: "missing page", url: "/knowledge/算法/链表/"},
		{name: "directory without index", url: "/noindex/"},
		{name: "file under missing dir", url: "/nope/index.html"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, redir, ok := s.resolve(tt.url)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v (file=%q redir=%q)", ok, tt.wantOK, file, redir)
			}
			if redir != tt.wantRedir {
				t.Errorf("redirect = %q, want %q", redir, tt.wantRedir)
			}
			if file != "" {
				want := filepath.Join(s.root, filepath.FromSlash(tt.wantFile))
				if file != want {
					t.Errorf("file = %q, want %q", file, want)
				}
			}
		})
	}
}

func TestServeHTTP(t *testing.T) {
	s, err := newSite(fixture(t))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("serves utf8 page", func(t *testing.T) {
		w := httptest.NewRecorder()
		s.ServeHTTP(w, httptest.NewRequest("GET", "/knowledge/%E7%AE%97%E6%B3%95/%E4%BA%8C%E5%8F%89%E6%A0%91/%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%8F%B3%E8%A7%86%E5%9B%BE/", nil))
		if w.Code != http.StatusOK {
			t.Fatalf("status = %d, want 200", w.Code)
		}
		if got := w.Body.String(); got != "<html>右视图</html>" {
			t.Errorf("body = %q", got)
		}
		if got := w.Header().Get("Content-Type"); got != "text/html; charset=utf-8" {
			t.Errorf("Content-Type = %q", got)
		}
	})

	t.Run("pagefind wasm is typed as wasm", func(t *testing.T) {
		w := httptest.NewRecorder()
		s.ServeHTTP(w, httptest.NewRequest("GET", "/pagefind/wasm.unknown.pagefind", nil))
		if got := w.Header().Get("Content-Type"); got != "application/wasm" {
			t.Errorf("Content-Type = %q, want application/wasm", got)
		}
	})

	t.Run("hashed assets are immutable, html is not", func(t *testing.T) {
		for path, want := range map[string]string{
			"/_astro/page.B1D-nYk3.js": "public, max-age=31536000, immutable",
			"/index.html":              "no-cache",
		} {
			w := httptest.NewRecorder()
			s.ServeHTTP(w, httptest.NewRequest("GET", path, nil))
			if got := w.Header().Get("Cache-Control"); got != want {
				t.Errorf("%s Cache-Control = %q, want %q", path, got, want)
			}
		}
	})

	t.Run("missing path returns 404 page with 404 status", func(t *testing.T) {
		w := httptest.NewRecorder()
		s.ServeHTTP(w, httptest.NewRequest("GET", "/knowledge/算法/链表/", nil))
		if w.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want 404", w.Code)
		}
		if got := w.Body.String(); got != "<html>not found</html>" {
			t.Errorf("body = %q, want the site 404 page", got)
		}
	})

	t.Run("dotfile returns 404 rather than its contents", func(t *testing.T) {
		w := httptest.NewRecorder()
		s.ServeHTTP(w, httptest.NewRequest("GET", "/.git/config", nil))
		if w.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want 404", w.Code)
		}
		if body := w.Body.String(); body == "[core]" {
			t.Error("leaked .git/config")
		}
	})

	t.Run("redirect keeps the query string", func(t *testing.T) {
		w := httptest.NewRecorder()
		s.ServeHTTP(w, httptest.NewRequest("GET", "/knowledge/%E7%AE%97%E6%B3%95?q=%E4%BA%8C%E5%8F%89%E6%A0%91", nil))
		if w.Code != http.StatusMovedPermanently {
			t.Fatalf("status = %d, want 301", w.Code)
		}
		if got, want := w.Header().Get("Location"), "/knowledge/%E7%AE%97%E6%B3%95/?q=%E4%BA%8C%E5%8F%89%E6%A0%91"; got != want {
			t.Errorf("Location = %q, want %q", got, want)
		}
	})

	t.Run("If-None-Match revalidates to 304", func(t *testing.T) {
		first := httptest.NewRecorder()
		s.ServeHTTP(first, httptest.NewRequest("GET", "/index.html", nil))
		tag := first.Header().Get("Etag")
		if tag == "" {
			t.Fatal("no ETag set")
		}
		r := httptest.NewRequest("GET", "/index.html", nil)
		r.Header.Set("If-None-Match", tag)
		w := httptest.NewRecorder()
		s.ServeHTTP(w, r)
		if w.Code != http.StatusNotModified {
			t.Errorf("status = %d, want 304", w.Code)
		}
	})

	t.Run("head has no body", func(t *testing.T) {
		w := httptest.NewRecorder()
		s.ServeHTTP(w, httptest.NewRequest("HEAD", "/index.html", nil))
		if w.Code != http.StatusOK {
			t.Fatalf("status = %d, want 200", w.Code)
		}
		if w.Body.Len() != 0 {
			t.Errorf("body length = %d, want 0", w.Body.Len())
		}
	})

	t.Run("write methods rejected", func(t *testing.T) {
		w := httptest.NewRecorder()
		s.ServeHTTP(w, httptest.NewRequest("POST", "/index.html", nil))
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("status = %d, want 405", w.Code)
		}
	})
}

func TestLooksLikeSite(t *testing.T) {
	root := fixture(t)

	if !looksLikeSite(root) {
		t.Error("the built site was not recognised")
	}
	// A directory with an index.html but none of this project's markers, which
	// is exactly the unrelated site that used to get served by mistake.
	other := t.TempDir()
	if err := os.WriteFile(filepath.Join(other, "index.html"), []byte("<html>别的站</html>"), 0o644); err != nil {
		t.Fatal(err)
	}
	if looksLikeSite(other) {
		t.Error("an unrelated directory with an index.html was accepted as the site root")
	}
	if looksLikeSite(filepath.Join(root, "nope")) {
		t.Error("a non-existent directory was accepted as the site root")
	}
}

func TestResolveRootExplicit(t *testing.T) {
	root := fixture(t)
	got, err := resolveRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if got != root {
		t.Errorf("resolveRoot(%q) = %q, want %q", root, got, root)
	}
}

func TestServeHTTPAgainstRealBuild(t *testing.T) {
	root, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "index.html")); err != nil {
		t.Skipf("no built site at %s", root)
	}
	s, err := newSite(root)
	if err != nil {
		t.Fatal(err)
	}

	// Paths taken from sitemap-0.xml, percent-encoded exactly as the browser
	// would request them.
	paths := []string{
		"/",
		"/knowledge/%E7%AE%97%E6%B3%95/",
		"/knowledge/%E7%AE%97%E6%B3%95/%E4%BA%8C%E5%8F%89%E6%A0%91/%E4%BA%8C%E5%8F%89%E6%A0%91%E4%B8%AD%E7%9A%84%E6%9C%80%E5%A4%A7%E8%B7%AF%E5%BE%84%E5%92%8C/",
		"/pagefind/wasm.unknown.pagefind",
		"/sitemap-index.xml",
		"/favicon.svg",
	}
	for _, p := range paths {
		r := httptest.NewRequest("GET", p, nil)
		r.URL.RawPath = p
		w := httptest.NewRecorder()
		s.ServeHTTP(w, r)
		if w.Code != http.StatusOK {
			t.Errorf("GET %s = %d, want 200", p, w.Code)
			continue
		}
		if body, _ := io.ReadAll(w.Result().Body); len(body) == 0 {
			t.Errorf("GET %s returned an empty body", p)
		}
		if w.Header().Get("Content-Type") == "" {
			t.Errorf("GET %s has no Content-Type", p)
		}
	}
}
