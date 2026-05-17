# 超级回文数

## 题目描述

如果一个正整数自身是回文数，并且它也是某个回文数的平方，那么我们称这个数为超级回文数。

现在，给定两个正整数 `left` 和 `right`（以字符串形式表示），返回落在范围 `[left, right]` 中的超级回文数的数目。

**示例 1：**

```
输入：left = "4", right = "1000"
输出：4
解释：4, 9, 121, 484 是超级回文数。
     4 = 2², 9 = 3², 121 = 11², 484 = 22²。
```

**示例 2：**

```
输入：left = "1", right = "10"
输出：2
解释：1, 4 是超级回文数。
```

**提示：**

- `1 <= left.length, right.length <= 18`
- `left` 和 `right` 是用字符串表示的整数
- `1 <= parseInt(left) <= parseInt(right) <= 10^18`

## 解题思路

### 核心思想：构造回文根

直接枚举 `[left, right]` 范围内的所有数并检查是不是超级回文，范围太大（10^18）不可行。

更好的思路：**构造回文根**。

因为超级回文数 = (回文根)²，且右边界最大 10^18，所以回文根最大为 10^9。

但回文根本身也是回文数，所以只需枚举所有回文数 ≤ 10^9 的即可，这大约是 10^5 个。

### 算法步骤

```
1. 将 left 和 right 转为整数 L、R
2. 从 1 开始枚举回文根：
   a. 生成回文数 p（通过数字镜像）
   b. 计算 p²
   c. 如果 p² > R，停止（后续只会更大）
   d. 如果 p² >= L 且 p² 也是回文数，计数 +1
3. 返回计数
```

### 生成回文数

有两种镜像方式：

```
奇数长度：1 → 101,  12 → 12321,  123 → 1234321
偶数长度：1 → 11,   12 → 1221,   123 → 123321
```

### 图解示例

```
寻找 [4, 1000] 内的超级回文数：

枚举回文根：
  1² = 1 → 不在范围
  2² = 4 → 是回文 ✓
  3² = 9 → 是回文 ✓
  11² = 121 → 是回文 ✓
  22² = 484 → 是回文 ✓
  101² = 10201 > 1000 → 停止

结果：{4, 9, 121, 484} → 4 个
```

### 复杂度分析

- **时间复杂度**：O(W^(1/4))，其中 W = 10^18，大约需要检查 ~100000 个回文根
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int superpalindromesInRange(String left, String right) {
        long L = Long.parseLong(left);
        long R = Long.parseLong(right);
        int count = 0;

        // 枚举所有可能的回文根
        for (int k = 1; ; k++) {
            // 奇数长度回文根
            long p = createPalindrome(k, true);
            long p2 = p * p;
            if (p2 > R) break;
            if (p2 >= L && isPalindrome(p2)) count++;

            // 偶数长度回文根
            p = createPalindrome(k, false);
            p2 = p * p;
            if (p2 > R) break;
            if (p2 >= L && isPalindrome(p2)) count++;
        }

        return count;
    }

    // 创建回文数
    private long createPalindrome(int num, boolean oddLength) {
        long palin = num;
        if (oddLength) num /= 10;
        while (num > 0) {
            palin = palin * 10 + num % 10;
            num /= 10;
        }
        return palin;
    }

    private boolean isPalindrome(long x) {
        String s = Long.toString(x);
        int i = 0, j = s.length() - 1;
        while (i < j) {
            if (s.charAt(i) != s.charAt(j)) return false;
            i++;
            j--;
        }
        return true;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 同时检查奇偶长度 | 回文根可以是奇数或偶数长度，两者都要枚举 |
| p² 可能溢出 long | 右边界 10^18 < Long.MAX_VALUE，所以 p² 不会溢出 |
| `p*p` 可能超过 R | 一旦超过就可以提前停止枚举 |
| 1 也是回文数 | 1²=1，在 left=1 时应该计数 |
| 生成长度控制 | while 循环中 num/10 用于奇数长度回文根，去掉最后一位再镜像 |
