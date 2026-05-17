# 二叉树中的最大路径和

## 题目描述

二叉树中的 **路径** 被定义为一条节点序列，序列中每对相邻节点之间有一条边。同一个节点在一条路径中**至多出现一次**。该路径**至少包含一个**节点，且不一定经过根节点。

**路径和**是路径中各节点值的总和。

给你一个二叉树的根节点 `root`，返回其**最大路径和**。

**示例 1：**

```
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 → 1 → 3，路径和为 2 + 1 + 3 = 6
```

**示例 2：**

```
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 → 20 → 7，路径和为 15 + 20 + 7 = 42
```

**提示：**

- 树中节点数目范围是 `[1, 3 * 10^4]`
- `-1000 <= Node.val <= 1000`

## 解题思路

### 核心思想：后序遍历 + 最大贡献值

对于二叉树中的任意节点，有两种角色：

1. **作为路径的"拐点"**：路径从左子树上来，经过该节点，再向右子树下去
   - 此时路径和 = node.val + leftGain + rightGain
2. **作为路径的一部分"向上贡献"**：选择左或右子树中更大的那条，加上自身值返回给父节点
   - 返回值 = node.val + max(leftGain, rightGain, 0)

### 算法步骤

```
后序遍历每个节点：

1. 计算左子树的最大贡献值 leftGain
   - 如果左子树贡献为负，取 0（不选左子树）

2. 计算右子树的最大贡献值 rightGain
   - 如果右子树贡献为负，取 0（不选右子树）

3. 当前节点作为拐点的路径和 = node.val + leftGain + rightGain
   - 用全局变量 maxSum 记录最大值

4. 返回当前节点向上贡献的值 = node.val + max(leftGain, rightGain)
```

### 图解示例

```
     -10
     /  \
    9    20
        /  \
       15   7

后序遍历：

节点 9:  leaf, leftGain=0, rightGain=0
  pathSum = 9+0+0=9,  maxSum=9
  向上贡献 = 9+0=9

节点 15: leaf, leftGain=0, rightGain=0
  pathSum = 15, maxSum=15
  向上贡献 = 15

节点 7:  leaf, leftGain=0, rightGain=0
  pathSum = 7, maxSum=15
  向上贡献 = 7

节点 20: leftGain=15, rightGain=7
  pathSum = 20+15+7=42, maxSum=42 ← 最大
  向上贡献 = 20+max(15,7)=35

节点 -10: leftGain=9, rightGain=35
  pathSum = -10+9+35=34, maxSum=42

结果：42
```

### 复杂度分析

- **时间复杂度**：O(n)，每个节点访问一次
- **空间复杂度**：O(h)，递归栈深度，h 为树的高度

## 代码

```java
class Solution {
    private int maxSum = Integer.MIN_VALUE;

    public int maxPathSum(TreeNode root) {
        maxGain(root);
        return maxSum;
    }

    private int maxGain(TreeNode node) {
        if (node == null) return 0;

        int leftGain = Math.max(maxGain(node.left), 0);
        int rightGain = Math.max(maxGain(node.right), 0);

        int pathSum = node.val + leftGain + rightGain;
        maxSum = Math.max(maxSum, pathSum);

        return node.val + Math.max(leftGain, rightGain);
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 负增益截断为 0 | `max(gain, 0)`，负数子树不选，舍弃它对路径的负贡献 |
| 全局变量 maxSum | 初始化为 `Integer.MIN_VALUE`，因为节点值可能全为负 |
| 向上贡献只能选一条路 | 节点作为路径的一部分时只能向左或向右，不能分叉 |
| 空节点返回 0 | 递归基：null 节点贡献为 0 |
