# 整数转罗马数字

## 题目描述

罗马数字包含以下七种字符：`I`(1)、`V`(5)、`X`(10)、`L`(50)、`C`(100)、`D`(500)、`M`(1000)。

通常，小的数字在大的数字右边。但特殊情况如 4 不写作 IIII，而是 IV（5-1）。同样 9 写作 IX。

给定一个整数，将其转为罗马数字。

**示例 1：**

```
输入：num = 3 → "III"
输入：num = 4 → "IV"
输入：num = 9 → "IX"
输入：num = 58 → "LVIII" (L=50, V=5, III=3)
输入：num = 1994 → "MCMXCIV" (M=1000, CM=900, XC=90, IV=4)
```

**提示：**

- `1 <= num <= 3999`

## 解题思路

### 核心思想：贪心匹配

将罗马数字的 13 个特殊符号按值从大到小排序，对 num 尽可能匹配最大的符号。

```
values  = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
symbols = ["M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"]

for i in [0..12]:
  while num >= values[i]:
    result += symbols[i]
    num -= values[i]
```

### 复杂度分析

- **时间复杂度**：O(1)，有限次循环
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public String intToRoman(int num) {
        int[] values = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        String[] symbols = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};

        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < values.length; i++) {
            while (num >= values[i]) {
                sb.append(symbols[i]);
                num -= values[i];
            }
        }

        return sb.toString();
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 6 种特殊组合 | IV(4)、IX(9)、XL(40)、XC(90)、CD(400)、CM(900) |
| 贪心的正确性 | 罗马数字的符号值从大到小排列，贪心总是最优 |
| num 范围 1-3999 | 题目限制，不会出现更大的值 |
