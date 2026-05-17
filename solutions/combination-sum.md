# 组合总和

## 题目描述

给你一个无重复元素的整数数组 `candidates` 和一个目标整数 `target`，找出所有和为 `target` 的组合。同一个数字可以无限制重复被选取。

**示例 1：**

```
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
```

**示例 2：**

```
输入：candidates = [2,3,5], target = 8
输出：[[2,2,2,2],[2,3,3],[3,5]]
```

**示例 3：**

```
输入：candidates = [2], target = 1
输出：[]
```

**提示：**

- `1 <= candidates.length <= 30`
- `2 <= candidates[i] <= 40`
- `candidates` 所有元素互不相同

## 解题思路

### 核心思想：回溯 + 剪枝

排序后回溯。每次从当前索引开始（可重复使用同一元素），当 sum > target 时剪枝。

```
排序
function backtrack(start, target, path):
  if target == 0: result.add(path); return
  for i in [start, n-1]:
    if candidates[i] > target: break
    path.add(candidates[i])
    backtrack(i, target - candidates[i], path)
    path.removeLast()
```

### 复杂度分析

- **时间复杂度**：O(n^(target/min))，最坏情况指数级
- **空间复杂度**：O(target/min)，递归深度

## 代码

```java
class Solution {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(candidates);
        backtrack(candidates, target, 0, new ArrayList<>(), res);
        return res;
    }

    private void backtrack(int[] candidates, int remain, int start,
            List<Integer> path, List<List<Integer>> res) {
        if (remain == 0) {
            res.add(new ArrayList<>(path));
            return;
        }

        for (int i = start; i < candidates.length; i++) {
            if (candidates[i] > remain) break;
            path.add(candidates[i]);
            backtrack(candidates, remain - candidates[i], i, path, res);
            path.remove(path.size() - 1);
        }
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 排序 + `> remain` 剪枝 | 排序后一旦当前数大于剩余目标，后面的都可以跳过 |
| 递归传 i 而非 i+1 | 允许重复使用同一个元素 |
| 结果要 new ArrayList | path 是复用的，必须复制一份加入结果 |
