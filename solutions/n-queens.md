# N 皇后

## 题目描述

按照国际象棋的规则，皇后可以攻击同一行、同一列、同一斜线上的棋子。`n` 皇后问题研究的是如何将 `n` 个皇后放置在 `n × n` 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 `n`，返回所有不同的 `n` 皇后问题的解决方案。

每一种解法包含一个不同的 `n` 皇后问题的棋子放置方案，该方案中 `'Q'` 和 `'.'` 分别代表了皇后和空位。

**示例 1：**

```
输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：4 皇后问题有两个不同的解法。
```

**示例 2：**

```
输入：n = 1
输出：[["Q"]]
```

**提示：**

- `1 <= n <= 9`

## 解题思路

### 核心思想：回溯法

逐行放置皇后，每行只能放一个。在放置时检查三个约束条件：

1. **列** — 该列不能有其它皇后
2. **主对角线（\）** — 行列差 `row - col` 为常数，不能重复
3. **副对角线（/）** — 行列和 `row + col` 为常数，不能重复

### 算法步骤

```
函数 backtrack(row):
  如果 row == n: 找到一个解，保存后返回
  
  遍历 col in [0, n):
    检查 cols[row-col] 和 diags1[row-col] 和 diags2[row+col] 是否被占用
    
    如果可用：
      放置皇后，标记三个集合
      backtrack(row + 1)
      回溯：撤销标记
```

### 图解示例

```
n = 4 的一个解：

行0:  .  Q  .  .      cols[1] 占用
行1:  .  .  .  Q      cols[3] 占用
行2:  Q  .  .  .      cols[0] 占用
行3:  .  .  Q  .      cols[2] 占用

主对角线差 (row-col)：
  (0,1): -1  (1,3): -2  (2,0): 2  (3,2): 1  → 全不同 ✓

副对角线和 (row+col)：
  (0,1): 1   (1,3): 4   (2,0): 2  (3,2): 5  → 全不同 ✓
```

### 复杂度分析

- **时间复杂度**：O(n!)，最坏情况需要探索所有排列，但剪枝大大减少搜索空间
- **空间复杂度**：O(n)，三个集合和递归栈深度

## 代码

```java
class Solution {
    public List<List<String>> solveNQueens(int n) {
        List<List<String>> res = new ArrayList<>();
        boolean[] cols = new boolean[n];
        boolean[] diag1 = new boolean[2 * n - 1]; // row - col + n - 1
        boolean[] diag2 = new boolean[2 * n - 1]; // row + col
        char[][] board = new char[n][n];
        for (int i = 0; i < n; i++) {
            Arrays.fill(board[i], '.');
        }
        backtrack(res, board, cols, diag1, diag2, 0, n);
        return res;
    }

    private void backtrack(List<List<String>> res, char[][] board,
            boolean[] cols, boolean[] diag1, boolean[] diag2,
            int row, int n) {
        if (row == n) {
            List<String> snapshot = new ArrayList<>();
            for (char[] r : board) {
                snapshot.add(new String(r));
            }
            res.add(snapshot);
            return;
        }

        for (int col = 0; col < n; col++) {
            int d1 = row - col + n - 1;
            int d2 = row + col;
            if (cols[col] || diag1[d1] || diag2[d2]) continue;

            board[row][col] = 'Q';
            cols[col] = diag1[d1] = diag2[d2] = true;
            backtrack(res, board, cols, diag1, diag2, row + 1, n);
            cols[col] = diag1[d1] = diag2[d2] = false;
            board[row][col] = '.';
        }
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 对角线用数组而非 Set | 用 `2n-1` 长度的 boolean 数组替代 Set，性能更好 |
| 主对角线索引偏移 | `row - col` 范围为 `[-(n-1), n-1]`，加 `n-1` 偏移到 `[0, 2n-2]` |
| 回溯必须还原所有标记 | 列、两条对角线和棋盘都要还原 |
| n <= 9 是有效约束 | 回溯法在 n=9 时约 37k 个解，仍在合理范围 |
