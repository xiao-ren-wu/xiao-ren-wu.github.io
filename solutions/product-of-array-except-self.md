# 除自身以外数组的乘积

## 题目描述

给你一个整数数组 `nums`，返回数组 `answer`，其中 `answer[i]` 等于 `nums` 中除 `nums[i]` 之外其余各元素的乘积。题目保证乘积在 32 位整数范围内。要求 O(n) 时间且不使用除法。

**示例 1：**

```
输入：nums = [1,2,3,4]
输出：[24,12,8,6]
```

**示例 2：**

```
输入：nums = [-1,1,0,-3,3]
输出：[0,0,9,0,0]
```

## 解题思路

### 核心思想：前缀积 × 后缀积

先从左到右计算前缀积，再从右到左乘以后缀积，每个位置的结果 = 左边所有数的积 × 右边所有数的积。

```
// 第一遍：前缀积
res[0] = 1
for i in [1, n-1]:
  res[i] = res[i-1] * nums[i-1]

// 第二遍：乘以后缀积
right = 1
for i in [n-1, 0]:
  res[i] *= right
  right *= nums[i]
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(1)（不计输出空间）

## 代码

```java
class Solution {
    public int[] productExceptSelf(int[] nums) {
        int n = nums.length;
        int[] res = new int[n];

        res[0] = 1;
        for (int i = 1; i < n; i++) {
            res[i] = res[i - 1] * nums[i - 1];
        }

        int right = 1;
        for (int i = n - 1; i >= 0; i--) {
            res[i] *= right;
            right *= nums[i];
        }

        return res;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 两次遍历 | 第一次前缀积，第二次后缀积 |
| 不用除法 | 题目要求不能使用除法（且存在 0 时除法会出错） |
| 32 位整数 | 题目保证不会溢出，无需用 long |
