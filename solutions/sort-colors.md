# 颜色分类

## 题目描述

给定一个包含红色、白色和蓝色共 n 个元素的数组 `nums`，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。使用整数 0、1、2 分别表示红色、白色、蓝色。要求不使用库函数且仅用 O(1) 额外空间的一趟扫描算法。

**示例 1：**

```
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
```

**示例 2：**

```
输入：nums = [2,0,1]
输出：[0,1,2]
```

## 解题思路

### 核心思想：三指针（荷兰国旗问题）

用三个指针：p0 指向 0 的右边界、p2 指向 2 的左边界、i 为当前遍历指针。

```
p0 = 0, p2 = n-1, i = 0

while i <= p2:
  if nums[i] == 0: swap(nums[i], nums[p0]); p0++; i++
  else if nums[i] == 2: swap(nums[i], nums[p2]); p2--
  else: i++
```

### 复杂度分析

- **时间复杂度**：O(n)，一趟扫描
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public void sortColors(int[] nums) {
        int p0 = 0, p2 = nums.length - 1, i = 0;

        while (i <= p2) {
            if (nums[i] == 0) {
                swap(nums, i, p0);
                p0++;
                i++;
            } else if (nums[i] == 2) {
                swap(nums, i, p2);
                p2--;
            } else {
                i++;
            }
        }
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
| 遇到 2 时 i 不前进 | 因为交换过来的可能是 0，需要再次判断 |
| p0 和 i 起始相同 | 确保 0 能正确交换到前面 |
| 循环条件 `i <= p2` | p2 之后都是 2，无需遍历 |
