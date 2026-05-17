# xiao-ren-wu.github.io — Agent Guide

## What this is

Simple GitHub Pages static site: a LeetCode problem list viewer (30-day challenge). Single page, Chinese UI.

## Repo structure

```
index.html          ← main app (HTML + CSS + JS + data + card rendering)
solution.html       ← solution viewer (loads & renders markdown from solutions/)
solutions/          ← individual solution markdown files, one per problem slug
  median-of-two-sorted-arrays.md
  substring-with-concatenation-of-all-words.md
  longest-valid-parentheses.md
  first-missing-positive.md
code.md             ← source file: concatenated solutions (used to generate solutions/)
```

No build tools, no package.json, no config. Pure static site.

## Key facts

- **Vanilla HTML/CSS/JS** — no frameworks, no bundlers, no npm.
- **Deployment**: push to `main`. GitHub Pages auto-deploys from `main` branch.
- **Local dev**: open `index.html` directly in a browser. No server required.
- **Data**: hardcoded JSON array inside the `<script>` tag (variable `jsonData`). This was extracted from LeetCode GraphQL API.
- **Language**: all UI text is zh-CN (Chinese).

## How to make changes

1. Edit `index.html` (or `solution.html` / `solutions/*.md`).
2. Open it in browser to verify.
3. Commit and push to `main` to deploy.

## Data model

Each problem object:
```json
{
  "difficulty": "HARD"|"MEDIUM"|"EASY",
  "questionFrontendId": "4",
  "title": "English title",
  "titleSlug": "english-title-slug",
  "translatedTitle": "中文标题",
  "topicTags": [{"nameTranslated": "数组"}, ...],
  "acRate": 0.44,
  "frequency": 62.3,
  "status": "TO_DO",
  "isInMyFavorites": false,
  "paidOnly": false
}
```

## UI features

- Search bar filters by Chinese/English title and category name.
- Difficulty filter buttons (全部/困难/中等/简单).
- Category filter buttons auto-generated from all unique tag names.
- Combined filtering: search + difficulty + category all work together.
- Each card shows: title, difficulty badge, tags, pass rate, frequency, and a "去做题" link to leetcode.cn.
- Each card shows a "题解" link if a solution exists, or "暂无题解" if not.
- Clicking "题解" opens `solution.html?slug=xxx` which renders the markdown with syntax highlighting (marked.js + highlight.js from CDN).
- `solution.html` has a "← 返回题目列表" link to go back to `index.html`.

## Style conventions

- Inline CSS in `<style>` tag (no external files).
- CSS custom properties not used.
- Gradient background (`linear-gradient(135deg, #667eea 0%, #764ba2 100%)`).
- Card design with left border accent, rounded corners, hover lift effect.
- Primary color: `#667eea`.
- Responsive breakpoint at 768px.

## Gotchas for agents

- **Single-page app** — main file is `index.html`. Most changes touch only this file.
- No linter, formatter, or type checker. Review manually.
- The JSON data set is large (~152KB file). Be careful not to corrupt it when editing.
- GitHub Pages has no server-side logic. Everything must be client-side.
- LeetCode link URLs use `leetcode.cn` (Chinese domain), not `leetcode.com`.
- Solution availability is tracked via a `solutionSlugs` Set in the JS of `index.html`. When adding a new solution, add the slug to this Set and create the corresponding `.md` file in `solutions/`.
- To add new solution files: create the `.md` in `solutions/<titleSlug>.md`, then add the slug to `solutionSlugs` in `index.html`.
