# 字典序的第 K 小数字

## 题目描述

给定整数 `n` 和 `k`，返回范围 `[1, n]` 中所有整数按字典序（即按数字的字符串形式排序）排列中的第 `k` 小的数字。

**示例 1：**

```
输入：n = 13, k = 2
输出：10
解释：字典序排列为 [1,10,11,12,13,2,3,4,5,6,7,8,9]，第 2 小的是 10。
```

**示例 2：**

```
输入：n = 1, k = 1
输出：1
```

**提示：**

- `1 <= k <= n <= 10^9`

## 解题思路

### 核心思想：字典树（10-ary tree）

将 1 到 n 的所有数字视为一棵 10 叉树的前序遍历：

```
                      1                    2 ...  9
                    / | \                / | \
                  10 11 12 ... 19      20 21 22 ...
                 /  \
               100 101 ...
```

字典序排列 = 这棵树的前序遍历顺序。

问题转化为：在 10 叉树上找第 k 个前序遍历的节点。

### 算法步骤

```
curr = 1
k = k - 1（跳过根节点 1）

while k > 0:
  count = 以 curr 为根的子树中的节点数（值 ≤ n）
  
  if count <= k:
    // 整棵子树都可以跳过，移到下一个兄弟节点
    k -= count
    curr++
  else:
    // 第 k 个节点在子树中，进入子树
    curr *= 10
    k--  // 跳过当前节点
```

### 计算子树节点数

```
function countNodes(prefix, n):
  count = 0
  first = prefix
  last = prefix
  while first <= n:
    count += min(last, n) - first + 1
    first *= 10
    last = last * 10 + 9
  return count
```

### 图解示例

```
n = 13, k = 2

curr=1, k=1（已减 1）

以 1 为根的子树节点数：
  第一层：1
  第二层：10-13（4个节点）
  共 5 个节点

count=5, k=1 → count > k，进入子树
  curr=10, k=0

k=0，停止，返回 curr=10 ✓
```

```
n = 100, k = 10

curr=1, k=9

子树计数：1,10-19,100 → 共 13 个
count=13, k=9 → 进入子树
  curr=10, k=8

子树(10): 10,100-109(但n=100只到100) → 共 2 个
count=2, k=8 → 跳过，curr=11, k=6

子树(11): 11,110-119(超n) → 共 1 个
count=1, k=6 → 跳过，curr=12, k=5

...继续直到找到第 10 个
```

### 复杂度分析

- **时间复杂度**：O(log²n)，每次计数 O(log n)，最多移动 O(log n) 次
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public int findKthNumber(int n, int k) {
        int curr = 1;
        k--;

        while (k > 0) {
            long count = countNodes(n, curr);
            if (count <= k) {
                k -= count;
                curr++;
            } else {
                curr *= 10;
                k--;
            }
        }

        return curr;
    }

    private long countNodes(int n, long prefix) {
        long count = 0;
        long first = prefix;
        long last = prefix;

        while (first <= n) {
            count += Math.min(last, n) - first + 1;
            first *= 10;
            last = last * 10 + 9;
        }

        return count;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| k 从 0 开始 | 初始 k-- 是因为我们跳过了根节点（curr=1 本身） |
| 用 long 防止溢出 | n 最大 10^9，first*10 可能超出 int 范围 |
| count ≤ k 时跳过整棵子树 | 第 k 个节点不在当前子树中，移到下一个兄弟 |
| count > k 时进入子树 | 第 k 个节点就在当前子树内，curr *= 10 进入下一层 |
| 不要生成完整的字典序数组 | n 可达 10^9，无法放入内存 |
