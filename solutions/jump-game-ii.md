# 跳跃游戏 II

## 题目描述

给你一个非负整数数组 `nums`，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。你的目标是使用最少的跳跃次数到达数组的最后一个位置。

**示例 1：**

```
输入：nums = [2,3,1,1,4]
输出：2
解释：跳到下标 1 的位置（跳 1 步），再跳 3 步到达最后一个位置。
```

**示例 2：**

```
输入：nums = [2,3,0,1,4]
输出：2
```

**提示：**

- `1 <= nums.length <= 10^4`
- `0 <= nums[i] <= 1000`

## 解题思路

### 核心思想：贪心 + BFS

维护当前跳跃能到达的最远边界。遍历当前范围内的所有位置，更新下一步能到达的最远位置。到达边界时跳跃次数 +1。

```
jumps = 0, curEnd = 0, farthest = 0

for i in [0, n-2]:
  farthest = max(farthest, i + nums[i])
  if i == curEnd:
    jumps++
    curEnd = farthest
```

### 图解示例

```
nums = [2,3,1,1,4]

i=0: farthest=2, i==curEnd(0) → jumps=1, curEnd=2
i=1: farthest=4
i=2: farthest=4, i==curEnd(2) → jumps=2, curEnd=4

结果：2 步
```

### 复杂度分析

- **时间复杂度**：O(n)，一次遍历
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int jump(int[] nums) {
        int jumps = 0, curEnd = 0, farthest = 0;

        for (int i = 0; i < nums.length - 1; i++) {
            farthest = Math.max(farthest, i + nums[i]);
            if (i == curEnd) {
                jumps++;
                curEnd = farthest;
            }
        }

        return jumps;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 只遍历到 n-2 | 最后一个位置不需要再跳 |
| curEnd 是当前跳跃的边界 | 到达边界时必须进行一次新的跳跃（除非已在终点） |
| 题目保证可到达终点 | 不需要检查不可达的情况 |
