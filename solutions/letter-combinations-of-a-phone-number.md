# 电话号码的字母组合

## 题目描述

给定一个仅包含数字 `2-9` 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。数字到字母的映射与电话按键相同。

```
2: abc    3: def    4: ghi
5: jkl    6: mno    7: pqrs
8: tuv    9: wxyz
```

**示例 1：**

```
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**示例 2：**

```
输入：digits = ""
输出：[]
```

**示例 3：**

```
输入：digits = "2"
输出：["a","b","c"]
```

**提示：**

- `0 <= digits.length <= 4`
- `digits[i]` 是范围 `['2', '9']` 的数字

## 解题思路

### 核心思想：回溯

将每个数字对应的字母作为选择列表，用回溯法枚举所有组合。

```
mapping = ["", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]

function backtrack(index, path):
  if index == digits.length():
    result.add(path)
    return

  letters = mapping[digits[index]]
  for each letter in letters:
    backtrack(index + 1, path + letter)
```

### 复杂度分析

- **时间复杂度**：O(4^n)，n 为数字个数，最坏每个数字映射 4 个字母
- **空间复杂度**：O(n)，递归栈深度

## 代码

```java
class Solution {
    private static final String[] MAPPING = {
        "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"
    };

    public List<String> letterCombinations(String digits) {
        List<String> res = new ArrayList<>();
        if (digits == null || digits.length() == 0) return res;
        backtrack(digits, 0, new StringBuilder(), res);
        return res;
    }

    private void backtrack(String digits, int index, StringBuilder path, List<String> res) {
        if (index == digits.length()) {
            res.add(path.toString());
            return;
        }

        String letters = MAPPING[digits.charAt(index) - '0'];
        for (char c : letters.toCharArray()) {
            path.append(c);
            backtrack(digits, index + 1, path, res);
            path.deleteCharAt(path.length() - 1);
        }
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 空串处理 | digits="" 返回 []，不是 [""] |
| MAPPING 中 0 和 1 为空 | 电话按键上 0 和 1 没有字母 |
| 回溯恢复状态 | 用 StringBuilder 时记得删除末尾字符 |
