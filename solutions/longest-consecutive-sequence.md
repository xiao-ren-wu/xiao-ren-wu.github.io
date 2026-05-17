# 最长连续序列

## 题目描述

给定一个未排序的整数数组 `nums`，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。要求 O(n)。

**示例 1：**

```
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1,2,3,4]，长度为 4。
```

**示例 2：**

```
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
```

## 解题思路

### 核心思想：哈希集合

将所有数放入 HashSet。遍历每个数，只有当它是某个连续序列的起点时（即 num-1 不在集合中）才计数。

```
set = new HashSet(nums)
maxLen = 0

for num in set:
  if !set.contains(num - 1):  // 是序列起点
    cur = num
    len = 1
    while set.contains(cur + 1):
      cur++; len++
    maxLen = max(maxLen, len)

return maxLen
```

### 复杂度分析

- **时间复杂度**：O(n)，每个数只被访问一次（跳过非起点）
- **空间复杂度**：O(n)

## 代码

```java
class Solution {
    public int longestConsecutive(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for (int num : nums) set.add(num);

        int maxLen = 0;

        for (int num : set) {
            if (!set.contains(num - 1)) {
                int cur = num;
                int len = 1;
                while (set.contains(cur + 1)) {
                    cur++;
                    len++;
                }
                maxLen = Math.max(maxLen, len);
            }
        }

        return maxLen;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 只从起点计数 | `!set.contains(num-1)` 确保每个序列只计数一次 |
| 遍历 set 而非数组 | 避免重复元素导致的重复计算 |
| 空间换时间 | O(n) 空间换取 O(n) 时间 |
