# 最大为 N 的数字组合

## 题目描述

给定一个按 **非递减顺序** 排列的数字数组 `digits`（数字是 1 到 9 的数字，且不包含 0）。你可以用任意次数 `digits` 中的数字来写数字。例如，如果 `digits = ["1","3","5"]`，我们可以写数字，如 `"13"`, `"551"`, 和 `"1351315"`。

返回可以生成的小于等于给定整数 `n` 的正整数的个数。

**示例 1：**

```
输入：digits = ["1","3","5","7"], n = 100
输出：20
解释：
可写出的 20 个数字是：
1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
```

**示例 2：**

```
输入：digits = ["1","4","9"], n = 1000000000
输出：29523
```

**示例 3：**

```
输入：digits = ["7"], n = 8
输出：1
```

**提示：**

- `1 <= digits.length <= 9`
- `digits[i]` 是范围 `['1', '9']` 中的数字
- `digits` 按非递减顺序排列
- `1 <= n <= 10^9`

## 解题思路

### 核心思想：分两部分计数

将问题分为两个部分：

1. **位数少于 n 的数字**：从 1 位到 len(n)-1 位，每位都有 `digits.length` 种选择
2. **位数等于 n 的数字**：逐位构造，需要仔细比较

### 算法步骤

```
function atMostNGivenDigitSet(digits, n):
  s = toString(n)
  len = s.length
  total = 0

  // 第一部分：所有位数少于 len 的数字
  for i in [1, len-1]:
    total += digits.length^i

  // 第二部分：位数等于 len 的数字
  for i in [0, len-1]:
    for each digit in digits:
      if digit < s[i]:
        total += digits.length^(len-i-1)
      else if digit == s[i]:
        // 继续处理下一位
        continue to next digit check
        如果都没有等于 s[i] 的数字 → return total
    // 如果循环完都没有提前返回，说明所有位都匹配了 s
    // 此时 n 本身就是一个有效数字
    return total + 1

  return total
```

### 图解示例

```
digits = [1,3,5,7], n = 100

len = 3

第一部分（1位和2位）：
  1位：4^1 = 4 个（1,3,5,7）
  2位：4^2 = 16 个
  小计：20

第二部分（3位，且 ≤ 100）：
  第0位('1')：digit<1 的没有，digit=1 ✓ → 继续
  第1位('0')：没有 digit < 0，也没有 digit = 0 → 停止
  所以没有 3 位数字

总计 = 20 ✓
```

```
digits = [1,3,5], n = 35

len = 2

第一部分（1位）：3^1 = 3 个

第二部分（2位，且 ≤ 35）：
  第0位('3')：digit<3 的有 '1' → 3^(2-0-1)=3个(11,13,15)
            digit=3 → 继续
  第1位('5')：digit<5 的有 '1','3' → 3^(2-1-1)=3^0=1 各 → 2个(31,33)
            digit=5 → n本身=35 ✓

总计 = 3 + 3 + 2 + 1 = 9
验证：{1,3,5,11,13,15,31,33,35} → 9个 ✓
```

### 复杂度分析

- **时间复杂度**：O(len(n) × |digits|)，最多 10×9
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int atMostNGivenDigitSet(String[] digits, int n) {
        String s = Integer.toString(n);
        int len = s.length();
        int d = digits.length;
        int total = 0;

        for (int i = 1; i < len; i++) {
            total += Math.pow(d, i);
        }

        for (int i = 0; i < len; i++) {
            boolean match = false;
            for (String digit : digits) {
                char c = digit.charAt(0);
                if (c < s.charAt(i)) {
                    total += Math.pow(d, len - i - 1);
                } else if (c == s.charAt(i)) {
                    match = true;
                    break;
                }
            }
            if (!match) return total;
        }

        return total + 1;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| digits 不含 0 | 所以不会出现前导零的问题，是简化条件 |
| digits 已排序 | 可以按顺序遍历，遇到大于当前位的就可以停止 |
| `Math.pow(d, len-i-1)` 可能不精确 | d 最大为 9，幂次不超过 9^9，用 double 转换没问题 |
| 位数等于 n 的处理 | 逐位检查，如果当前位在 digits 中没有匹配则立即返回 |
| n 本身也算 | 第三部分末尾要 +1（如果没有提前返回，n 本身就是一个有效数字） |
