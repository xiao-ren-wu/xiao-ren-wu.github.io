# 最小覆盖子串 (LCR 017)

## 题目描述

给定两个字符串 `s` 和 `t`，请在 `s` 中找出包含 `t` 所有字符的最小子串。（剑指 Offer II 017）

如果 `s` 中不存在这样的子串，则返回空字符串 `""`。

**示例 1：**

```
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释："BANC" 覆盖了 'A'、'B'、'C' 各一个。
```

**示例 2：**

```
输入：s = "a", t = "a"
输出："a"
```

**示例 3：**

```
输入：s = "a", t = "aa"
输出：""
解释：t 中有两个 'a'，s 中只有一个，无法覆盖。
```

**提示：**

- `1 <= s.length, t.length <= 10^5`
- `s` 和 `t` 由英文字母组成

## 解题思路

### 核心思想：滑动窗口

用两个指针维护一个窗口。右指针扩展窗口直到包含 t 的所有字符，然后左指针收缩窗口以找到最短长度。

### 关键优化：计数器和种类数

- `needMap`：t 中每个字符需要的数量
- `required`：t 中不同字符的种类数
- `window`：当前窗口中各字符的计数
- `formed`：当前窗口中已达标的字符种类数

当 `formed == required` 时，窗口已覆盖 t。

### 算法步骤

```
right 遍历 s：
  1. 将 s[right] 加入 window
  2. 如果 window[c] == needMap[c]，formed++
  3. 当 formed == required：
     a. 更新最小长度和起始位置
     b. 尝试收缩 left：移出 s[left]
        - 如果 window[c] < needMap[c]，formed--
        - left++
  4. right++
```

### 图解示例

```
s = "ADOBECODEBANC", t = "ABC"
needMap = {A:1, B:1, C:1}, required=3

A D O B E C O D E B A N C
↑                         right=0, window={A:1}, formed=1
↑ ↑                       right=1, D 不在 t 中
↑   ↑                     right=2, O 不在 t 中
↑     ↑                   right=3, window={A:1,B:1}, formed=2
↑       ↑                 right=4, E 不在 t 中
↑         ↑               right=5, window={A:1,B:1,C:1}, formed=3
↑         ↑               ← 覆盖！left=0，但 A 移出后 formed=2
  ↑       ↑               left=1 → 不在 t 中
    ↑     ↑               left=2 → 不在 t 中
      ↑   ↑               left=3 → 移出 B, formed=1

继续扩展...最终得到 "BANC" 最短。
```

### 复杂度分析

- **时间复杂度**：O(m + n)，每个字符最多被访问两次
- **空间复杂度**：O(k)，k 是 t 中不同字符的数量

## 代码

```java
class Solution {
    public String minWindow(String s, String t) {
        Map<Character, Integer> needMap = new HashMap<>();
        for (char c : t.toCharArray()) {
            needMap.put(c, needMap.getOrDefault(c, 0) + 1);
        }

        int required = needMap.size();
        int formed = 0;
        Map<Character, Integer> window = new HashMap<>();

        int left = 0, right = 0;
        int minLen = Integer.MAX_VALUE;
        int minStart = 0;

        while (right < s.length()) {
            char c = s.charAt(right);
            window.put(c, window.getOrDefault(c, 0) + 1);

            if (needMap.containsKey(c) && window.get(c).intValue() == needMap.get(c).intValue()) {
                formed++;
            }

            while (formed == required && left <= right) {
                if (right - left + 1 < minLen) {
                    minLen = right - left + 1;
                    minStart = left;
                }

                char leftChar = s.charAt(left);
                window.put(leftChar, window.get(leftChar) - 1);
                if (needMap.containsKey(leftChar)
                        && window.get(leftChar).intValue() < needMap.get(leftChar).intValue()) {
                    formed--;
                }
                left++;
            }

            right++;
        }

        return minLen == Integer.MAX_VALUE ? "" : s.substring(minStart, minStart + minLen);
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 本题与 LeetCode 76 相同 | LCR 017 是剑指 Offer II 版本，解法完全一致 |
| `intValue()` 比较 Integer | 避免 `==` 比较 Integer 对象的缓存问题 |
| required 是字符种类计数 | "AABC" 的 required=3（A、B、C），不是 4 |
| formed 当且仅当达标才 +1/ -1 | 只跟踪从达标 ↔ 不达标的变化时刻 |
| 记录起始位置，最后再截取 | 避免频繁 substring 造成 O(n²) 开销 |
