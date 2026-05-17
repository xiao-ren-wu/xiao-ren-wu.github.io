# 在排序数组中查找元素的第一个和最后一个位置

## 题目描述

给你一个按照非递减顺序排列的整数数组 `nums`，和一个目标值 `target`。找出给定目标值在数组中的开始位置和结束位置。如果不存在返回 `[-1, -1]`。要求 O(log n)。

**示例 1：**

```
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
```

**示例 2：**

```
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
```

**示例 3：**

```
输入：nums = [], target = 0
输出：[-1,-1]
```

## 解题思路

### 核心思想：两次二分

用两次二分分别找第一个位置和最后一个位置。

```
找第一个 ≥ target 的位置（左边界）：
  while left < right:
    mid = (left + right) / 2
    if nums[mid] >= target: right = mid
    else: left = mid + 1

找第一个 > target 的位置（右边界+1）：
  while left < right:
    mid = (left + right) / 2
    if nums[mid] > target: right = mid
    else: left = mid + 1
```

### 复杂度分析

- **时间复杂度**：O(log n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int[] searchRange(int[] nums, int target) {
        int[] res = new int[]{-1, -1};
        if (nums.length == 0) return res;

        int left = 0, right = nums.length;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] >= target) right = mid;
            else left = mid + 1;
        }
        if (left == nums.length || nums[left] != target) return res;
        res[0] = left;

        right = nums.length;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] > target) right = mid;
            else left = mid + 1;
        }
        res[1] = left - 1;
        return res;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 二分查找左闭右开 | `right = nums.length` 避免 +1/-1 混淆 |
| 第一个位置的条件是 `>=` | 找第一个 ≥ target 的位置 |
| 最后一个位置的条件是 `>` | 找第一个 > target 的位置再 -1 |
| 检查越界和值 | 左边界找到后要检查是否越界以及值是否等于 target |
