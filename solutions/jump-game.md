# 跳跃游戏

## 题目描述

给你一个非负整数数组 `nums`，你最初位于数组的第一个下标。数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个下标。

**示例 1：**

```
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步到下标 1，再跳 3 步到达最后一个下标。
```

**示例 2：**

```
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标 3 的位置。该位置最大跳跃长度为 0，无法到达终点。
```

## 解题思路

### 核心思想：贪心

维护能到达的最远位置，遍历数组，如果当前位置不可达则返回 false，否则更新最远位置。

```
reachable = 0
for i in [0, n-1]:
  if i > reachable: return false
  reachable = max(reachable, i + nums[i])
return true
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public boolean canJump(int[] nums) {
        int reachable = 0;

        for (int i = 0; i < nums.length; i++) {
            if (i > reachable) return false;
            reachable = Math.max(reachable, i + nums[i]);
        }

        return true;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 提前终止 | `i > reachable` 说明当前位置不可达，直接返回 false |
| 不需要记录跳跃次数 | 只需判断能否到达，不同于跳跃游戏 II |
| 最远距离单调不降 | reachable 只会增大，不会减小 |
