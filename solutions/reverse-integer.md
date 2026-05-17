# 整数反转

## 题目描述

给你一个 32 位的有符号整数 `x`，返回将 `x` 中的数字部分反转后的结果。如果反转后整数超过 32 位有符号整数的范围 `[−2³¹, 2³¹ − 1]`，就返回 0。

**示例 1：**

```
输入：x = 123
输出：321
```

**示例 2：**

```
输入：x = -123
输出：-321
```

**示例 3：**

```
输入：x = 120
输出：21
```

**示例 4：**

```
输入：x = 0
输出：0
```

**提示：**

- `-2³¹ <= x <= 2³¹ - 1`

## 解题思路

### 核心思想：逐位弹出

```
rev = 0
while x != 0:
  pop = x % 10
  x /= 10
  rev = rev * 10 + pop
```

关键问题：反转后可能溢出 int。需要在溢出前检查。

### 溢出检查

```
如果 rev > Integer.MAX_VALUE/10 或 rev < Integer.MIN_VALUE/10：
  反转后必然溢出，返回 0

如果 rev == Integer.MAX_VALUE/10：
  需要检查 pop > 7（因为 MAX_VALUE = 2147483647）
  
如果 rev == Integer.MIN_VALUE/10：
  需要检查 pop < -8（因为 MIN_VALUE = -2147483648）
```

### 复杂度分析

- **时间复杂度**：O(log x)，x 的位数
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int reverse(int x) {
        int rev = 0;

        while (x != 0) {
            int pop = x % 10;
            x /= 10;

            if (rev > Integer.MAX_VALUE / 10
                    || (rev == Integer.MAX_VALUE / 10 && pop > 7)) {
                return 0;
            }
            if (rev < Integer.MIN_VALUE / 10
                    || (rev == Integer.MIN_VALUE / 10 && pop < -8)) {
                return 0;
            }

            rev = rev * 10 + pop;
        }

        return rev;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| Java 取模负数结果 | `-123 % 10 = -3`，所以反转负数时不需要特殊处理 |
| 溢出检查在计算之前 | 先判断 rev*10 是否会溢出，再执行运算 |
| pop>7 / pop<-8 的判断 | Integer.MAX_VALUE=2147483647 末位 7，MIN_VALUE=-2147483648 末位 -8 |
| 反转后末尾的 0 自动消失 | 如 120 反转后得到 21，中间过程自动处理 |
