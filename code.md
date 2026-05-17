# 寻找两个正序数组的中位数

## 题目描述

给定两个大小分别为 `m` 和 `n` 的正序（从小到大）数组 `nums1` 和 `nums2`。请你找出并返回这两个正序数组的 **中位数**。

算法的时间复杂度应该为 `O(log (m+n))`。

**示例 1：**

```
输入：nums1 = [1,3], nums2 = [2]
输出：2.0
解释：合并数组 = [1,2,3] ，中位数 2
```

**示例 2：**

```
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.5
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
```

**提示：**

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-10^6 <= nums1[i], nums2[i] <= 10^6`

---

## 解题思路

### 核心思想：二分查找分割法

中位数的本质：**将一组数分成左右两个数量相等的部分，且左边所有数 ≤ 右边所有数**。

对于两个有序数组，我们在两个数组上各画一条分割线，将数组分成左右两部分：

```
nums1: [a0, a1, ..., a(i-1)] | [ai, ..., a(m-1)]
         ← i 个元素 →          ← m-i 个 →

nums2: [b0, b1, ..., b(j-1)] | [bj, ..., b(n-1)]
         ← j 个元素 →          ← n-j 个 →
```

合并后的左半部分 = `nums1[0..i-1]` + `nums2[0..j-1]`，共 `i + j` 个元素。
合并后的右半部分 = `nums1[i..m-1]` + `nums2[j..n-1]`，共 `(m-i) + (n-j)` 个元素。

### 分割条件

我们需要找到一种分割，使得：

1. **左右元素数量相等**（或左边多一个）：
   
   - `i + j = (m + n + 1) / 2`（向上取整）

2. **左边所有元素 ≤ 右边所有元素**，等价于：
   
   - `nums1[i-1] ≤ nums2[j]`（nums1 左最大 ≤ nums2 右最小）
   - `nums2[j-1] ≤ nums1[i]`（nums2 左最大 ≤ nums1 右最小）

### 为什么需要交叉检查两个条件？

合并后的左边 = nums1 左半部分 + nums2 左半部分，最大值来自这两者中的较大者。
合并后的右边 = nums1 右半部分 + nums2 右半部分，最小值来自这两者中的较小者。

需要保证左边任意元素 ≤ 右边任意元素，所以既要检查 nums1 左边与 nums2 右边的交界，也要检查 nums2 左边与 nums1 右边的交界。

### 二分搜索过程

已知 `j = totalLeft - i`，问题简化为在 `[0, m]` 范围内二分搜索合适的 `i`：

- **命中了**：`nums1[i-1] ≤ nums2[j]` 且 `nums2[j-1] ≤ nums1[i]`
  
  - 奇数个元素：中位数 = `max(nums1[i-1], nums2[j-1])`
  - 偶数个元素：中位数 = `(max(左) + min(右)) / 2`

- **i 太大**（`nums1[i-1] > nums2[j]`）：nums1 左边拿太多了，`i` 左移

- **i 太小**（`nums2[j-1] > nums1[i]`）：nums1 左边拿太少了，`i` 右移

### 边界处理（哨兵值）

| 边界情况    | 问题                           | 哨兵方案                           |
| ------- | ---------------------------- | ------------------------------ |
| `i = 0` | nums1 左边无元素，`nums1[i-1]` 不存在 | `leftMax1 = MIN_VALUE`，条件自动成立  |
| `i = m` | nums1 右边无元素，`nums1[i]` 不存在   | `rightMin1 = MAX_VALUE`，条件自动成立 |
| `j = 0` | nums2 左边无元素，`nums2[j-1]` 不存在 | `leftMax2 = MIN_VALUE`，条件自动成立  |
| `j = n` | nums2 右边无元素，`nums2[j]` 不存在   | `rightMin2 = MAX_VALUE`，条件自动成立 |

### 为什么选择较短的数组进行二分搜索？

`j = totalLeft - i` 必须满足 `0 ≤ j ≤ n`，即：

- `i = 0` 时：`j = totalLeft ≤ n` → 需要 `totalLeft ≤ n`
- `i = m` 时：`j = totalLeft - m ≥ 0` → 需要 `totalLeft ≥ m`

因为 `totalLeft = (m+n+1)/2`，所以需要 `m ≤ totalLeft`，即 `m ≤ n`。

**保证始终对较短的数组做二分搜索**，这样 `i` 的搜索范围更小，时间复杂度为 `O(log(min(m, n)))`。

### 复杂度分析

- **时间复杂度**：`O(log(min(m, n)))` — 对较短的数组进行二分搜索
- **空间复杂度**：`O(1)` — 只使用了常数个变量

---

## 图解示例

```
nums1 = [1, 2]     m = 2
nums2 = [3, 4, 5]  n = 3

