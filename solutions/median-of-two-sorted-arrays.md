# 寻找两个正序数组的中位数

## 题目描述

给定两个大小分别为 `m` 和 `n` 的正序（从小到大）数组 `nums1` 和 `nums2`。请你找出并返回这两个正序数组的 **中位数**。

算法的时间复杂度应该为 `O(log (m+n))`。

**示例 1：**

```
输入：nums1 = [1,3], nums2 = [2]
输出：2.0
解释：合并数组 = [1,2,3] ，中位数 2
```

**示例 2：**

```
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.5
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
```

**提示：**

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-10^6 <= nums1[i], nums2[i] <= 10^6`

---

## 解题思路

### 核心思想：二分查找分割法

中位数的本质：**将一组数分成左右两个数量相等的部分，且左边所有数 ≤ 右边所有数**。

对于两个有序数组，我们在两个数组上各画一条分割线，将数组分成左右两部分：

```
nums1: [a0, a1, ..., a(i-1)] | [ai, ..., a(m-1)]
         ← i 个元素 →          ← m-i 个 →

nums2: [b0, b1, ..., b(j-1)] | [bj, ..., b(n-1)]
         ← j 个元素 →          ← n-j 个 →
```

合并后的左半部分 = `nums1[0..i-1]` + `nums2[0..j-1]`，共 `i + j` 个元素。
合并后的右半部分 = `nums1[i..m-1]` + `nums2[j..n-1]`，共 `(m-i) + (n-j)` 个元素。

### 分割条件

我们需要找到一种分割，使得：

1. **左右元素数量相等**（或左边多一个）：
   
   - `i + j = (m + n + 1) / 2`（向上取整）

2. **左边所有元素 ≤ 右边所有元素**，等价于：
   
   - `nums1[i-1] ≤ nums2[j]`（nums1 左最大 ≤ nums2 右最小）
   - `nums2[j-1] ≤ nums1[i]`（nums2 左最大 ≤ nums1 右最小）

### 为什么需要交叉检查两个条件？

合并后的左边 = nums1 左半部分 + nums2 左半部分，最大值来自这两者中的较大者。
合并后的右边 = nums1 右半部分 + nums2 右半部分，最小值来自这两者中的较小者。

需要保证左边任意元素 ≤ 右边任意元素，所以既要检查 nums1 左边与 nums2 右边的交界，也要检查 nums2 左边与 nums1 右边的交界。

### 二分搜索过程

已知 `j = totalLeft - i`，问题简化为在 `[0, m]` 范围内二分搜索合适的 `i`：

- **命中了**：`nums1[i-1] ≤ nums2[j]` 且 `nums2[j-1] ≤ nums1[i]`
  
  - 奇数个元素：中位数 = `max(nums1[i-1], nums2[j-1])`
  - 偶数个元素：中位数 = `(max(左) + min(右)) / 2`

- **i 太大**（`nums1[i-1] > nums2[j]`）：nums1 左边拿太多了，`i` 左移

- **i 太小**（`nums2[j-1] > nums1[i]`）：nums1 左边拿太少了，`i` 右移

### 边界处理（哨兵值）

| 边界情况    | 问题                           | 哨兵方案                           |
| ------- | ---------------------------- | ------------------------------ |
| `i = 0` | nums1 左边无元素，`nums1[i-1]` 不存在 | `leftMax1 = MIN_VALUE`，条件自动成立  |
| `i = m` | nums1 右边无元素，`nums1[i]` 不存在   | `rightMin1 = MAX_VALUE`，条件自动成立 |
| `j = 0` | nums2 左边无元素，`nums2[j-1]` 不存在 | `leftMax2 = MIN_VALUE`，条件自动成立  |
| `j = n` | nums2 右边无元素，`nums2[j]` 不存在   | `rightMin2 = MAX_VALUE`，条件自动成立 |

### 为什么选择较短的数组进行二分搜索？

`j = totalLeft - i` 必须满足 `0 ≤ j ≤ n`，即：

- `i = 0` 时：`j = totalLeft ≤ n` → 需要 `totalLeft ≤ n`
- `i = m` 时：`j = totalLeft - m ≥ 0` → 需要 `totalLeft ≥ m`

因为 `totalLeft = (m+n+1)/2`，所以需要 `m ≤ totalLeft`，即 `m ≤ n`。

**保证始终对较短的数组做二分搜索**，这样 `i` 的搜索范围更小，时间复杂度为 `O(log(min(m, n)))`。

### 复杂度分析

- **时间复杂度**：`O(log(min(m, n)))` — 对较短的数组进行二分搜索
- **空间复杂度**：`O(1)` — 只使用了常数个变量

---

## 图解示例

```
nums1 = [1, 2]     m = 2
nums2 = [3, 4, 5]  n = 3

