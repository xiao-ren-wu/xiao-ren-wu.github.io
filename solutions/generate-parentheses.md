# 括号生成

## 题目描述

数字 `n` 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且有效的括号组合。

**示例 1：**

```
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
```

**示例 2：**

```
输入：n = 1
输出：["()"]
```

**提示：**

- `1 <= n <= 8`

## 解题思路

### 核心思想：回溯 + 剪枝

在构造过程中维护左右括号的数量，保证：

1. 左括号数量 ≤ n
2. 右括号数量 ≤ 左括号数量（保证当前前缀始终有效）

```
function backtrack(left, right, path):
  if path.length == 2*n:
    result.add(path)
    return

  if left < n:
    backtrack(left+1, right, path + "(")
  if right < left:
    backtrack(left, right+1, path + ")")
```

### 复杂度分析

- **时间复杂度**：O(4^n / √n)，第 n 个卡特兰数
- **空间复杂度**：O(n)，递归深度

## 代码

```java
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<>();
        backtrack(n, 0, 0, new StringBuilder(), res);
        return res;
    }

    private void backtrack(int n, int left, int right, StringBuilder path, List<String> res) {
        if (path.length() == 2 * n) {
            res.add(path.toString());
            return;
        }

        if (left < n) {
            path.append('(');
            backtrack(n, left + 1, right, path, res);
            path.deleteCharAt(path.length() - 1);
        }
        if (right < left) {
            path.append(')');
            backtrack(n, left, right + 1, path, res);
            path.deleteCharAt(path.length() - 1);
        }
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 右括号条件 `right < left` | 不是 `right < n`，必须保证当前前缀中左括号比右括号多 |
| 只有在左括号未用完时才添加 `(` | n 对括号最多 n 个左括号 |
| 回溯恢复 StringBuilder | 删除末尾字符 |
