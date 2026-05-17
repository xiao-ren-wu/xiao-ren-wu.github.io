# 最长回文子串

## 题目描述

给你一个字符串 `s`，找到 `s` 中最长的回文子串。

**示例 1：**

```
输入：s = "babad"
输出："bab"
解释："aba" 也是有效答案。
```

**示例 2：**

```
输入：s = "cbbd"
输出："bb"
```

**提示：**

- `1 <= s.length <= 1000`
- `s` 仅由数字和英文字母组成

## 解题思路

### 核心思想：中心扩展法

回文串要么是奇数长度（中心是一个字符），要么是偶数长度（中心是两个字符）。枚举每个可能的中心，向两边扩展，记录最长回文。

```
for i in [0, n-1]:
  扩展以 i 为中心（奇数长度）
  扩展以 i 和 i+1 为中心（偶数长度）
  更新最长回文
```

### 图解示例

```
s = "babad"

以 i=1 ('a') 为中心扩展：
  b ← a → b  → "bab" 长度 3
  (再扩) ← 已到边界

以 i=2 ('b') 为中心扩展：
  a ← b → a  → "aba" 长度 3

最长 = "bab" 或 "aba"
```

### 复杂度分析

- **时间复杂度**：O(n²)，最坏情况每个中心扩展 O(n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    private int start = 0, maxLen = 0;

    public String longestPalindrome(String s) {
        if (s == null || s.length() < 2) return s;

        for (int i = 0; i < s.length(); i++) {
            expand(s, i, i);
            expand(s, i, i + 1);
        }

        return s.substring(start, start + maxLen);
    }

    private void expand(String s, int left, int right) {
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            left--;
            right++;
        }
        int len = right - left - 1;
        if (len > maxLen) {
            maxLen = len;
            start = left + 1;
        }
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 中心有 2n-1 个 | n 个字符中心 + n-1 个字符间中心 |
| while 结束后 left 和 right 多走了一步 | 回文子串范围为 [left+1, right-1] |
| DP 也可解但非最优 | O(n²) 时间 O(n²) 空间，不如中心扩展的 O(1) 空间 |
