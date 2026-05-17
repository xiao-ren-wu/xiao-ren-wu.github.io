# 乘积最大子数组

## 题目描述

给你一个整数数组 `nums`，请你找出数组中乘积最大的非空连续子数组，并返回该子数组所对应的乘积。

**示例 1：**

```
输入：nums = [2,3,-2,4]
输出：6
解释：子数组 [2,3] 有最大乘积 6。
```

**示例 2：**

```
输入：nums = [-2,0,-1]
输出：0
```

## 解题思路

### 核心思想：动态规划

由于负数会使最大变最小、最小变最大，需要同时维护当前的最大乘积和最小乘积。

```
maxCurr = minCurr = res = nums[0]

for i in [1, n-1]:
  if nums[i] < 0: 交换 maxCurr 和 minCurr
  maxCurr = max(nums[i], maxCurr * nums[i])
  minCurr = min(nums[i], minCurr * nums[i])
  res = max(res, maxCurr)
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int maxProduct(int[] nums) {
        int maxCurr = nums[0], minCurr = nums[0], res = nums[0];

        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < 0) {
                int tmp = maxCurr;
                maxCurr = minCurr;
                minCurr = tmp;
            }

            maxCurr = Math.max(nums[i], maxCurr * nums[i]);
            minCurr = Math.min(nums[i], minCurr * nums[i]);

            res = Math.max(res, maxCurr);
        }

        return res;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 负数交换 | 遇到负数，最大变最小、最小变最大 |
| 丢弃子数组 | 如果 `nums[i] > maxCurr*nums[i]`，说明前面的子数组可以丢弃 |
| 0 的影响 | 遇到 0 会重置 maxCurr 和 minCurr 为 0，相当于重新开始 |
