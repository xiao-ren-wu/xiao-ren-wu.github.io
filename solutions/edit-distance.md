# 编辑距离

## 题目描述

给你两个单词 `word1` 和 `word2`，请返回将 `word1` 转换成 `word2` 所使用的最少操作数。可以执行三种操作：插入一个字符、删除一个字符、替换一个字符。

**示例 1：**

```
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse → rorse (替换 'h' 为 'r')
rorse → rose (删除 'r')
rose → ros (删除 'e')
```

**示例 2：**

```
输入：word1 = "intention", word2 = "execution"
输出：5
```

## 解题思路

### 核心思想：动态规划

定义 `dp[i][j]` 为 word1 前 i 个字符转为 word2 前 j 个字符的最小编辑距离。

```
if word1[i-1] == word2[j-1]:
  dp[i][j] = dp[i-1][j-1]  // 字符相同，无需操作
else:
  dp[i][j] = min(
    dp[i-1][j] + 1,      // 删除 word1[i-1]
    dp[i][j-1] + 1,      // 插入 word2[j-1]
    dp[i-1][j-1] + 1     // 替换 word1[i-1] → word2[j-1]
  )
```

### 状态转移表示例

```
word1 = "horse", word2 = "ros"

    ''  r  o  s
''   0  1  2  3
h    1  1  2  3
o    2  2  1  2
r    3  2  2  2
s    4  3  3  2
e    5  4  4  3
```

### 复杂度分析

- **时间复杂度**：O(mn)
- **空间复杂度**：O(mn)，可优化为 O(min(m,n))

## 代码

```java
class Solution {
    public int minDistance(String word1, String word2) {
        int m = word1.length(), n = word2.length();
        int[][] dp = new int[m + 1][n + 1];

        for (int i = 0; i <= m; i++) dp[i][0] = i;
        for (int j = 0; j <= n; j++) dp[0][j] = j;

        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (word1.charAt(i - 1) == word2.charAt(j - 1)) {
                    dp[i][j] = dp[i - 1][j - 1];
                } else {
                    dp[i][j] = Math.min(
                        Math.min(dp[i - 1][j] + 1, dp[i][j - 1] + 1),
                        dp[i - 1][j - 1] + 1
                    );
                }
            }
        }

        return dp[m][n];
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 初始化行列 | dp[i][0] = i 表示删除 i 次，dp[0][j] = j 表示插入 j 次 |
| 三种操作对应关系 | dp[i-1][j] 删除、dp[i][j-1] 插入、dp[i-1][j-1] 替换 |
| 字符相等时直接继承 | 不需要操作，继承左上角的值 |
