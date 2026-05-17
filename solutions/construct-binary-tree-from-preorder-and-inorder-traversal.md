# 从前序与中序遍历序列构造二叉树

## 题目描述

给定两个整数数组 `preorder` 和 `inorder`，其中 `preorder` 是二叉树的先序遍历，`inorder` 是同一棵树的中序遍历，请构造二叉树并返回根节点。

**示例 1：**

```
输入：preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出：[3,9,20,null,null,15,7]
```

**示例 2：**

```
输入：preorder = [-1], inorder = [-1]
输出：[-1]
```

## 解题思路

### 核心思想：递归 + 哈希优化

前序的第一个元素是根节点，在中序中找到根节点位置，左边是左子树，右边是右子树。

```
preorder = [根 | 左子树 | 右子树]
inorder  = [左子树 | 根 | 右子树]

用 HashMap 存 inorder 中值→索引的映射，O(1) 查找。
```

```
function build(preorder, preStart, preEnd, inorder, inStart, inEnd):
  if preStart > preEnd: return null

  rootVal = preorder[preStart]
  root = new TreeNode(rootVal)
  rootIndex = map.get(rootVal)
  leftSize = rootIndex - inStart

  root.left = build(preStart+1, preStart+leftSize, inStart, rootIndex-1)
  root.right = build(preStart+leftSize+1, preEnd, rootIndex+1, inEnd)

  return root
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(n)

## 代码

```java
class Solution {
    private Map<Integer, Integer> inorderMap;

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        inorderMap = new HashMap<>();
        for (int i = 0; i < inorder.length; i++) {
            inorderMap.put(inorder[i], i);
        }
        return build(preorder, 0, preorder.length - 1,
                     inorder, 0, inorder.length - 1);
    }

    private TreeNode build(int[] preorder, int preStart, int preEnd,
                           int[] inorder, int inStart, int inEnd) {
        if (preStart > preEnd) return null;

        int rootVal = preorder[preStart];
        TreeNode root = new TreeNode(rootVal);
        int rootIndex = inorderMap.get(rootVal);
        int leftSize = rootIndex - inStart;

        root.left = build(preorder, preStart + 1, preStart + leftSize,
                          inorder, inStart, rootIndex - 1);
        root.right = build(preorder, preStart + leftSize + 1, preEnd,
                           inorder, rootIndex + 1, inEnd);

        return root;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 中序根索引的哈希预处理 | 每步 O(1) 查找，避免 O(n) 扫描 |
| leftSize 计算 | `rootIndex - inStart` 为左子树节点数 |
| 边界更新 | 前序范围根据 leftSize 划分，后续问题用根索引划分 |
