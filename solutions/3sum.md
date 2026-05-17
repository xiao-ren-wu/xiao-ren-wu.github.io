# 三数之和

## 题目描述

给你一个整数数组 `nums`，判断是否存在三元组 `[nums[i], nums[j], nums[k]]` 满足 `i != j`、`i != k` 且 `j != k`，同时还满足 `nums[i] + nums[j] + nums[k] == 0`。请你返回所有和为 0 且不重复的三元组。

**示例 1：**

```
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
```

**示例 2：**

```
输入：nums = [0,1,1]
输出：[]
```

**示例 3：**

```
输入：nums = [0,0,0]
输出：[[0,0,0]]
```

**提示：**

- `3 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## 解题思路

### 核心思想：排序 + 双指针

先排序，然后固定一个数，用双指针在剩余部分找两数之和等于目标值。

关键：**去重**。

```
排序 nums

for i in [0, n-3]:
  跳过重复的 i（去重）
  if nums[i] > 0: break（因为后面都 > 0）

  left = i+1, right = n-1, target = -nums[i]
  while left < right:
    sum = nums[left] + nums[right]
    if sum == target:
      记录 [i, left, right]
      left++, right--
      跳过重复的 left 和 right（去重）
    else if sum < target: left++
    else: right--
```

### 复杂度分析

- **时间复杂度**：O(n²)，排序 O(n log n)，两层循环 O(n²)
- **空间复杂度**：O(1)（不计输出空间）

## 代码

```java
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);

        for (int i = 0; i < nums.length - 2; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            if (nums[i] > 0) break;

            int left = i + 1, right = nums.length - 1;
            while (left < right) {
                int sum = nums[left] + nums[right];
                if (sum == -nums[i]) {
                    res.add(Arrays.asList(nums[i], nums[left], nums[right]));
                    while (left < right && nums[left] == nums[left + 1]) left++;
                    while (left < right && nums[right] == nums[right - 1]) right--;
                    left++;
                    right--;
                } else if (sum < -nums[i]) {
                    left++;
                } else {
                    right--;
                }
            }
        }

        return res;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 排序是前提 | 双指针依赖有序数组，去重也依赖排序 |
| 三重去重 | i 去重、left 去重、right 去重，缺一不可 |
| 固定数大于 0 可提前结束 | 因为排序后后面的数都 ≥ 当前数，三数和不可能为 0 |
