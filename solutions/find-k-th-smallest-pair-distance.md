# 找出第 K 小的数对距离

## 题目描述

数对 `(a, b)` 由整数 `a` 和 `b` 组成，其数对距离定义为 `a` 和 `b` 的绝对差值。

给你一个整数数组 `nums` 和一个整数 `k`，返回所有数对距离中第 `k` 小的数对距离。

**示例 1：**

```
输入：nums = [1,3,1], k = 1
输出：0
解释：数对距离排序为：|1-1|=0, |1-3|=2, |3-1|=2，最小的距离是 0。
```

**示例 2：**

```
输入：nums = [1,1,1], k = 2
输出：0
```

**示例 3：**

```
输入：nums = [1,6,1], k = 3
输出：5
```

**提示：**

- `n == nums.length`
- `2 <= n <= 10^4`
- `0 <= nums[i] <= 10^6`
- `1 <= k <= n * (n - 1) / 2`

## 解题思路

### 核心思想：二分答案 + 双指针计数

直接找第 k 小的距离很难，但可以换个思路：

> 给定一个距离 `d`，有多少对数对的距离 ≤ d？

如果能高效地回答这个问题，就可以对答案做**二分搜索**。

### 算法步骤

```
1. 排序 nums

2. 二分搜索距离 d，范围 [0, max(nums) - min(nums)]

3. 对于每个 mid：
   - 用双指针统计距离 ≤ mid 的数对个数 count
   - 如果 count >= k：说明答案 ≤ mid，缩小右边界
   - 如果 count < k：说明答案 > mid，扩大左边界

4. 返回左边界
```

### 双指针计数

```
countPairs(nums, mid):
  count = 0
  left = 0
  for right in [0..n-1]:
    while nums[right] - nums[left] > mid:
      left++
    count += right - left
  return count
```

### 图解示例

```
nums = [1,3,1] → 排序后 [1,1,3]

二分搜索距离：
  左=0, 右=2
  mid=1: 距离≤1的数对有 (1,1) 这一个，count=1, k=1 → count≥k，右移
  mid=0: 距离≤0的数对只有 (1,1) → count=1, k=1 → count≥k，右移
  左=0, 右=0 → 答案=0

统计距离≤0：
  right=0: nums[0]-nums[0]=0≤0 → count+=0-0=0
  right=1: nums[1]-nums[0]=0≤0 → count+=1-0=1
  right=2: nums[2]=3, left右移到1(3-1=2>0), 再右移(3-1=2>0) → left=2, count+=2-2=0
  total = 1 ✓
```

### 复杂度分析

- **时间复杂度**：O(n log n + n log W)，排序 O(n log n)，二分 O(log W) × 计数 O(n)
- **空间复杂度**：O(log n)，排序栈空间

## 代码

```java
class Solution {
    public int smallestDistancePair(int[] nums, int k) {
        Arrays.sort(nums);

        int left = 0;
        int right = nums[nums.length - 1] - nums[0];

        while (left < right) {
            int mid = left + (right - left) / 2;
            if (countPairs(nums, mid) >= k) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        return left;
    }

    private int countPairs(int[] nums, int maxDist) {
        int count = 0;
        int left = 0;

        for (int right = 0; right < nums.length; right++) {
            while (nums[right] - nums[left] > maxDist) {
                left++;
            }
            count += right - left;
        }

        return count;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 需要先排序 | 双指针计数依赖有序数组 |
| `count >= k` 时右边界收缩 | 不是直接返回，要找到最小的距离，而不是任意一个 |
| 计数用 `right - left` | 对于固定的 right，left 到 right-1 之间的所有数对距离都 ≤ maxDist |
| k 从 1 开始 | 第 k 小，k=1 表示最小的距离 |
| 最大距离 = max - min | 排序后最后一个元素减第一个元素 |
