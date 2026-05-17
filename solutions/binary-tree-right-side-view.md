# 二叉树的右视图

## 题目描述

给定一个二叉树的根节点 `root`，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

**示例 1：**

```
输入：root = [1,2,3,null,5,null,4]
输出：[1,3,4]
```

**示例 2：**

```
输入：root = [1,null,3]
输出：[1,3]
```

**示例 3：**

```
输入：root = []
输出：[]
```

## 解题思路

### 核心思想：BFS 层序遍历

层序遍历，每层取最后一个节点即可。也可用 DFS：先遍历右子树，每层第一个访问的节点即为右视图节点。

```
// BFS 方法
queue = [root]
while queue 不为空:
  size = queue.size()
  for i in [0, size-1]:
    node = queue.poll()
    if i == size-1: res.add(node.val)
    if node.left: queue.offer(node.left)
    if node.right: queue.offer(node.right)
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(n)

## 代码

```java
class Solution {
    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        if (root == null) return res;

        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);

        while (!queue.isEmpty()) {
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                TreeNode node = queue.poll();
                if (i == size - 1) res.add(node.val);
                if (node.left != null) queue.offer(node.left);
                if (node.right != null) queue.offer(node.right);
            }
        }

        return res;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 每层最后一个节点 | BFS 中取 `i == size-1` |
| DFS 可替代 | 先右后左递归，每深度首个节点即为右视图节点 |
| 空树处理 | root == null 返回空列表 |
