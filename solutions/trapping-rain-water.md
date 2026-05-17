# 接雨水

## 题目描述

给定 `n` 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

**示例 1：**

```
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水。
```

**示例 2：**

```
输入：height = [4,2,0,3,2,5]
输出：9
```

**提示：**

- `n == height.length`
- `1 <= n <= 2 * 10^4`
- `0 <= height[i] <= 10^5`

## 解题思路

### 核心思想

每个位置能接的雨水量取决于它左右两侧最高柱子中较矮的那个，减去当前位置的高度：

```
water[i] = min(maxLeft[i], maxRight[i]) - height[i]
```

### 解法一：双指针（最优）

用两个指针从左右向中间遍历，同时维护 `leftMax` 和 `rightMax`：

- 如果 `height[left] < height[right]`，说明左侧的瓶颈由 `leftMax` 决定，处理左指针
- 否则处理右指针

```
初始化：left = 0, right = n-1, leftMax = 0, rightMax = 0, ans = 0

while left < right:
    if height[left] < height[right]:
        if height[left] >= leftMax:
            leftMax = height[left]
        else:
            ans += leftMax - height[left]
        left++
    else:
        if height[right] >= rightMax:
            rightMax = height[right]
        else:
            ans += rightMax - height[right]
        right--
```

### 图解示例

```
height = [0,1,0,2,1,0,1,3,2,1,2,1]

雨水累积过程：
  ↓                     ↓
  0 1 0 2 1 0 1 3 2 1 2 1
  leftMax=0 rightMax=1 → leftMax太小，左移

    ↓                   ↓
  0 1 0 2 1 0 1 3 2 1 2 1
  leftMax=1 rightMax=1 → 高度1，无积水，左移

      ↓                 ↓
  0 1 0 2 1 0 1 3 2 1 2 1
  leftMax=1 rightMax=1 → water=1-0=1，左移

...最终累计得到 6
```

### 复杂度分析

| 解法 | 时间复杂度 | 空间复杂度 |
|------|-----------|-----------|
| 双指针 | O(n) | O(1) |
| DP（左右数组） | O(n) | O(n) |
| 单调栈 | O(n) | O(n) |

## 代码

```java
class Solution {
    public int trap(int[] height) {
        int left = 0, right = height.length - 1;
        int leftMax = 0, rightMax = 0;
        int ans = 0;

        while (left < right) {
            if (height[left] < height[right]) {
                if (height[left] >= leftMax) {
                    leftMax = height[left];
                } else {
                    ans += leftMax - height[left];
                }
                left++;
            } else {
                if (height[right] >= rightMax) {
                    rightMax = height[right];
                } else {
                    ans += rightMax - height[right];
                }
                right--;
            }
        }

        return ans;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 双指针只在矮的一侧计算 | 因为较矮的一侧已经确定了瓶颈高度，高的一侧即使有更高的柱子也不影响 |
| 左右指针相遇时停止 | 最后一个位置不需要处理（无积水） |
| leftMax/rightMax 实时更新 | 不是整个数组的全局最大，而是当前指针遍历过的最大 |
| 单调栈解法核心 | 栈底到栈顶递减，遇到更高的元素时弹栈计算面积 |
