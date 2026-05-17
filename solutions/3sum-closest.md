# 最接近的三数之和

## 题目描述

给你一个长度为 `n` 的整数数组 `nums` 和一个目标值 `target`。请你从 `nums` 中选出三个整数，使它们的和与 `target` 最接近。返回这三个数的和。

**示例 1：**

```
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
```

**示例 2：**

```
输入：nums = [0,0,0], target = 1
输出：0
```

**提示：**

- `3 <= nums.length <= 1000`
- `-1000 <= nums[i] <= 1000`
- `-10^4 <= target <= 10^4`

## 解题思路

### 核心思想：排序 + 双指针

与三数之和类似，固定一个数后用双指针找最接近 target 的组合。

```
排序
best = nums[0] + nums[1] + nums[2]

for i in [0, n-3]:
  left = i+1, right = n-1
  while left < right:
    sum = nums[i] + nums[left] + nums[right]
    if sum == target: return target
    if |sum - target| < |best - target|: best = sum
    if sum < target: left++
    else: right--
```

### 复杂度分析

- **时间复杂度**：O(n²)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        int best = nums[0] + nums[1] + nums[2];

        for (int i = 0; i < nums.length - 2; i++) {
            int left = i + 1, right = nums.length - 1;
            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];
                if (sum == target) return target;
                if (Math.abs(sum - target) < Math.abs(best - target)) {
                    best = sum;
                }
                if (sum < target) left++;
                else right--;
            }
        }

        return best;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 不需要去重 | 找最接近的和，重复组合不影响结果 |
| 命中直接返回 | sum==target 时已是最优，无需继续 |
| best 的初始化 | 用前三个元素的和即可 |
