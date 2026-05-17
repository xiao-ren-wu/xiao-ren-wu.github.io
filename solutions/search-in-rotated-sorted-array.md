# 搜索旋转排序数组

## 题目描述

整数数组 `nums` 按升序排列，数组中的值互不相同。在某个下标 `k` 上进行了旋转。给你旋转后的数组和一个整数 `target`，如果存在返回下标，否则返回 -1。要求 O(log n)。

**示例 1：**

```
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
```

**示例 2：**

```
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
```

**示例 3：**

```
输入：nums = [1], target = 0
输出：-1
```

## 解题思路

### 核心思想：二分查找

旋转数组有一个性质：将数组从中间分开，**总有一半是有序的**。

```
while left <= right:
  mid = (left + right) / 2
  if nums[mid] == target: return mid

  if nums[left] <= nums[mid]:  // 左半有序
    if target 在 [nums[left], nums[mid]) 内:
      right = mid - 1
    else:
      left = mid + 1
  else:  // 右半有序
    if target 在 (nums[mid], nums[right]] 内:
      left = mid + 1
    else:
      right = mid - 1
```

### 复杂度分析

- **时间复杂度**：O(log n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int search(int[] nums, int target) {
        int left = 0, right = nums.length - 1;

        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] == target) return mid;

            if (nums[left] <= nums[mid]) {
                if (nums[left] <= target && target < nums[mid]) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            } else {
                if (nums[mid] < target && target <= nums[right]) {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            }
        }

        return -1;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 用 `<=` 判断左半有序 | `nums[left] <= nums[mid]` 等号处理 mid==left 的情况 |
| 边界检查要包含 `=` | target 与 nums[left]/nums[right] 比较时用 `<=` |
| 先判断哪半有序 | 再判断 target 是否在有序的那半边 |
