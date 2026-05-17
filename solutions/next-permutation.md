# 下一个排列

## 题目描述

整数数组的一个 **排列** 就是将其所有成员以序列或线性顺序排列。整数数组的 **下一个排列** 是指其整数的下一个字典序更大的排列。

给你一个整数数组 `nums`，找出 `nums` 的下一个排列。如果不存在下一个更大的排列，则将数组重新排列为最小的排列（即升序排列）。

必须原地修改。

**示例 1：**

```
输入：nums = [1,2,3]
输出：[1,3,2]
```

**示例 2：**

```
输入：nums = [3,2,1]
输出：[1,2,3]
```

**示例 3：**

```
输入：nums = [1,1,5]
输出：[1,5,1]
```

**提示：**

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 100`

## 解题思路

### 核心思想：找规律

```
步骤：
1. 从右向左找到第一个相邻升序对 (i, i+1)：nums[i] < nums[i+1]
2. 如果找不到（整个数组降序），反转整个数组
3. 在 [i+1, n-1] 中从右向左找到第一个 > nums[i] 的数 nums[j]
4. 交换 nums[i] 和 nums[j]
5. 反转 [i+1, n-1]
```

### 图解示例

```
nums = [1,2,3]

从右向左找升序对：
  3 > 2 ✓ → 没有
  2 < 3 ✓ → i=1

在 [2, n-1] 中找 > nums[1]=2 的最右边的数 → nums[2]=3
交换 → [1,3,2]
反转 [2, n-1] → [1,3,2]

结果：[1,3,2]
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public void nextPermutation(int[] nums) {
        int i = nums.length - 2;
        while (i >= 0 && nums[i] >= nums[i + 1]) i--;

        if (i >= 0) {
            int j = nums.length - 1;
            while (nums[j] <= nums[i]) j--;
            swap(nums, i, j);
        }

        reverse(nums, i + 1, nums.length - 1);
    }

    private void swap(int[] nums, int i, int j) {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }

    private void reverse(int[] nums, int start, int end) {
        while (start < end) {
            swap(nums, start++, end--);
        }
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| i 从右向左找 | 找到第一次出现 nums[i] < nums[i+1] 的位置 |
| j 也从右向左找 | 找第一个 > nums[i] 的数（保证是最小的大数） |
| 找不到 i 时直接反转 | 说明整个数组是降序的，下一个排列是最小的升序排列 |
| 反转的是 i+1 到末尾 | 反转后这部分变成升序，是 i 位置后的最小排列 |
