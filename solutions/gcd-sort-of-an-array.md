# 数组的最大公因数排序

## 题目描述

给你一个整数数组 `nums`，你可以在 `nums` 上执行下述操作：

如果两个元素 `nums[i]` 和 `nums[j]` 的最大公因数（GCD）大于 1，你可以交换 `nums[i]` 和 `nums[j]` 的位置。

返回 `true`，如果可以通过任意次操作将 `nums` 按非递减顺序排序；否则返回 `false`。

**示例 1：**

```
输入：nums = [7,21,3]
输出：true
解释：可以按如下顺序排序：
- 7 和 21 的最大公因数是 7 > 1，可以交换 7 和 21 的位置。
  得到 [21,7,3]。
- 21 和 3 的最大公因数是 3 > 1，可以交换 21 和 3 的位置。
  得到 [3,7,21]。
```

**示例 2：**

```
输入：nums = [5,2,6,2]
输出：false
解释：无法交换任何元素，因为任意两个数的最大公因数都 ≤ 1。
```

**示例 3：**

```
输入：nums = [10,5,9,3,15]
输出：true
解释：
我们可以交换 10 和 15（gcd=5 > 1），然后交换 10 和 9（gcd=3 > 1），再交换 10 和 3（gcd=3 > 1）。
经过交换可以得到 [3,5,9,10,15]。
```

**提示：**

- `1 <= nums.length <= 3 * 10^4`
- `1 <= nums[i] <= 10^5`

## 解题思路

### 核心思想：并查集 + 素数连通

两个数如果 gcd > 1，说明它们共享至少一个公共质因子。通过这个公共质因子，它们可以间接连通。

核心思路：

1. 每个数与其所有质因子建立连接（Union）
2. 排序数组，检查原数组中每个位置的值与排序后该位置的值是否在同一个连通分量中
3. 如果是，则可以通过交换到达目标位置

### 算法步骤

```
1. 初始化并查集，大小为 max(nums) + 1

2. 对每个 num 分解质因数：
   for each 质因子 factor of num:
     union(num, factor)

3. 复制 nums 并排序得到 sorted

4. 遍历 i：
   if find(nums[i]) != find(sorted[i]):
     return false

5. return true
```

### 图解示例

```
nums = [7, 21, 3]

并查集连通关系：
  7 → primeFactor=7 → union(7,7)
  21 → primeFactor=3 → union(21,3)
     → primeFactor=7 → union(21,7)
  3 → primeFactor=3 → union(3,3)

连通分量：
  {7, 21} 通过质因子 7 连通
  {21, 3} 通过质因子 3 连通
  → {3, 7, 21} 全部连通 ✓

排序后：[3, 7, 21]
原索引0(7) 与 排序后索引0(3) → find(7)==find(3) ✓
原索引1(21) 与 排序后索引1(7) → find(21)==find(7) ✓
原索引2(3) 与 排序后索引2(21) → find(3)==find(21) ✓
```

### 复杂度分析

- **时间复杂度**：O(n × √m + n log n)，m = max(nums)，筛法预处理质因子
- **空间复杂度**：O(m + n)

## 代码

```java
class Solution {
    public boolean gcdSort(int[] nums) {
        int max = 0;
        for (int num : nums) max = Math.max(max, num);

        int[] parent = new int[max + 1];
        for (int i = 0; i <= max; i++) parent[i] = i;

        // 筛法找到每个数的最小质因子
        int[] spf = new int[max + 1];
        for (int i = 2; i <= max; i++) {
            if (spf[i] == 0) {
                for (int j = i; j <= max; j += i) {
                    if (spf[j] == 0) spf[j] = i;
                }
            }
        }

        // 对每个 num 连接其所有质因子
        for (int num : nums) {
            int x = num;
            while (x > 1) {
                int p = spf[x];
                union(parent, num, p);
                while (x % p == 0) x /= p;
            }
        }

        // 检查
        int[] sorted = nums.clone();
        Arrays.sort(sorted);
        for (int i = 0; i < nums.length; i++) {
            if (find(parent, nums[i]) != find(parent, sorted[i])) {
                return false;
            }
        }
        return true;
    }

    private int find(int[] parent, int x) {
        if (parent[x] != x) parent[x] = find(parent, parent[x]);
        return parent[x];
    }

    private void union(int[] parent, int a, int b) {
        int ra = find(parent, a);
        int rb = find(parent, b);
        if (ra != rb) parent[ra] = rb;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 质因子分解用 SPF（最小质因子） | 筛法预处理 SPF 可以快速分解所有数的质因子 |
| union 连通的是 num 与它的质因子 | 不是 num 之间直接连通，是通过质因子间接连通 |
| 并查集大小是 max(nums)+1 | 因为质因子的值也在 [2, max] 范围内 |
| 检查的是排序前后相同位置 | 如果每个位置的值都能通过交换到达目标位置，整体就可排序 |
| 1 没有质因子 | 1 不参与连通，它的位置无法与任何其他元素交换 |
