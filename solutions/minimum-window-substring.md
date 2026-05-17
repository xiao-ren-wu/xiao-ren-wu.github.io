# 最小覆盖子串

## 题目描述

给你一个字符串 `s` 和一个字符串 `t`，请在 `s` 中找出包含 `t` 所有字符的最小子串。

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

用一个可变大小的窗口在 `s` 上滑动，每次扩展右边界，当窗口包含 `t` 中所有字符时，尝试收缩左边界以找到最短长度。

### 关键优化：用 `required` 计数而非逐个比较

- `needMap`：t 中每个字符的需求量
- `required`：t 中不同字符的种类数（如 t="AABC" 则 required=3）
- `window`：当前窗口中每个字符的已匹配量
- `formed`：当前窗口中已达标的字符种类数

当 `formed == required` 时，窗口满足覆盖条件。

### 算法步骤

```
right 指针遍历 s，将字符加入窗口：

1. 遇到 t 中的字符 → 更新 window 计数
   - 如果 window[c] == needMap[c]，formed++

2. 当 formed == required（窗口已覆盖）：
   a. 更新最小长度
   b. 收缩 left：移出左边字符
      - 如果是 t 中的字符且 window[c] < needMap[c]，formed--
      - left++

3. 继续扩展 right
```

### 图解示例

```
s = "ADOBECODEBANC", t = "ABC"
needMap = {A:1, B:1, C:1}, required = 3

窗口滑动过程：
A D O B E C O D E B A N C
↑                          → right=0, window={A:1}, formed=1
↑ ↑                        → right=1, D 不在 t 中
↑   ↑                      → right=2, O 不在 t 中
↑     ↑                    → right=3, window={A:1, B:1}, formed=2
↑       ↑                  → right=4, E 不在 t 中
↑         ↑                → right=5, window={A:1, B:1, C:1}, formed=3 ← 覆盖！
↑         ↑                → left=0, 移出 A, formed=2
  ↑       ↑                → left=1, 移出 D(不在)
    ↑     ↑                → left=2, 移出 O(不在)
      ↑   ↑                → left=3, 移出 B, formed=1
继续扩展...

最终得到 "BANC" 是最短的覆盖子串。
```

### 复杂度分析

- **时间复杂度**：O(m + n)，左右指针各遍历一次
- **空间复杂度**：O(k)，k 是 t 中不同字符的种类数

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
| 用 `intValue()` 比较 Integer | 避免 `Integer` 对象缓存导致 `==` 比较出错 |
| `required` 是字符种类数 | 不是字符总个数，如 `"AABC"` 的 required=3（A、B、C） |
| formed 只增不减的逻辑 | 当从 "达标" 变成 "不达标" 时 formed--，只比 `==` |
| 窗口收缩时 leftChar 要重新取 | 不是缓存之前的 left 位置，因为 left 已经变了 |
| 用 `minStart` 记录起始位置 | 最后才截取子串，避免频繁 substring 操作 |
