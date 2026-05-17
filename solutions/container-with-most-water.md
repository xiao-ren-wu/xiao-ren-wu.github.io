# 盛最多水的容器

## 题目描述

给定一个长度为 `n` 的整数数组 `height`，有 `n` 条垂线，第 `i` 条线的两个端点是 `(i, 0)` 和 `(i, height[i])`。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。返回容器可以储存的最大水量。

**示例 1：**

```
输入：[1,8,6,2,5,4,8,3,7]
输出：49
```

**示例 2：**

```
输入：height = [1,1]
输出：1
```

**提示：**

- `n == height.length`
- `2 <= n <= 10^5`
- `0 <= height[i] <= 10^4`

## 解题思路

### 核心思想：双指针

容器的水量 = 宽度 × 高度（两侧中较矮的）。用两个指针从两端向中间移动，每次移动较矮的那一侧。

```
left=0, right=n-1, maxArea=0

while left < right:
  area = (right-left) × min(height[left], height[right])
  maxArea = max(maxArea, area)
  if height[left] < height[right]: left++
  else: right--
```

### 复杂度分析

- **时间复杂度**：O(n)，一次遍历
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int maxArea(int[] height) {
        int left = 0, right = height.length - 1;
        int maxArea = 0;

        while (left < right) {
            int area = (right - left) * Math.min(height[left], height[right]);
            maxArea = Math.max(maxArea, area);

            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }
        }

        return maxArea;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 移动较矮的一侧 | 只有移动矮的才有可能获得更大面积，移动高的只会让宽度减小 |
| 相等时任意移动 | 可以左右都移动，也可以任选一边 |
| 宽度不断缩小 | 移动指针后宽度减 1，但高度可能变大 |
