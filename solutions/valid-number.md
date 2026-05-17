# 有效数字

## 题目描述

有效数字（按顺序）可以分成以下几个部分：

1. 一个 **小数** 或者 **整数**
2. （可选）一个 `'e'` 或 `'E'`，后面跟着一个 **整数**

**小数** 可以分成以下几个部分：
- （可选）一个符号字符（`'+'` 或 `'-'`）
- 下述格式之一：
  1. 至少一位数字，后面跟着一个点 `'.'`
  2. 至少一位数字，后面跟着一个点 `'.'`，后面再跟着至少一位数字
  3. 一个点 `'.'`，后面跟着至少一位数字

**整数** 可以分成以下几个部分：
- （可选）一个符号字符（`'+'` 或 `'-'`）
- 至少一位数字

部分有效数字：`["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]`

部分无效数字：`["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]`

**示例 1：**

```
输入：s = "0"
输出：true
```

**示例 2：**

```
输入：s = "e"
输出：false
```

**示例 3：**

```
输入：s = "."
输出：false
```

**示例 4：**

```
输入：s = ".1"
输出：true
```

**提示：**

- `1 <= s.length <= 20`
- `s` 仅含英文字母（大写和小写）、数字（`0-9`）、加号 `'+'`、减号 `'-'`、空格 `' '` 或者点 `'.'`

## 解题思路

### 核心思想：按规则扫描

用几个布尔标记位记录扫描过程中的状态：

| 标记 | 含义 |
|------|------|
| `seenDigit` | 是否见过数字 |
| `seenDot` | 是否见过小数点 |
| `seenE` | 是否见过 e/E |
| `numberAfterE` | e 之后是否有数字（必须是整数） |

### 扫描规则

遍历每个字符 `c`：

1. **数字 `0-9`**：`seenDigit = true`，如果在 e 之后则 `numberAfterE = true`
2. **`'+'` / `'-'`**：只能在开头或 e/E 的后面紧邻的位置出现
3. **`'.'`**：只能出现一次，且不能在 e 之后
4. **`'e'` / `'E'`**：只能出现一次，且之前必须有数字，之后必须有数字
5. **其它字符**：直接返回 false

### 图解示例

```
s = "-123.456e789"

  - → 符号位，合法（开头）
  1 → seenDigit = true
  2 → seenDigit = true
  3 → seenDigit = true
  . → seenDot = true（第一次，且未出现 e）
  4 → seenDigit = true
  5 → seenDigit = true
  6 → seenDigit = true
  e → seenE = true，seenDigit = true ✓，numberAfterE = false
  7 → seenDigit = true，numberAfterE = true
  8 → seenDigit = true，numberAfterE = true
  9 → seenDigit = true，numberAfterE = true

结果：seenDigit = true && numberAfterE = true → true
```

### 复杂度分析

- **时间复杂度**：O(n)，一次遍历
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public boolean isNumber(String s) {
        boolean seenDigit = false;
        boolean seenDot = false;
        boolean seenE = false;
        boolean numberAfterE = true;

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);

            if (Character.isDigit(c)) {
                seenDigit = true;
                if (seenE) numberAfterE = true;
            } else if (c == '+' || c == '-') {
                if (i > 0 && s.charAt(i - 1) != 'e' && s.charAt(i - 1) != 'E') {
                    return false;
                }
            } else if (c == '.') {
                if (seenDot || seenE) return false;
                seenDot = true;
            } else if (c == 'e' || c == 'E') {
                if (seenE || !seenDigit) return false;
                seenE = true;
                numberAfterE = false;
            } else {
                return false;
            }
        }

        return seenDigit && numberAfterE;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| `numberAfterE` 初始为 true | 当没有 e 时，不需要 e 后有数字；有 e 后再设回 false 等待数字 |
| 符号位位置限制 | 只能在开头或 e/E 后面紧邻一个，如 `-3e-5` 合法，`3-5` 不合法 |
| 小数点不能在 e 之后 | 科学计数法指数必须是整数，如 `3e-2.5` 不合法 |
| `"."` 单独不合法 | 必须有数字配合，如 `".1"` 或 `"1."` 合法 |
| 首字符的 `'.'` 合法 | 只要后面有数字即可，如 `".5"` → true |
