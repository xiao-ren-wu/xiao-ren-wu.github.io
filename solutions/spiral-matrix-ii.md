# 螺旋矩阵 II

## 题目描述

给你一个正整数 `n`，生成一个包含 `1` 到 `n²` 所有元素，且元素按顺时针顺序螺旋排列的 `n × n` 正方形矩阵。

**示例 1：**

```
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
```

**示例 2：**

```
输入：n = 1
输出：[[1]]
```

## 解题思路

### 核心思想：边界收缩

与螺旋矩阵相同思路，用四个边界变量控制方向，依次填充数字。

```
top=0, bottom=n-1, left=0, right=n-1
num = 1

while num <= n*n:
  从左到右填充上边（top++）
  从上到下填充右边（right--）
  从右到左填充下边（bottom--）
  从下到上填充左边（left++）
```

### 复杂度分析

- **时间复杂度**：O(n²)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int[][] generateMatrix(int n) {
        int[][] matrix = new int[n][n];
        int top = 0, bottom = n - 1;
        int left = 0, right = n - 1;
        int num = 1;

        while (num <= n * n) {
            for (int i = left; i <= right; i++) matrix[top][i] = num++;
            top++;

            for (int i = top; i <= bottom; i++) matrix[i][right] = num++;
            right--;

            for (int i = right; i >= left; i--) matrix[bottom][i] = num++;
            bottom--;

            for (int i = bottom; i >= top; i--) matrix[i][left] = num++;
            left++;
        }

        return matrix;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 边界收缩时机 | 每填完一条边立即收缩，防止重复 |
| 无需检查边界 | n×n 方阵，每次循环四条边都会填满 |
| 奇数的中心元素 | 由最后一次循环的左边或上边填充 |
