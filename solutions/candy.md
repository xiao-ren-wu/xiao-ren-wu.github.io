# 分发糖果

## 题目描述

`n` 个孩子站成一排，给你一个整数数组 `ratings` 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

- 每个孩子至少分配到 1 个糖果
- 相邻两个孩子中，评分更高的孩子会获得更多的糖果

请你计算最少需要准备多少颗糖果。

**示例 1：**

```
输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
```

**示例 2：**

```
输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只要 1 颗糖果也可以满足条件（评分相等时不要求更多）。
```

**提示：**

- `n == ratings.length`
- `1 <= n <= 2 * 10^4`
- `0 <= ratings[i] <= 2 * 10^4`

## 解题思路

### 核心思想：两遍扫描

"每个孩子至少一颗" + "相邻中评分高的得更多" → 这相当于两个约束条件：

1. **从左到右**：如果 `ratings[i] > ratings[i-1]`，则 `candies[i] > candies[i-1]`
2. **从右到左**：如果 `ratings[i] > ratings[i+1]`，则 `candies[i] > candies[i+1]`

单独满足一个约束很简单，同时满足就需要两趟。

### 算法步骤

```
1. 初始化 candies 数组，全部设为 1

2. 左到右扫描：
   for i = 1 to n-1:
     if ratings[i] > ratings[i-1]:
       candies[i] = candies[i-1] + 1

3. 右到左扫描：
   for i = n-2 downto 0:
     if ratings[i] > ratings[i+1]:
       candies[i] = max(candies[i], candies[i+1] + 1)

4. 求和
```

### 图解示例

```
ratings = [1, 0, 2]

初始化：candies = [1, 1, 1]

左到右：
  i=1: 0 < 1 → 不变，candies = [1, 1, 1]
  i=2: 2 > 0 → candies[2] = 1+1=2，candies = [1, 1, 2]

右到左：
  i=1: 0 < 2 → 不变，candies = [1, 1, 2]
  i=0: 1 > 0 → candies[0] = max(1, 1+1)=2，candies = [2, 1, 2]

总和 = 2+1+2 = 5
```

```
ratings = [1, 2, 2]

初始化：candies = [1, 1, 1]

左到右：
  i=1: 2 > 1 → candies[1]=2，candies = [1, 2, 1]
  i=2: 2 = 2 → 不变，candies = [1, 2, 1]

右到左：
  i=1: 2 = 2 → 不变，candies = [1, 2, 1]
  i=0: 1 < 2 → 不变，candies = [1, 2, 1]

总和 = 1+2+1 = 4
```

### 复杂度分析

- **时间复杂度**：O(n)，两趟线性扫描
- **空间复杂度**：O(n)，糖果分配数组

## 代码

```java
class Solution {
    public int candy(int[] ratings) {
        int n = ratings.length;
        int[] candies = new int[n];
        Arrays.fill(candies, 1);

        for (int i = 1; i < n; i++) {
            if (ratings[i] > ratings[i - 1]) {
                candies[i] = candies[i - 1] + 1;
            }
        }

        for (int i = n - 2; i >= 0; i--) {
            if (ratings[i] > ratings[i + 1]) {
                candies[i] = Math.max(candies[i], candies[i + 1] + 1);
            }
        }

        int sum = 0;
        for (int c : candies) sum += c;
        return sum;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 第二趟用 `max` 而非直接赋值 | 第一趟已经确定了部分约束，第二趟不能破坏它，只能取较大值 |
| 相等时不要求更多糖果 | `ratings[i] == ratings[i+1]` 时，candies[i] 可以不调整 |
| 初始化全为 1 | 保证"每个孩子至少 1 个糖果"的约束 |
| O(1) 空间也可以做 | 使用 slope-counting 方法，但实现更复杂，两趟扫描更直观 |
