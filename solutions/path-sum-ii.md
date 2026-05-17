# 路径总和 II

## 题目描述

给你二叉树的根节点 `root` 和一个整数目标和 `targetSum`，找出所有从根节点到叶子节点路径总和等于给定目标和的路径。

**示例 1：**

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
```

**示例 2：**

```
输入：root = [1,2,3], targetSum = 5
输出：[]
```

**示例 3：**

```
输入：root = [1,2], targetSum = 0
输出：[]
```

## 解题思路

### 核心思想：DFS + 回溯

从根节点开始 DFS，记录路径和路径和。到达叶子节点时判断是否等于 targetSum。

```
function dfs(root, target, path, res):
  if root == null: return

  path.add(root.val)
  target -= root.val

  if root.left == null && root.right == null && target == 0:
    res.add(new ArrayList(path))

  dfs(root.left, target, path, res)
  dfs(root.right, target, path, res)

  path.remove(path.size() - 1)  // 回溯
```

### 复杂度分析

- **时间复杂度**：O(n²)，最坏情况下复制路径 O(n)
- **空间复杂度**：O(n)，递归深度和路径存储

## 代码

```java
class Solution {
    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {
        List<List<Integer>> res = new ArrayList<>();
        dfs(root, targetSum, new ArrayList<>(), res);
        return res;
    }

    private void dfs(TreeNode root, int remain,
            List<Integer> path, List<List<Integer>> res) {
        if (root == null) return;

        path.add(root.val);
        remain -= root.val;

        if (root.left == null && root.right == null && remain == 0) {
            res.add(new ArrayList<>(path));
        }

        dfs(root.left, remain, path, res);
        dfs(root.right, remain, path, res);

        path.remove(path.size() - 1);
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 必须到叶子节点 | 路径必须从根到叶子，不含非叶子节点 |
| 回溯恢复 | 递归结束后删除当前节点值 |
| 结果复制 | 添加结果时必须 new ArrayList，path 是复用的 |
