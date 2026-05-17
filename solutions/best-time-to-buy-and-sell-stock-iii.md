# 买卖股票的最佳时机 III

## 题目描述

给定一个数组，它的第 `i` 个元素是一支给定的股票在第 `i` 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 **两笔** 交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

**示例 1：**

```
输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股价=0）买入，第 6 天（股价=3）卖出，利润=3
     在第 7 天（股价=1）买入，第 8 天（股价=4）卖出，利润=3
     总利润 = 3+3 = 6
```

**示例 2：**

```
输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股价=1）买入，第 5 天（股价=5）卖出，利润=4
     只能做一笔交易（因为只有一次买卖才能获得最大利润）
```

**示例 3：**

```
输入：prices = [7,6,4,3,1]
输出：0
解释：无法完成任何交易，最大利润为 0
```

**提示：**

- `1 <= prices.length <= 10^5`
- `0 <= prices[i] <= 10^5`

## 解题思路

### 核心思想：状态机 DP

最多两笔交易，可以将交易过程拆解为 4 个有序状态：

```
状态 0: 第一次买入  (buy1)
状态 1: 第一次卖出  (sell1)
状态 2: 第二次买入  (buy2)
状态 3: 第二次卖出  (sell2)
```

每个状态从前一个状态转移而来：

```
buy1  = max(buy1,  -price)      // 第一次买入，花钱
sell1 = max(sell1, buy1 + price) // 第一次卖出，赚钱
buy2  = max(buy2, sell1 - price) // 第二次买入，用第一次赚的钱
sell2 = max(sell2, buy2 + price) // 第二次卖出
```

### 算法步骤

```
初始化：
  buy1 = buy2 = -∞（或 -prices[0]）
  sell1 = sell2 = 0

遍历每个价格 price：
  buy1  = max(buy1,  -price)
  sell1 = max(sell1, buy1 + price)
  buy2  = max(buy2, sell1 - price)
  sell2 = max(sell2, buy2 + price)

返回 sell2
```

### 图解示例

```
prices = [3,3,5,0,0,3,1,4]

Day 1: price=3  buy1=-3  sell1=0   buy2=-3 sell2=0
Day 2: price=3  buy1=-3  sell1=0   buy2=-3 sell2=0
Day 3: price=5  buy1=-3  sell1=2   buy2=-3 sell2=2
Day 4: price=0  buy1=0   sell1=2   buy2=2  sell2=2
Day 5: price=0  buy1=0   sell1=2   buy2=2  sell2=2
Day 6: price=3  buy1=0   sell1=3   buy2=2  sell2=5
Day 7: price=1  buy1=0   sell1=3   buy2=2  sell2=5
Day 8: price=4  buy1=0   sell1=3   buy2=2  sell2=6 ← 答案
```

从 Day 6 可以看出 sell2=5 已经完成了两笔交易（0→3 和 0→3 但第一次卖在3，第二次买在0卖在3）。Day 8 得到最终利润 6。

### 复杂度分析

- **时间复杂度**：O(n)，一次遍历
- **空间复杂度**：O(1)，只有 4 个变量

## 代码

```java
class Solution {
    public int maxProfit(int[] prices) {
        int buy1 = Integer.MIN_VALUE;
        int sell1 = 0;
        int buy2 = Integer.MIN_VALUE;
        int sell2 = 0;

        for (int price : prices) {
            buy1 = Math.max(buy1, -price);
            sell1 = Math.max(sell1, buy1 + price);
            buy2 = Math.max(buy2, sell1 - price);
            sell2 = Math.max(sell2, buy2 + price);
        }

        return sell2;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 状态顺序不能乱 | buy1 → sell1 → buy2 → sell2 必须按这个顺序，因为必须先买才能卖 |
| buy1 初始化为 `-prices[0]` | 负无穷也可以，但 -prices[0] 更自然 |
| `Integer.MIN_VALUE` 做初始值 | 表示"还未买入"的状态，Math.max 时不会影响结果 |
| 可以只有一次交易 | 如果一次交易利润更大，sell2 会继承 sell1 的结果（因为 buy2 也可能继承 buy1） |
| 扩展到 k 笔交易 | 用两个数组 DP，通用解法见"买卖股票的最佳时机 IV" |
