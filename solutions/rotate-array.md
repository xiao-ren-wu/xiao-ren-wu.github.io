# 轮转数组

## 题目描述

给定一个整数数组 `nums`，将数组中的元素向右轮转 `k` 个位置，其中 `k` 是非负数。要求 O(1) 额外空间。

**示例 1：**

```
输入：nums = [1,2,3,4,5,6,7], k = 3
输出：[5,6,7,1,2,3,4]
```

**示例 2：**

```
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
```

## 解题思路

### 核心思想：三次翻转

```
k = k % n
翻转整个数组         → [7,6,5,4,3,2,1]
翻转前 k 个元素      → [5,6,7,4,3,2,1]
翻转后 n-k 个元素    → [5,6,7,1,2,3,4]
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public void rotate(int[] nums, int k) {
        int n = nums.length;
        k %= n;

        reverse(nums, 0, n - 1);
        reverse(nums, 0, k - 1);
        reverse(nums, k, n - 1);
    }

    private void reverse(int[] nums, int left, int right) {
        while (left < right) {
            int tmp = nums[left];
            nums[left] = nums[right];
            nums[right] = tmp;
            left++;
            right--;
        }
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| k 取模 | k 可能大于 n，先取模避免多余操作 |
| 三次翻转顺序 | 整体反转 → 前段反转 → 后段反转（向右轮转） |
| 向左轮转 | 整体反转 → 后段反转 → 前段反转 |
