# 二叉树的锯齿形层序遍历

## 题目描述

给你二叉树的根节点 `root`，返回其节点值的锯齿形层序遍历（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

**示例 1：**

```
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]
```

**示例 2：**

```
输入：root = [1]
输出：[[1]]
```

**示例 3：**

```
输入：root = []
输出：[]
```

## 解题思路

### 核心思想：BFS + 奇偶层翻转

正常 BFS 层序遍历，偶数层（从 0 开始）从左到右，奇数层从右到左。

```
queue = [root]
level = 0

while queue 不为空:
  levelSize = queue.size()
  deque = LinkedList  // 用双端队列
  for i in [0, levelSize-1]:
    node = queue.poll()
    if level % 2 == 0:
      list.add(node.val)       // 偶数层：尾插
    else:
      list.addFirst(node.val)  // 奇数层：头插
    if node.left: queue.offer(node.left)
    if node.right: queue.offer(node.right)
  result.add(list)
  level++
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(n)

## 代码

```java
class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) return res;

        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        boolean leftToRight = true;

        while (!queue.isEmpty()) {
            int size = queue.size();
            LinkedList<Integer> level = new LinkedList<>();

            for (int i = 0; i < size; i++) {
                TreeNode node = queue.poll();

                if (leftToRight) {
                    level.addLast(node.val);
                } else {
                    level.addFirst(node.val);
                }

                if (node.left != null) queue.offer(node.left);
                if (node.right != null) queue.offer(node.right);
            }

            res.add(level);
            leftToRight = !leftToRight;
        }

        return res;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 用 LinkedList 而非 ArrayList | 头插效率高 O(1) |
| levelSize 固定 | 提前取 size，避免队列变化影响 |
| 方向标记 | 可用 boolean leftToRight 或用层级奇偶判断 |
