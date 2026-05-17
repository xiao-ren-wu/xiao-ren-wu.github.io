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
