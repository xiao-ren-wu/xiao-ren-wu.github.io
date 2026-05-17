# 螺旋矩阵

## 题目描述

给你一个 `m` 行 `n` 列的矩阵 `matrix`，请按照顺时针螺旋顺序，返回矩阵中的所有元素。

**示例 1：**

```
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
```

**示例 2：**

```
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
```

## 解题思路

### 核心思想：边界收缩

用四个变量标记上下左右边界，按顺时针方向遍历，每遍历完一条边就收缩边界。

```
top=0, bottom=m-1, left=0, right=n-1

while top <= bottom && left <= right:
  从左到右遍历上边（top++）
  从上到下遍历右边（right--）
  if top <= bottom:
    从右到左遍历下边（bottom--）
  if left <= right:
    从下到上遍历左边（left++）
```

### 复杂度分析

- **时间复杂度**：O(mn)
- **空间复杂度**：O(1)（不计输出空间）

## 代码

```java
class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res = new ArrayList<>();
        int top = 0, bottom = matrix.length - 1;
        int left = 0, right = matrix[0].length - 1;

        while (top <= bottom && left <= right) {
            for (int i = left; i <= right; i++) res.add(matrix[top][i]);
            top++;

            for (int i = top; i <= bottom; i++) res.add(matrix[i][right]);
            right--;

            if (top <= bottom) {
                for (int i = right; i >= left; i--) res.add(matrix[bottom][i]);
                bottom--;
            }

            if (left <= right) {
                for (int i = bottom; i >= top; i--) res.add(matrix[i][left]);
                left++;
            }
        }

        return res;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 下边和左边要检查边界 | 收缩后可能 top>bottom 或 left>right，需检查 |
| 单行或单列 | 处理时注意不要重复遍历 |
| 非方阵 | m ≠ n 时也能正确处理 |
