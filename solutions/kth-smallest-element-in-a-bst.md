# 二叉搜索树中第 K 小的元素

## 题目描述

给定一个二叉搜索树的根节点 `root`，和一个整数 `k`，请你设计一个算法查找其中第 `k` 小的元素（从 1 开始计数）。

**示例 1：**

```
输入：root = [3,1,4,null,2], k = 1
输出：1
```

**示例 2：**

```
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3
```

## 解题思路

### 核心思想：中序遍历

BST 的中序遍历得到升序序列，第 k 个元素即为第 k 小。

```
// 递归
count = k
result = null
function inorder(root):
  if root == null || found: return
  inorder(root.left)
  count--
  if count == 0: result = root.val; return
  inorder(root.right)
```

### 复杂度分析

- **时间复杂度**：O(H + k)，H 为树高
- **空间复杂度**：O(H)，递归栈空间

## 代码

```java
class Solution {
    private int count;
    private int result;

    public int kthSmallest(TreeNode root, int k) {
        count = k;
        inorder(root);
        return result;
    }

    private void inorder(TreeNode root) {
        if (root == null) return;

        inorder(root.left);

        count--;
        if (count == 0) {
            result = root.val;
            return;
        }

        inorder(root.right);
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 从 1 开始计数 | k=1 时第一个元素即为第 1 小 |
| 提早终止 | 找到后可通过全局变量停止后续遍历 |
| 中序是升序 | BST 中序遍历天然有序 |
