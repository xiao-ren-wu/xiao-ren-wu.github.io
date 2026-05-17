# 数据流的中位数

## 题目描述

中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

设计一个支持以下两种操作的数据结构：

- `void addNum(int num)` — 从数据流中添加一个整数到数据结构中
- `double findMedian()` — 返回目前所有元素的中位数

**示例：**

```
addNum(1)
addNum(2)
findMedian() → 1.5
addNum(3)
findMedian() → 2.0
```

**进阶：**

1. 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化？
2. 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化？

## 解题思路

### 核心思想：两个堆

将数据分成两半：

- **左半部分**：用**大顶堆**（max-heap）存储，堆顶是左半部分的最大值
- **右半部分**：用**小顶堆**（min-heap）存储，堆顶是右半部分的最小值

保持 `left.size() == right.size()` 或 `left.size() == right.size() + 1`

这样中位数总是：
- 奇数个元素 → `left.peek()`（左半多一个）
- 偶数个元素 → `(left.peek() + right.peek()) / 2.0`

### 算法步骤

```
addNum(num):
  1. 先将 num 加入左堆（大顶堆）
  2. 将左堆的最大值移到右堆
     （确保左堆的所有元素 ≤ 右堆的所有元素）
  3. 如果右堆比左堆多，将右堆的最小值移回左堆
     （保持大小平衡）

findMedian():
  if left.size() > right.size():
    return left.peek()
  else:
    return (left.peek() + right.peek()) / 2.0
```

### 图解示例

```
addNum(1):   left=[1]  right=[]     → 中位数=1
addNum(2):   left=[1]  right=[2]    → 中位数=(1+2)/2=1.5
addNum(3):   left=[2,1] right=[3]   → 中位数=2
addNum(4):   left=[2,1] right=[3,4] → 中位数=(2+3)/2=2.5
addNum(5):   left=[3,2,1] right=[4,5] → 中位数=3
```

### 复杂度分析

| 操作 | 复杂度 |
|------|--------|
| addNum | O(log n) — 堆的插入和删除 |
| findMedian | O(1) — 直接取堆顶 |
| 空间 | O(n) |

## 代码

```java
class MedianFinder {
    private PriorityQueue<Integer> left;  // 大顶堆
    private PriorityQueue<Integer> right; // 小顶堆

    public MedianFinder() {
        left = new PriorityQueue<>((a, b) -> b - a);
        right = new PriorityQueue<>();
    }

    public void addNum(int num) {
        left.offer(num);
        right.offer(left.poll());

        if (right.size() > left.size()) {
            left.offer(right.poll());
        }
    }

    public double findMedian() {
        if (left.size() > right.size()) {
            return left.peek();
        }
        return (left.peek() + right.peek()) / 2.0;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 左大右小 | 左堆存较小的一半（大顶堆），右堆存较大的一半（小顶堆） |
| 三步操作确保有序性 | 先加左→平衡最大值到右→再平衡大小，保证左所有值 ≤ 右所有值 |
| `b - a` 的坑 | 当 a 和 b 差值很大时可能溢出，但本题在 int 范围内安全 |
| 进阶优化 | 如果范围 0-100，用计数数组 + 维护两个计数指针可做到 O(1) add |
| 堆的初始容量 | 可以指定初始容量减少扩容开销，但非必须 |
