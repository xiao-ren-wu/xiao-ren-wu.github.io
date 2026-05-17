# 单词拆分

## 题目描述

给你一个字符串 `s` 和一个字符串列表 `wordDict` 作为字典。请判断是否可以利用字典中出现的单词拼接出 s。字典中的单词可以重复使用。

**示例 1：**

```
输入：s = "leetcode", wordDict = ["leet","code"]
输出：true
```

**示例 2：**

```
输入：s = "applepenapple", wordDict = ["apple","pen"]
输出：true
```

**示例 3：**

```
输入：s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
输出：false
```

## 解题思路

### 核心思想：动态规划

`dp[i]` 表示 s 的前 i 个字符能否被拆分为字典中的单词。

```
dp[0] = true（空串）
for i in [1, n]:
  for j in [0, i-1]:
    if dp[j] && s[j:i] in wordDict:
      dp[i] = true; break
```

### 复杂度分析

- **时间复杂度**：O(n²)，n 为字符串长度
- **空间复杂度**：O(n)

## 代码

```java
class Solution {
    public boolean wordBreak(String s, List<String> wordDict) {
        Set<String> set = new HashSet<>(wordDict);
        boolean[] dp = new boolean[s.length() + 1];
        dp[0] = true;

        for (int i = 1; i <= s.length(); i++) {
            for (int j = 0; j < i; j++) {
                if (dp[j] && set.contains(s.substring(j, i))) {
                    dp[i] = true;
                    break;
                }
            }
        }

        return dp[s.length()];
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| wordDict 转 Set | O(1) 查找字典单词 |
| dp[0] = true | 空串可拼接，作为递推基础 |
| 剪枝 break | 找到一种拆分即可跳出内层循环 |