totalLeft = (2+3+1)/2 = 3
```

**第一次迭代：** `i = 1, j = 2`

```
nums1: [1] | [2]
nums2: [3, 4] | [5]

leftMax1 = 1, rightMin2 = 5      1 ≤ 5  ✅
leftMax2 = 4, rightMin1 = 2      4 ≤ 2  ❌
→ i 太小，右移
```

**第二次迭代：** `i = 2, j = 1`

```
nums1: [1, 2] | []
nums2: [3] | [4, 5]

leftMax1 = 2, rightMin2 = 4      2 ≤ 4  ✅
leftMax2 = 3, rightMin1 = MAX    3 ≤ MAX  ✅
→ 命中！

中位数 = max(leftMax1, leftMax2) = max(2, 3) = 3
```

---

## 最终代码

```java
/**
 * Median of Two Sorted Arrays — O(log(min(m, n)))
 *
 * Binary search the smaller array to find a partition where all left elements
 * are ≤ all right elements. Uses sentinel values (MIN_VALUE/MAX_VALUE) to
 * handle boundary cases uniformly.
 */
public class MedianOfTwoSortedArrays {

    public static double findMedianSortedArrays(int[] nums1, int[] nums2) {
        if (nums1.length > nums2.length) {
            return findMedianSortedArrays(nums2, nums1);
        }

        int m = nums1.length;
        int n = nums2.length;
        int totalLeft = (m + n + 1) / 2;

        int lo = 0;
        int hi = m;

        while (lo <= hi) {
            int i = lo + (hi - lo) / 2;
            int j = totalLeft - i;

            int leftMax1  = (i == 0)   ? Integer.MIN_VALUE : nums1[i - 1];
            int rightMin1 = (i == m)   ? Integer.MAX_VALUE : nums1[i];
            int leftMax2  = (j == 0)   ? Integer.MIN_VALUE : nums2[j - 1];
            int rightMin2 = (j == n)   ? Integer.MAX_VALUE : nums2[j];

            if (leftMax1 <= rightMin2 && leftMax2 <= rightMin1) {
                if ((m + n) % 2 == 1) {
                    return Math.max(leftMax1, leftMax2);
                } else {
                    return (Math.max(leftMax1, leftMax2) + Math.min(rightMin1, rightMin2)) / 2.0;
                }
            } else if (leftMax1 > rightMin2) {
                hi = i - 1;
            } else {
                lo = i + 1;
            }
        }

        throw new IllegalArgumentException("Input arrays are not sorted");
    }