totalLeft = (2+3+1)/2 = 3
```

**第一次迭代：** `i = 1, j = 2`

```
nums1: [1] | [2]
nums2: [3, 4] | [5]

leftMax1 = 1, rightMin2 = 5      1 ≤ 5  ✅
leftMax2 = 4, rightMin1 = 2      4 ≤ 2  ❌
→ i 太小，右移
```

**第二次迭代：** `i = 2, j = 1`

```
nums1: [1, 2] | []
nums2: [3] | [4, 5]

leftMax1 = 2, rightMin2 = 4      2 ≤ 4  ✅
leftMax2 = 3, rightMin1 = MAX    3 ≤ MAX  ✅
→ 命中！

中位数 = max(leftMax1, leftMax2) = max(2, 3) = 3
```

---

## 最终代码

```java
/**
 * Median of Two Sorted Arrays — O(log(min(m, n)))
 *
 * Binary search the smaller array to find a partition where all left elements
 * are ≤ all right elements. Uses sentinel values (MIN_VALUE/MAX_VALUE) to
 * handle boundary cases uniformly.
 */
public class MedianOfTwoSortedArrays {

    public static double findMedianSortedArrays(int[] nums1, int[] nums2) {
        if (nums1.length > nums2.length) {
            return findMedianSortedArrays(nums2, nums1);
        }

        int m = nums1.length;
        int n = nums2.length;
        int totalLeft = (m + n + 1) / 2;

        int lo = 0;
        int hi = m;

        while (lo <= hi) {
            int i = lo + (hi - lo) / 2;
            int j = totalLeft - i;

            int leftMax1  = (i == 0)   ? Integer.MIN_VALUE : nums1[i - 1];
            int rightMin1 = (i == m)   ? Integer.MAX_VALUE : nums1[i];
            int leftMax2  = (j == 0)   ? Integer.MIN_VALUE : nums2[j - 1];
            int rightMin2 = (j == n)   ? Integer.MAX_VALUE : nums2[j];

            if (leftMax1 <= rightMin2 && leftMax2 <= rightMin1) {
                if ((m + n) % 2 == 1) {
                    return Math.max(leftMax1, leftMax2);
                } else {
                    return (Math.max(leftMax1, leftMax2) + Math.min(rightMin1, rightMin2)) / 2.0;
                }
            } else if (leftMax1 > rightMin2) {
                hi = i - 1;
            } else {
                lo = i + 1;
            }
        }

        throw new IllegalArgumentException("Input arrays are not sorted");
    }

    public static void main(String[] args) {
        System.out.println("Example 1: "
            + findMedianSortedArrays(new int[]{1, 3}, new int[]{2})
            + " (expected 2.0)");
        System.out.println("Example 2: "
            + findMedianSortedArrays(new int[]{1, 2}, new int[]{3, 4})
            + " (expected 2.5)");
        System.out.println("Edge (one empty): "
            + findMedianSortedArrays(new int[]{}, new int[]{1})
            + " (expected 1.0)");
        System.out.println("Edge (all same): "
            + findMedianSortedArrays(new int[]{1, 1}, new int[]{1, 1})
            + " (expected 1.0)");
        System.out.println("Edge (negative): "
            + findMedianSortedArrays(new int[]{-5, -3, -1}, new int[]{-4, -2, 0})
            + " (expected -2.5)");
    }
}
```

---

## 常见错误总结

| 错误                                   | 后果                | 修正                 |
| ------------------------------------ | ----------------- | ------------------ |
| 哨兵值用反：`j == n` 时给 `MIN_VALUE`        | 分割条件永远不满足，搜索乱掉    | 给 `MAX_VALUE`      |
| 二分边界用 `right = i` 而非 `right = i - 1` | i 被重复判定，可能导致死循环   | 用 `right = i - 1`  |
| 不交换数组：没有保证 `m ≤ n`                   | `j` 可能越界（负数或超过 n） | 始终让 `nums1` 是较短的数组 |

---
