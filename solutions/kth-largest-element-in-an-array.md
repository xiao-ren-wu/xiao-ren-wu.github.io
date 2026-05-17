# 数组中的第 K 个最大元素

## 题目描述

给定整数数组 `nums` 和整数 `k`，请返回数组中第 k 个最大的元素。要求 O(n)。

**示例 1：**

```
输入：nums = [3,2,1,5,6,4], k = 2
输出：5
```

**示例 2：**

```
输入：nums = [3,2,3,1,2,4,5,5,6], k = 4
输出：4
```

## 解题思路

### 核心思想：快速选择

基于快排的 partition，每次确定 pivot 的位置。如果 pivot 正好是第 k 大（即 n-k），直接返回；否则只在包含目标的那一侧继续查找。

```
function quickSelect(nums, left, right, k):
  pivotIndex = partition(nums, left, right)
  if pivotIndex == k: return nums[k]
  else if pivotIndex < k: return quickSelect(nums, pivotIndex+1, right, k)
  else: return quickSelect(nums, left, pivotIndex-1, k)

function partition:
  pivot = nums[right]
  i = left
  for j in [left, right-1]:
    if nums[j] <= pivot: swap(nums, i, j); i++
  swap(nums, i, right)
  return i
```

### 复杂度分析

- **平均时间复杂度**：O(n)
- **最坏时间复杂度**：O(n²)，可通过随机 pivot 避免
- **空间复杂度**：O(log n)，递归深度

## 代码

```java
class Solution {
    public int findKthLargest(int[] nums, int k) {
        int n = nums.length;
        return quickSelect(nums, 0, n - 1, n - k);
    }

    private int quickSelect(int[] nums, int left, int right, int k) {
        if (left == right) return nums[left];

        int pivotIndex = partition(nums, left, right);
        if (pivotIndex == k) return nums[k];
        else if (pivotIndex < k) return quickSelect(nums, pivotIndex + 1, right, k);
        else return quickSelect(nums, left, pivotIndex - 1, k);
    }

    private int partition(int[] nums, int left, int right) {
        int pivot = nums[right];
        int i = left;
        for (int j = left; j < right; j++) {
            if (nums[j] <= pivot) {
                swap(nums, i, j);
                i++;
            }
        }
        swap(nums, i, right);
        return i;
    }

    private void swap(int[] nums, int i, int j) {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 第 k 大 = 第 n-k 小 | 找第 k 大在升序中的索引为 n-k |
| 只搜索一侧 | 不像快排两边都递归，快速选择只递归包含目标的一侧 |
| 最坏情况优化 | 随机选择 pivot 可避免 O(n²) |