    public static void main(String[] args) {
        System.out.println("Example 1: "
            + findMedianSortedArrays(new int[]{1, 3}, new int[]{2})
            + " (expected 2.0)");
        System.out.println("Example 2: "
            + findMedianSortedArrays(new int[]{1, 2}, new int[]{3, 4})
            + " (expected 2.5)");
        System.out.println("Edge (one empty): "
            + findMedianSortedArrays(new int[]{}, new int[]{1})
            + " (expected 1.0)");
        System.out.println("Edge (all same): "
            + findMedianSortedArrays(new int[]{1, 1}, new int[]{1, 1})
            + " (expected 1.0)");
        System.out.println("Edge (negative): "
            + findMedianSortedArrays(new int[]{-5, -3, -1}, new int[]{-4, -2, 0})
            + " (expected -2.5)");
    }
}
```

---

## 常见错误总结

| 错误                                   | 后果                | 修正                 |
| ------------------------------------ | ----------------- | ------------------ |
| 哨兵值用反：`j == n` 时给 `MIN_VALUE`        | 分割条件永远不满足，搜索乱掉    | 给 `MAX_VALUE`      |
| 二分边界用 `right = i` 而非 `right = i - 1` | i 被重复判定，可能导致死循环   | 用 `right = i - 1`  |
| 不交换数组：没有保证 `m ≤ n`                   | `j` 可能越界（负数或超过 n） | 始终让 `nums1` 是较短的数组 |

---

# 串联所有单词的子串

## 题目描述

给定一个字符串 `s` 和一个字符串数组 `words`。`words` 中所有字符串 **长度相同**。

`s` 中的 **串联子串** 是指一个包含 `words` 中所有字符串以任意顺序排列连接起来的子串。

例如，如果 `words = ["ab","cd","ef"]`，那么 `"abcdef"`，`"abefcd"`，`"cdabef"`，`"cdefab"`，`"efabcd"` 和 `"efcdab"` 都是串联子串。`"acdbef"` 不是串联子串，因为他不是任何 `words` 排列的连接。

返回所有串联子串在 `s` 中的开始索引。你可以以 **任意顺序** 返回答案。

**示例 1：**

```
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar"。
```

**示例 2：**

```
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[8]
解释：从索引 8 开始的子串是 "goodgoodbestword"，是 ["word","good","best","word"] 的一个排列。
```

**示例 3：**

```
输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]
```

**提示：**

- `1 <= s.length <= 10^4`
- `1 <= words.length <= 5000`
- `1 <= words[i].length <= 30`
- `words[i]` 和 `s` 由小写英文字母组成

---

## 解题思路

### 核心观察

`words` 中所有字符串 **长度相同**，这是本题最关键的条件。它意味着我们可以按单词粒度（`wordLen`）来切割 `s`，从而将问题转化为 **词级别的滑动窗口**。

设 `wordLen = words[0].length()`，`wordCount = words.length`，串联子串的总长度固定为 `wordLen * wordCount`。

### 算法：词级滑动窗口

#### 为什么需要枚举 offset？

串联子串可以从任意位置开始，但所有单词长度相同，所以起始位置对 `wordLen` 取模的结果决定了它落在哪个偏移类中。我们只需枚举 `[0, wordLen)` 这 `wordLen` 个偏移量，在每个偏移类内以 `wordLen` 为步长滑动窗口。

```
s = "barfoothefoobarman"

offset=0: bar | foo | the | foo | bar | man   匹配 [bar,foo]
offset=1: arf | oot | hef | oob | arm         乱序，不可能匹配
offset=2: rfo | oth | efo | oba | rm          乱序，不可能匹配
```

#### 滑动窗口规则

对每个 offset：

1. **遇到不在 `needMap` 中的词** → 整个窗口作废，清空 window，重置 left
2. **遇到合法词** → 加入 window，更新 matched
3. **如果当前词超量**（`window.get(cur) > needMap.get(cur)`）→ 从左侧逐个移出单词，同时更新 window 和 matched，直到当前词数量合法
4. **如果 matched == needMap.size()** → 记录 left，然后左移一个单词，继续滑动

### 为什么 matched 用词种类数而非总词数？

`matched` 统计的是「有多少种单词的数量已经恰好达标」。当 `matched == needMap.size()` 时，说明每种词的数量都正好满足要求，**总词数也必然等于 `wordCount`**。这种方式在收缩窗口时只需比较 `equals` 就能知道是否需要减少 matched，比维护总词数更简洁。

### 复杂度分析

- **时间复杂度**：`O(n)` — 每个字符最多被访问两次（进入窗口和移出窗口），`n = s.length()`
- **空间复杂度**：`O(m)` — `m = words` 中不同单词的数量

---

## 图解示例

```
s = "barfoothefoobarman", words = ["foo","bar"]

needMap = {"foo":1, "bar":1}, wordLen=3, wordCount=2, totalLeft=6
```

**offset = 0：**

```
right=0: "bar" → window={bar:1}    matched=1
right=3: "foo" → window={bar:1, foo:1}  matched=2
  matched(2)==needMap.size(2) → 记录 left=0
  左移: window={bar:0, foo:1}, matched=1

right=6: "the" → 非法词，重置
  matched=0, window={}, left=9

right=9:  "foo" → window={foo:1}    matched=1
right=12: "bar" → window={foo:1, bar:1}  matched=2
  matched(2)==needMap.size(2) → 记录 left=9
  左移: window={foo:0, bar:1}, matched=1

结果: [0, 9]
```

---

## 题解代码

```java
import java.util.*;

