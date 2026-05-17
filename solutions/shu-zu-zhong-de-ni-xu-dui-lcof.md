# 交易逆序对的总数 (LCR 170)

## 题目描述

在股票交易中，如果前一天的股价高于后一天的股价，则存在一个「交易逆序对」。给定一个数组 `prices`，返回交易逆序对的总数。

即：给定一个数组，计算数组中逆序对的总数。逆序对是指满足 `i < j` 且 `prices[i] > prices[j]` 的数对。

**示例 1：**

```
输入：prices = [7,5,6,4]
输出：5
解释：逆序对为 (7,5)、(7,6)、(7,4)、(5,4)、(6,4)，共 5 个。
```

**示例 2：**

```
输入：prices = [1,3,2,3,1]
输出：4
```

**提示：**

- `0 <= prices.length <= 50000`
- `0 <= prices[i] <= 50000`

## 解题思路

### 核心思想：归并排序统计

在归并排序的合并过程中，当左半部分的 `nums[i] > nums[j]` 时，说明左半部分从 i 到 mid 的所有元素都大于 nums[j]，它们都可以与 nums[j] 构成逆序对。

### 算法步骤

```
mergeSort(nums, left, right):
  if left >= right: return 0

  mid = (left + right) / 2
  count = mergeSort(left, mid) + mergeSort(mid+1, right)

  // 合并 [left..mid] 和 [mid+1..right]
  i = left, j = mid+1, k = 0
  while i <= mid && j <= right:
    if nums[i] <= nums[j]:
      temp[k++] = nums[i++]
    else:
      count += mid - i + 1  // nums[i..mid] 都 > nums[j]
      temp[k++] = nums[j++]

  // 复制剩余元素
  while i <= mid: temp[k++] = nums[i++]
  while j <= right: temp[k++] = nums[j++]

  // 复制回原数组
  for k in [0, right-left]: nums[left+k] = temp[k]

  return count
```

### 图解示例

```
prices = [7,5,6,4]

归并排序过程：

         [7, 5, 6, 4]
        /            \
    [7, 5]          [6, 4]
    /    \          /    \
  [7]   [5]       [6]   [4]

合并 [7] 和 [5]: 7>5 → count=1, 结果[5,7]
合并 [6] 和 [4]: 6>4 → count=1, 结果[4,6]

合并 [5,7] 和 [4,6]:
  i=0(5), j=0(4): 5>4 → count+=2-0=2, count=4, 取 4
  i=0(5), j=1(6): 5≤6, 取 5
  i=1(7), j=1(6): 7>6 → count+=2-1=1, count=5, 取 6
  i=1(7), j=2: 左半剩余, 取 7
  结果 [4,5,6,7]

总逆序对 = 5 ✓
```

### 复杂度分析

- **时间复杂度**：O(n log n)，归并排序
- **空间复杂度**：O(n)，临时数组

## 代码

```java
class Solution {
    public int reversePairs(int[] nums) {
        if (nums == null || nums.length < 2) return 0;
        int[] temp = new int[nums.length];
        return mergeSort(nums, temp, 0, nums.length - 1);
    }

    private int mergeSort(int[] nums, int[] temp, int left, int right) {
        if (left >= right) return 0;

        int mid = left + (right - left) / 2;
        int count = mergeSort(nums, temp, left, mid)
                  + mergeSort(nums, temp, mid + 1, right);

        int i = left, j = mid + 1, k = 0;
        while (i <= mid && j <= right) {
            if (nums[i] <= nums[j]) {
                temp[k++] = nums[i++];
            } else {
                count += mid - i + 1;
                temp[k++] = nums[j++];
            }
        }

        while (i <= mid) temp[k++] = nums[i++];
        while (j <= right) temp[k++] = nums[j++];

        System.arraycopy(temp, 0, nums, left, k);

        return count;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 当 `nums[i] > nums[j]` 时计数加 `mid-i+1` | 左半部分从 i 到 mid 的所有元素都大于 nums[j] |
| 相等时不计数 | `nums[i] <= nums[j]` 不是逆序对，先取左边的保持稳定性 |
| count 可能超过 int 范围 | 当 n=50000 时，最大逆序对数为 n(n-1)/2 ≈ 1.25×10^9，int 可容纳 |
| 原数组会被修改 | 归并排序会原地排序数组，如果需要保留原数组需先复制 |
| 与 LeetCode 493 的区别 | 本题是严格逆序对（>），不是重要逆序对（> 2×） |
