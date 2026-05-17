# 二叉树的最近公共祖先

## 题目描述

给定一个二叉树，找到该树中两个指定节点的最近公共祖先。

最近公共祖先的定义：对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。

**示例 1：**

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
```

**示例 2：**

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
```

**示例 3：**

```
输入：root = [1,2], p = 1, q = 2
输出：1
```

## 解题思路

### 核心思想：递归查找

从根节点 DFS，如果当前节点等于 p 或 q 则返回当前节点。分别从左子树和右子树查找 p、q 的 LCA。

```
function lowestCommonAncestor(root, p, q):
  if root == null || root == p || root == q: return root

  left = lowestCommonAncestor(root.left, p, q)
  right = lowestCommonAncestor(root.right, p, q)

  if left != null && right != null: return root  // p、q 分别在左右子树
  return left != null ? left : right
```

### 复杂度分析

- **时间复杂度**：O(n)，每个节点最多访问一次
- **空间复杂度**：O(n)，递归栈空间

## 代码

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || root == p || root == q) return root;

        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);

        if (left != null && right != null) return root;
        return left != null ? left : right;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 自己也是祖先 | 如果 p 是 q 的祖先，LCA 就是 p |
| 左右都非空 | 说明 p、q 分居两侧，root 即 LCA |
| 仅一侧非空 | 说明 p、q 都在同一侧，继续向该侧递归 |