class Solution {
    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> resList = new ArrayList<>();
        int wordLen = words[0].length();
        if (wordLen * words.length > s.length()) {
            return resList;
        }
        Map<String, Integer> needMap = new HashMap<>();
        for (String word : words) {
            needMap.put(word, needMap.getOrDefault(word, 0) + 1);
        }
        for (int offset = 0; offset < wordLen; offset++) {
            int left = offset;
            int matched = 0;
            Map<String, Integer> window = new HashMap<>();
            for (int right = offset; right <= s.length() - wordLen; right += wordLen) {
                String cur = s.substring(right, right + wordLen);
                if (!needMap.containsKey(cur)) {
                    matched = 0;
                    window.clear();
                    left = right + wordLen;
                    continue;
                }
                window.put(cur, window.getOrDefault(cur, 0) + 1);
                if (Objects.equals(window.get(cur), needMap.get(cur))) {
                    matched++;
                }
                while (window.get(cur) > needMap.get(cur)) {
                    String leftWord = s.substring(left, left + wordLen);
                    if (Objects.equals(window.get(leftWord), needMap.get(leftWord))) {
                        matched--;
                    }
                    window.put(leftWord, window.get(leftWord) - 1);
                    left += wordLen;
                }
                if (matched == needMap.size()) {
                    resList.add(left);
                    String leftWord = s.substring(left, left + wordLen);
                    if (Objects.equals(window.get(leftWord), needMap.get(leftWord))) {
                        matched--;
                    }
                    window.put(leftWord, window.get(leftWord) - 1);
                    left += wordLen;
                }
            }
        }
        return resList;
    }
}
```

---

## 关键细节总结

| 细节                   | 说明                                                      |
| -------------------- | ------------------------------------------------------- |
| 枚举 `[0, wordLen)` 偏移 | 所有串联子串的起始位置必然落在这 wordLen 个余数类中                          |
| `matched` 统计词种类数     | 无需维护总词数，减少 equals 判断的次数                                 |
| while 收缩取 `leftWord` | 必须读 `s.substring(left, left+wordLen)` 拿到实际词，不能假定是 `cur` |
| 收缩时联动更新 `matched`    | 只有当 `Objects.equals` 成立（从达标变成不达标）时才递减 matched           |
| 遇到非法词清空重置            | 非法词之前积累的单词全部无效，不能继续累加                                   |
| 找到结果后左移一个单词          | 保持窗口继续滑动，注意 leftWord 的取值要在 left 更新之前                    |

## 常见错误

| 错误                                               | 后果                   |
| ------------------------------------------------ | -------------------- |
| while 只 `left += wordLen` 不更新 window             | 死循环                  |
| while 默认 left 位置 = cur 直接 `window.put(cur, ...)` | 删词错误，window 数据错乱     |
| 找到结果后不更新 matched                                 | 下一轮 matched 偏大，漏判或误判 |
| 遇到非法词只 continue 不清空窗口                            | 非法词前数据污染后续匹配         |
| 循环条件 `right < s.length()`                        | 少处理最后一个单词位置          |

---

# 最长有效括号

## 题目描述

给你一个只包含 `'('` 和 `')'` 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

左右括号匹配，即每个左括号都有对应的右括号将其闭合的字符串是格式正确的，比如 `"(()())"`。

**示例 1：**

```
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
```

**示例 2：**

```
输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
```

**示例 3：**

```
输入：s = ""
输出：0
```

---

## 解题思路

### 核心问题

求最长的**连续**有效括号子串长度。有效 = 每个 `(` 都能在它**之后**匹配到一个 `)`，且没有多余闭合。

关键难点：有效括号子串不一定是嵌套最深的，而是**连续闭合最长的**。如 `")()())"` 最长的是 `"()()"`（长度 4），而非 `"(())"`。

### 解法对比

| 解法        | 时间       | 空间       | 核心思想                   |
| --------- | -------- | -------- | ---------------------- |
| 栈         | O(n)     | O(n)     | 用栈跟踪未匹配括号的索引           |
| 动态规划      | O(n)     | O(n)     | dp[i] = 以 s[i] 结尾的有效长度 |
| **双指针扫描** | **O(n)** | **O(1)** | **双向扫描，左右计数**          |

---

## 解法一：双指针扫描（最优，O(1) 空间）

### 为什么需要双向扫描？

**正向扫描**：`left` 计 `(`，`right` 计 `)`

- `left == right` → 有效，更新答案
- `right > left` → 无效，重置

但正向扫描无法处理 `"(()"` 这种左括号多于右括号的情况——`left` 始终大于 `right`，永远不会触发 `left == right`。

**反向扫描**：逻辑翻转，`(` 计为 left，`)` 计为 right

- `left == right` → 有效，更新答案
- `left > right` → 无效，重置（反向时 left 代表 `(`，不能多于 `)`）

**互补原理**：正向解决 `right > left` 的截断，反向解决 `left > right` 的截断。两个方向互补，覆盖所有情况。

### 代码

```java
class Solution {
    public int longestValidParentheses(String s) {
        int maxLen = 0;
        int left = 0, right = 0;

        // 正向扫描：处理 "right > left" 的截断
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                left++;
            } else {
                right++;
            }
            if (left == right) {
                maxLen = Math.max(maxLen, 2 * right);
            } else if (right > left) {
                left = 0;
                right = 0;
            }
        }

        // 反向扫描：处理 "left > right" 的截断
        left = 0;
        right = 0;
        for (int i = s.length() - 1; i >= 0; i--) {
            if (s.charAt(i) == '(') {
                left++;
            } else {
                right++;
            }
            if (left == right) {
                maxLen = Math.max(maxLen, 2 * left);
            } else if (left > right) {
                left = 0;
                right = 0;
            }
        }

        return maxLen;
    }
}
```

**复杂度**：时间 O(n)，空间 O(1)

---

## 解法二：动态规划

### 状态定义

`dp[i]` = 以 `s[i]` 结尾的最长有效括号子串长度。

只有 `s[i] == ')'` 才可能形成有效括号子串。

### 状态转移

**情况 1：`s[i] == ')'` 且 `s[i-1] == '('` → 模式 `"...()"`**

```
dp[i] = dp[i-2] + 2      (i >= 2)
dp[i] = 2                (i == 1)
```

**情况 2：`s[i] == ')'` 且 `s[i-1] == ')'` → 模式 `"...))"`**

找到与 `s[i]` 配对的 `(` 的位置：`i - dp[i-1] - 1`

如果 `s[i - dp[i-1] - 1] == '('`，则：

```
dp[i] = dp[i-1] + 2 + dp[i - dp[i-1] - 2]
```

其中：

- `dp[i-1]` — 中间那段有效长度
- `2` — 本次配对的一对括号
- `dp[i - dp[i-1] - 2]` — 再往前的有效长度（拼接）

### 图解

```
索引:  0  1  2  3  4  5
s =   (  (  )  (  )  )

dp[4]: s[4]=')', s[3]=='(' → 情况1
       dp[4] = dp[2] + 2 = 2 + 2 = 4

dp[5]: s[5]=')', s[4]==')' → 情况2
       match = 5 - dp[4] - 1 = 5 - 4 - 1 = 0
       s[0]=='(' → 配对成功
       dp[5] = dp[4] + 2 + dp[-1] = 4 + 2 + 0 = 6
```

### 代码

```java
class Solution {
    public int longestValidParentheses(String s) {
        int maxLen = 0;
        int n = s.length();
        int[] dp = new int[n];

        for (int i = 1; i < n; i++) {
            if (s.charAt(i) == ')') {
                if (s.charAt(i - 1) == '(') {
                    // 情况1: ...()
                    dp[i] = (i >= 2 ? dp[i - 2] : 0) + 2;
                } else {
                    // 情况2: ...))
                    int match = i - dp[i - 1] - 1;
                    if (match >= 0 && s.charAt(match) == '(') {
                        dp[i] = dp[i - 1] + 2;
                        if (match - 1 >= 0) {
                            dp[i] += dp[match - 1];
                        }
                    }
                }
                maxLen = Math.max(maxLen, dp[i]);
            }
            // s[i] == '(' → dp[i] = 0，无需处理
        }

        return maxLen;
    }
}
```

**复杂度**：时间 O(n)，空间 O(n)

---

## 解法三：栈

### 核心思想

- 始终保持栈底为**当前有效子串的起始位置的前一个索引**
- 初始化压入 `-1`，表示有效子串从 0 开始
- 遇到 `(` 压入其索引
- 遇到 `)` 弹出栈顶
  - 如果栈空了，说明当前 `)` 没有匹配的 `(`，将当前索引入栈作为新的起始基准
  - 如果栈非空，`i - stack.peek()` 即为以当前位置结尾的有效长度

### 代码

```java
import java.util.ArrayDeque;
import java.util.Deque;

class Solution {
    public int longestValidParentheses(String s) {
        int maxLen = 0;
        Deque<Integer> stack = new ArrayDeque<>();
        stack.push(-1);

        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                stack.push(i);
            } else {
                stack.pop();
                if (stack.isEmpty()) {
                    stack.push(i);
                } else {
                    maxLen = Math.max(maxLen, i - stack.peek());
                }
            }
        }

        return maxLen;
    }
}
```

**复杂度**：时间 O(n)，空间 O(n)

---

## 总结对比

| 解法  | 优势           | 劣势       |
| --- | ------------ | -------- |
| 双指针 | 空间 O(1)、代码简洁 | 需要理解双向互补 |
| DP  | 状态转移清晰、易扩展   | 空间 O(n)  |
| 栈   | 思路最直观        | 空间 O(n)  |

双指针是面试最优解，DP 和栈也常考，建议三者都掌握。

---

# 缺失的第一个正数

## 题目描述

给你一个未排序的整数数组 `nums`，请你找出其中没有出现的最小的正整数。

要求实现时间复杂度为 O(n) 并且只使用**常数级别额外空间**的解决方案。

**示例 1：**

```
输入：nums = [1,2,0]
输出：3
解释：范围 [1,2] 中的数字都在数组中。
```

**示例 2：**

```
输入：nums = [3,4,-1,1]
输出：2
解释：1 在数组中，但 2 没有。
```

**示例 3：**

```
输入：nums = [7,8,9,11,12]
输出：1
解释：最小的正数 1 没有出现。
```

---

## 解题思路

### 核心观察

对于长度为 `n` 的数组，答案一定在 `[1, n+1]` 范围内。因为如果 `1~n` 全部出现，答案就是 `n+1`。

### 关键约束：O(1) 额外空间

不能用哈希表。**把数组本身当作哈希表**——将每个正整数 `x` 放到索引 `x-1` 的位置上。

### 算法：原地置换

**三步走：**

1. **原地置换** — 遍历数组，把每个在 `[1, n]` 范围内的数放到它应该在的位置
   
   - 当前值 `val = nums[i]`，它应该去的位置是 `nums[val-1]`
   - 如果 `nums[val-1] != val`，就交换，直到当前位置的值是正确值或者不需要处理

2. **扫描找答案** — 遍历 `i = 0..n-1`，第一个 `nums[i] != i+1` 的位置，答案就是 `i+1`

3. **兜底** — 如果 `1~n` 都在正确位置，答案就是 `n+1`

### 图解

```
nums = [3, 4, -1, 1]    n = 4

