# 全排列

## 题目描述

给定一个不含重复数字的数组 `nums`，返回其所有可能的全排列。

**示例 1：**

```
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**示例 2：**

```
输入：nums = [0,1]
输出：[[0,1],[1,0]]
```

**示例 3：**

```
输入：nums = [1]
输出：[[1]]
```

**提示：**

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- `nums` 中的整数互不相同

## 解题思路

### 核心思想：回溯

用 `used` 数组标记已使用的元素，每次选择一个未使用的元素加入路径。

```
function backtrack(path, used):
  if path.size == nums.length:
    result.add(new ArrayList(path))
    return
  for i in [0, n-1]:
    if used[i]: continue
    used[i] = true
    path.add(nums[i])
    backtrack(path, used)
    path.removeLast()
    used[i] = false
```

### 复杂度分析

- **时间复杂度**：O(n × n!)，n! 种排列，每种复制 O(n)
- **空间复杂度**：O(n)，递归深度和 used 数组

## 代码

```java
class Solution {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        boolean[] used = new boolean[nums.length];
        backtrack(nums, used, new ArrayList<>(), res);
        return res;
    }

    private void backtrack(int[] nums, boolean[] used,
            List<Integer> path, List<List<Integer>> res) {
        if (path.size() == nums.length) {
            res.add(new ArrayList<>(path));
            return;
        }

        for (int i = 0; i < nums.length; i++) {
            if (used[i]) continue;
            used[i] = true;
            path.add(nums[i]);
            backtrack(nums, used, path, res);
            path.remove(path.size() - 1);
            used[i] = false;
        }
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| used 数组标记 | 由于元素互不相同，无需跳过重复值 |
| 结果复制 | path 是复用对象，加结果时要 new ArrayList |
| 回溯恢复状态 | 删除末尾元素且重置 used |
