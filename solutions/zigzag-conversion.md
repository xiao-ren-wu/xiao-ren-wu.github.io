# Z 字形变换

## 题目描述

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 `"PAYPALISHIRING"`，行数为 3 时，排列如下：

```
P   A   H   N
A P L S I I G
Y   I   R
```

之后，你的输出需要从左往右逐行读取：`"PAHNAPLSIIGYIR"`

**示例 1：**

```
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
```

**示例 2：**

```
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
```

**提示：**

- `1 <= s.length <= 1000`
- `s` 由英文字母（小写和大写）、`,` 和 `.` 组成
- `1 <= numRows <= 1000`

## 解题思路

### 核心思想：模拟行走方向

用 `numRows` 个 StringBuilder 分别存储每一行的字符。用一个方向变量控制上下移动：

```
rows[0..numRows-1]
curRow = 0, goingDown = true

for each char in s:
  rows[curRow].append(char)
  if curRow == 0: goingDown = true
  if curRow == numRows-1: goingDown = false
  curRow += goingDown ? 1 : -1

拼接所有 rows
```

### 图解示例

```
s = "PAYPALISHIRING", numRows = 3

行0: P     A     H     N
行1: A  P  L  S  I  I  G
行2: Y     I     R

方向: → 下 → 上 → 下 → 上 ...
P(下) A(上) Y(下) P(上) A(下) L(上) ...

最终拼接: P + AHN + APLSIIG + YIR = "PAHNAPLSIIGYIR"
```

### 复杂度分析

- **时间复杂度**：O(n)，遍历一次
- **空间复杂度**：O(n)，存储结果

## 代码

```java
class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) return s;

        StringBuilder[] rows = new StringBuilder[numRows];
        for (int i = 0; i < numRows; i++) rows[i] = new StringBuilder();

        int curRow = 0;
        boolean goingDown = true;

        for (char c : s.toCharArray()) {
            rows[curRow].append(c);
            if (curRow == 0) goingDown = true;
            if (curRow == numRows - 1) goingDown = false;
            curRow += goingDown ? 1 : -1;
        }

        StringBuilder res = new StringBuilder();
        for (StringBuilder row : rows) res.append(row);
        return res.toString();
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| numRows=1 时直接返回 | 因为方向不会反转，会死循环 |
| 方向在首行和末行切换 | 不是固定的"到底就换"，而是在 0 和 numRows-1 处触发 |
| 相邻字符的索引差 | 第一行和最后一行的字符间隔为 2*(numRows-1) |
