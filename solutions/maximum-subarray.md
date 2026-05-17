# 最大子数组和

## 题目描述

给你一个整数数组 `nums`，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

**示例 1：**

```
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6。
```

**示例 2：**

```
输入：nums = [1]
输出：1
```

**示例 3：**

```
输入：nums = [5,4,-1,7,8]
输出：23
```

## 解题思路

### 核心思想：Kadane 算法

遍历数组，维护以当前位置结尾的最大子数组和。如果之前的和为负，则舍弃重新开始。

```
maxCurr = maxEndingHere = nums[0]

for i in [1, n-1]:
  maxEndingHere = max(nums[i], maxEndingHere + nums[i])
  maxCurr = max(maxCurr, maxEndingHere)
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int maxSubArray(int[] nums) {
        int maxCurr = nums[0];
        int maxEndingHere = nums[0];

        for (int i = 1; i < nums.length; i++) {
            maxEndingHere = Math.max(nums[i], maxEndingHere + nums[i]);
            maxCurr = Math.max(maxCurr, maxEndingHere);
        }

        return maxCurr;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 初始化用 nums[0] | 不能初始化为 0，因为数组可能全为负 |
| 不用分治的原因 | Kadane O(n) 比分治 O(n log n) 更优 |
| 返回最大值 | 不是返回 maxEndingHere，而是全局最大值 |
