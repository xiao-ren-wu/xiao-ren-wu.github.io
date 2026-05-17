# 单词搜索

## 题目描述

给定一个 `m × n` 二维字符网格 `board` 和一个字符串单词 `word`。如果 word 存在于网格中，返回 true；否则返回 false。单词必须按字母顺序，通过相邻的单元格（水平或垂直）构成。同一个单元格的字母不允许被重复使用。

**示例 1：**

```
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCCED'
输出：true
```

**示例 2：**

```
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'SEE'
输出：true
```

**示例 3：**

```
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCB'
输出：false
```

## 解题思路

### 核心思想：DFS + 回溯

从每个格子出发，尝试匹配 word。用临时标记避免重复使用。

```
function dfs(board, i, j, word, index):
  if index == word.length: return true
  if 越界 或 board[i][j] != word[index]: return false

  temp = board[i][j]
  board[i][j] = '#'  标记已访问
  found = dfs(i+1,j) || dfs(i-1,j) || dfs(i,j+1) || dfs(i,j-1)
  board[i][j] = temp  恢复
  return found
```

### 复杂度分析

- **时间复杂度**：O(mn × 3^k)，k 为 word 长度（首字母有 4 个方向，后续最多 3 个）
- **空间复杂度**：O(k)，递归深度

## 代码

```java
class Solution {
    public boolean exist(char[][] board, String word) {
        int m = board.length, n = board[0].length;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (dfs(board, i, j, word, 0)) return true;
            }
        }

        return false;
    }

    private boolean dfs(char[][] board, int i, int j, String word, int index) {
        if (index == word.length()) return true;
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length
                || board[i][j] != word.charAt(index)) return false;

        char tmp = board[i][j];
        board[i][j] = '#';

        boolean found = dfs(board, i + 1, j, word, index + 1)
                     || dfs(board, i - 1, j, word, index + 1)
                     || dfs(board, i, j + 1, word, index + 1)
                     || dfs(board, i, j - 1, word, index + 1);

        board[i][j] = tmp;
        return found;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 标记已访问 | 用 '#' 或特殊字符标记，避免重复使用 |
| 回溯恢复 | 递归结束后必须恢复原字符 |
| 剪枝优化 | 可先统计 board 字符频率，若 word 中存在超出频率的字符直接 false |
