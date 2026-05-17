# 最短回文串

## 题目描述

给定一个字符串 `s`，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回用这种方式转换的最短回文串。

**示例 1：**

```
输入：s = "aacecaaa"
输出："aaacecaaa"
```

**示例 2：**

```
输入：s = "abcd"
输出："dcbabcd"
```

**提示：**

- `0 <= s.length <= 5 * 10^4`
- `s` 仅由小写英文字母组成

## 解题思路

### 核心思想

在字符串前面添加字符使其成为回文，等价于：

1. 找到 **最长的回文前缀**
2. 将剩余的**后缀反转**后，拼接到原字符串前面

```
s = "aacecaaa"

最长回文前缀： "aacecaa"
剩余后缀：     "a"
反转后缀：     "a"
结果：         "a" + "aacecaaa" = "aaacecaaa"
```

### 解法：KMP（最优）

要找到 s 的最长回文前缀，可以构造字符串 `t = s + "#" + reverse(s)`，然后对 t 求 KMP 的部分匹配表（next 数组）。next 数组的最后一个值就是「s 的最长回文前缀长度」。

原理：回文前缀正读反读相同，等价于它的反转是自身的后缀。KMP 的 next 数组正好能找出最长的既是前缀又是后缀的子串。

### 图解示例

```
s = "aacecaaa"
rev = "aaacecaa"
t = "aacecaaa#aaacecaa"

KMP next 数组（部分匹配表）：
t =  a a c e c a a a # a a a c e c a a
next=0 1 0 0 0 1 1 1 0 1 2 2 2 3 4 5 6

  ↑ last value = 7

最长回文前缀长度 = 7 → "aacecaa"
剩余后缀 = s.substring(7) = "a"
结果 = "a" + "aacecaaa" = "aaacecaaa"
```

```
s = "abcd"
rev = "dcba"
t = "abcd#dcba"

KMP next 数组末尾值 = 1（最长回文前缀 "a"）
剩余后缀 = "bcd"
结果 = "dcb" + "abcd" = "dcbabcd"
```

### 复杂度分析

- **时间复杂度**：O(n)，KMP 线性时间
- **空间复杂度**：O(n)

## 代码

```java
class Solution {
    public String shortestPalindrome(String s) {
        String rev = new StringBuilder(s).reverse().toString();
        String t = s + "#" + rev;

        int[] next = new int[t.length()];
        for (int i = 1; i < t.length(); i++) {
            int j = next[i - 1];
            while (j > 0 && t.charAt(i) != t.charAt(j)) {
                j = next[j - 1];
            }
            if (t.charAt(i) == t.charAt(j)) {
                j++;
            }
            next[i] = j;
        }

        int palindromeLen = next[t.length() - 1];
        String suffix = s.substring(palindromeLen);
        return new StringBuilder(suffix).reverse().toString() + s;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| `"#"` 分隔符的作用 | 防止前后缀越界匹配，例如 s="aaa"，不加分隔符会错误地匹配到更长的回文 |
| KMP next 数组含义 | next[i] 表示 t[0..i] 中最长公共前后缀的长度 |
| 最后一个值就是答案 | next[t.length()-1] 直接给出最长回文前缀的长度 |
| 空串处理 | s 为空时，rev 也为空，t 为 "#"，next 末尾为 0，返回空串正确 |
| 暴力法也可以但不推荐 | 从中间向两侧找最长回文前缀，最坏 O(n²)，n=50000 时超时 |
