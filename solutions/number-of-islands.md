# 岛屿数量

## 题目描述

给你一个由 `'1'`（陆地）和 `'0'`（水）组成的二维网格，请你计算网格中岛屿的数量。岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

**示例 1：**

```
输入：
[
  ['1','1','1','1','0'],
  ['1','1','0','1','0'],
  ['1','1','0','0','0'],
  ['0','0','0','0','0']
]
输出：1
```

**示例 2：**

```
输入：
[
  ['1','1','0','0','0'],
  ['1','1','0','0','0'],
  ['0','0','1','0','0'],
  ['0','0','0','1','1']
]
输出：3
```

## 解题思路

### 核心思想：DFS 标记

遍历网格，遇到 '1' 时计数 +1，然后 DFS 将所有相邻的 '1' 标记为 '0'（沉没）。

```
count = 0
for i in [0, m-1]:
  for j in [0, n-1]:
    if grid[i][j] == '1':
      count++
      dfs(grid, i, j)

function dfs(grid, i, j):
  如果越界或 grid[i][j] != '1': return
  grid[i][j] = '0'
  dfs(i+1, j) dfs(i-1, j) dfs(i, j+1) dfs(i, j-1)
```

### 复杂度分析

- **时间复杂度**：O(mn)
- **空间复杂度**：O(mn)，最坏情况递归深度

## 代码

```java
class Solution {
    public int numIslands(char[][] grid) {
        int count = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == '1') {
                    count++;
                    dfs(grid, i, j);
                }
            }
        }
        return count;
    }

    private void dfs(char[][] grid, int i, int j) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length
                || grid[i][j] != '1') return;

        grid[i][j] = '0';
        dfs(grid, i + 1, j);
        dfs(grid, i - 1, j);
        dfs(grid, i, j + 1);
        dfs(grid, i, j - 1);
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 沉没标记 | 将 '1' 改为 '0' 避免重复计数，节省 visited 数组空间 |
| 方向覆盖 | 四个方向都要 DFS |
| BFS 也可 | 可用队列做 BFS 替代递归，避免栈溢出 |
