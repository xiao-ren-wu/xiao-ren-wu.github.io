# 删除无效的括号

## 题目描述

给你一个由若干括号和字母组成的字符串 `s`，删除最小数量的无效括号，使得输入的字符串有效。返回所有可能的结果。

答案可以按 **任意顺序** 返回。

**示例 1：**

```
输入：s = "()())()"
输出：["(())()","()()()"]
```

**示例 2：**

```
输入：s = "(a)())()"
输出：["(a())()","(a)()()"]
```

**示例 3：**

```
输入：s = ")("
输出：[""]
```

**提示：**

- `1 <= s.length <= 25`
- `s` 由小写英文字母和 `'('`、`')'` 组成
- 字符串中最多含 `20` 个括号

## 解题思路

### 核心思想

#### 第一步：确定最少需要删除的括号数量

遍历字符串，统计无效的左括号和右括号数量：

```
leftRemove = 需要删除的左括号数
rightRemove = 需要删除的右括号数

遍历规则：
  遇到 '(' → leftRemove++
  遇到 ')' → if leftRemove > 0 then leftRemove-- else rightRemove++
```

#### 第二步：DFS 回溯生成所有有效结果

从第一个字符开始，尝试两种选择：保留或删除（如果是括号且还有删除配额）。

```
void dfs(index, leftCount, rightCount, leftRemove, rightRemove, path):
  if index == s.length():
    if leftRemove==0 && rightRemove==0:
      将 path 加入结果
    return

  // 尝试删除当前字符（如果是括号且有配额）
  if s[index] == '(' && leftRemove > 0:
    dfs(index+1, leftCount, rightCount, leftRemove-1, rightRemove, path)
  if s[index] == ')' && rightRemove > 0:
    dfs(index+1, leftCount, rightCount, leftRemove, rightRemove-1, path)

  // 保留当前字符
  path.append(s[index])
  if s[index] != '(' && s[index] != ')':
    dfs(index+1, leftCount, rightCount, leftRemove, rightRemove, path)
  else if s[index] == '(':
    dfs(index+1, leftCount+1, rightCount, leftRemove, rightRemove, path)
  else if s[index] == ')' && leftCount > rightCount:
    dfs(index+1, leftCount, rightCount+1, leftRemove, rightRemove, path)
  path.removeLast()
```

### 图解示例

```
s = "()())()"

第一步：计算删除数量
  ( → leftRemove=1
  ) → leftRemove=0
  ( → leftRemove=1
  ) → leftRemove=0
  ) → rightRemove=1
  ( → leftRemove=1
  ) → leftRemove=0

删除配额：leftRemove=0, rightRemove=1

第二步：DFS 回溯，删除 1 个多余的右括号

结果：
  "(())()" ← 删除索引 3 的 ')'
  "()()()" ← 删除索引 4 的 ')'
```

### 复杂度分析

- **时间复杂度**：O(2^n)，n 为括号数量（≤ 20），实际剪枝后远小于 2^n
- **空间复杂度**：O(n)，递归深度

## 代码

```java
class Solution {
    public List<String> removeInvalidParentheses(String s) {
        int leftRemove = 0, rightRemove = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') {
                leftRemove++;
            } else if (c == ')') {
                if (leftRemove > 0) leftRemove--;
                else rightRemove++;
            }
        }

        Set<String> result = new HashSet<>();
        dfs(s, 0, 0, 0, leftRemove, rightRemove, new StringBuilder(), result);
        return new ArrayList<>(result);
    }

    private void dfs(String s, int index, int leftCount, int rightCount,
            int leftRemove, int rightRemove, StringBuilder path,
            Set<String> result) {
        if (index == s.length()) {
            if (leftRemove == 0 && rightRemove == 0) {
                result.add(path.toString());
            }
            return;
        }

        char c = s.charAt(index);
        int len = path.length();

        if (c == '(' && leftRemove > 0) {
            dfs(s, index + 1, leftCount, rightCount,
                leftRemove - 1, rightRemove, path, result);
        }
        if (c == ')' && rightRemove > 0) {
            dfs(s, index + 1, leftCount, rightCount,
                leftRemove, rightRemove - 1, path, result);
        }

        path.append(c);
        if (c != '(' && c != ')') {
            dfs(s, index + 1, leftCount, rightCount,
                leftRemove, rightRemove, path, result);
        } else if (c == '(') {
            dfs(s, index + 1, leftCount + 1, rightCount,
                leftRemove, rightRemove, path, result);
        } else if (rightCount < leftCount) {
            dfs(s, index + 1, leftCount, rightCount + 1,
                leftRemove, rightRemove, path, result);
        }
        path.setLength(len);
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 用 Set 去重 | 同一个结果可能通过不同的删除顺序产生 |
| 保留字母字符 | 字母不能删除，直接保留并继续 |
| 右括号不能多于左括号 | 在保留 ')' 时检查 rightCount < leftCount，保证前缀始终有效 |
| 删除配额递减 | 当删除一个括号时，对应的 leftRemove/rightRemove 减 1 |
| 回溯恢复 StringBuilder | path.setLength(len) 恢复状态 |