i=0: nums[0]=3  → 应放 index=2
     swap(0, 2) → [-1, 4, 3, 1]
     nums[0]=-1 → 跳过（不在 [1,4] 内）

i=1: nums[1]=4  → 应放 index=3
     swap(1, 3) → [-1, 1, 3, 4]
     nums[1]=1  → 应放 index=0
     swap(1, 0) → [1, -1, 3, 4]
     nums[1]=-1 → 跳过

i=2: nums[2]=3  → 应放 index=2，已在正确位置 ✓

i=3: nums[3]=4  → 应放 index=3，已在正确位置 ✓

扫描：nums[0]=1 ✓, nums[1]=-1 ≠ 2 → 答案 = 2
```

### 代码

```java
class Solution {
    public int firstMissingPositive(int[] nums) {
        int n = nums.length;

        // 原地置换：把每个正整数 val 放到索引 val-1
        for (int i = 0; i < n; i++) {
            while (nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i]) {
                int val = nums[i];
                nums[i] = nums[val - 1];
                nums[val - 1] = val;
            }
        }

        // 扫描：找到第一个 nums[i] != i+1 的位置
        for (int i = 0; i < n; i++) {
            if (nums[i] != i + 1) {
                return i + 1;
            }
        }

        // 兜底：1~n 全部出现，答案就是 n+1
        return n + 1;
    }
}
```

### 复杂度

- **时间复杂度**：O(n) — 每个元素最多被交换到正确位置一次，总共最多 n 次交换
- **空间复杂度**：O(1) — 原地交换，只用了常数变量

### 易错点

| 易错点                          | 说明                              |
| ---------------------------- | ------------------------------- |
| `while` 而非 `if`              | 交换后当前位置的新值可能仍需处理，必须循环直到值正确或不可处理 |
| `nums[i] <= n`               | 防止 `nums[i] - 1` 越界             |
| `nums[nums[i]-1] != nums[i]` | 目标位置已有正确值时不再交换，防止死循环（如 `[1,1]`） |
| 原地修改                         | 会破坏原数组（面试时可提"但这是允许的"）           |
