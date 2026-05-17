# 四数之和

## 题目描述

给你一个由 `n` 个整数组成的数组 `nums`，和一个目标值 `target`。找出并返回满足条件且不重复的四元组。

**示例 1：**

```
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
```

**示例 2：**

```
输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]
```

**提示：**

- `1 <= nums.length <= 200`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`

## 解题思路

### 核心思想：排序 + 双指针 + 双层剪枝

在三数之和的基础上再套一层循环。固定前两个数，用双指针找后两个数。

关键：两处剪枝 + 三重去重。

```
排序
for i in [0, n-4]:
  去重 i
  剪枝：最小和 > target → break
        最大和 < target → continue

  for j in [i+1, n-3]:
    去重 j
    剪枝：当前最小和 > target → break
          当前最大和 < target → continue

    left=j+1, right=n-1
    while left < right:
      sum = nums[i]+nums[j]+nums[left]+nums[right]
      if sum == target: 记录并去重
      else if sum < target: left++
      else: right--
```

### 复杂度分析

- **时间复杂度**：O(n³)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public List<List<Integer>> fourSum(int[] nums, int target) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);
        int n = nums.length;

        for (int i = 0; i < n - 3; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            if ((long) nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target) break;
            if ((long) nums[i] + nums[n - 3] + nums[n - 2] + nums[n - 1] < target) continue;

            for (int j = i + 1; j < n - 2; j++) {
                if (j > i + 1 && nums[j] == nums[j - 1]) continue;
                if ((long) nums[i] + nums[j] + nums[j + 1] + nums[j + 2] > target) break;
                if ((long) nums[i] + nums[j] + nums[n - 2] + nums[n - 1] < target) continue;

                int left = j + 1, right = n - 1;
                while (left < right) {
                    long sum = (long) nums[i] + nums[j] + nums[left] + nums[right];
                    if (sum == target) {
                        res.add(Arrays.asList(nums[i], nums[j], nums[left], nums[right]));
                        while (left < right && nums[left] == nums[left + 1]) left++;
                        while (left < right && nums[right] == nums[right - 1]) right--;
                        left++; right--;
                    } else if (sum < target) {
                        left++;
                    } else {
                        right--;
                    }
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
| 用 long 防止溢出 | nums[i] 范围较大，四数之和可能溢出 int |
| 双重剪枝 | 最小和 > target 则退出，最大和 < target 则跳过 |
| 三重重去重 | i、j、left/right 都要去重 |
