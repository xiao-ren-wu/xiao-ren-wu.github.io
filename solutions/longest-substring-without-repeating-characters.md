# 无重复字符的最长子串

## 题目描述

给定一个字符串 `s`，请你找出其中不含有重复字符的最长子串的长度。

**示例 1：**

```
输入：s = "abcabcbb"
输出：3
解释：无重复字符的最长子串是 "abc"，长度为 3。
```

**示例 2：**

```
输入：s = "bbbbb"
输出：1
解释：无重复字符的最长子串是 "b"，长度为 1。
```

**示例 3：**

```
输入：s = "pwwkew"
输出：3
解释：无重复字符的最长子串是 "wke"，长度为 3。
```

**提示：**

- `0 <= s.length <= 5 * 10^4`
- `s` 由英文字母、数字、符号和空格组成

## 解题思路

### 核心思想：滑动窗口

用两个指针维护一个不含重复字符的窗口，右指针扩展窗口，当遇到重复字符时左指针跳转到重复字符的下一个位置。

用 HashMap 记录每个字符最后出现的位置。

```
left = 0, maxLen = 0

for right in [0, n-1]:
  if map 中存在 s[right]:
    left = max(left, map[s[right]] + 1)
  map[s[right]] = right
  maxLen = max(maxLen, right - left + 1)
```

### 图解示例

```
s = "abcabcbb"

a b c a b c b b
↑
left=0 right=0 map={a:0} max=1

a b c a b c b b
↑     ↑
left=1 right=3 发现 a 重复，left 跳到 1，map[a]=3
此时窗口为 "bca"，max=3

最终 max=3
```

### 复杂度分析

- **时间复杂度**：O(n)，每个字符访问一次
- **空间复杂度**：O(min(m, n))，m 为字符集大小

## 代码

```java
class Solution {
    public int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> map = new HashMap<>();
        int maxLen = 0;
        int left = 0;

        for (int right = 0; right < s.length(); right++) {
            char c = s.charAt(right);
            if (map.containsKey(c)) {
                left = Math.max(left, map.get(c) + 1);
            }
            map.put(c, right);
            maxLen = Math.max(maxLen, right - left + 1);
        }

        return maxLen;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| left 取 max 而非直接赋值 | 因为 left 只会向右移动，防止跳回已经处理过的位置 |
| 字符集较大的处理 | 可以用数组替代 HashMap 优化（ASCII 128 或 256） |
| 空串返回 0 | s.length()=0 时循环不执行，返回 0 正确 |
