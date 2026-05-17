# 合并区间

## 题目描述

以数组 `intervals` 表示若干个区间的集合，其中单个区间为 `intervals[i] = [starti, endi]`。请你合并所有重叠的区间，并返回一个不重叠的区间数组。

**示例 1：**

```
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠，合并为 [1,6]。
```

**示例 2：**

```
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可视为重叠。
```

## 解题思路

### 核心思想：排序 + 遍历

按左端点排序，遍历时判断当前 intervals[i] 是否与 merged 最后一个区间重叠。

```
按 intervals[i][0] 升序排序
merged = []
for interval in intervals:
  if merged 为空 || interval[0] > merged.last[1]:
    无重叠，直接加入
  else:
    有重叠，更新 merged.last[1] = max(merged.last[1], interval[1])
```

### 复杂度分析

- **时间复杂度**：O(n log n)，排序为主要开销
- **空间复杂度**：O(log n) 或 O(n)，取决于排序算法

## 代码

```java
class Solution {
    public int[][] merge(int[][] intervals) {
        if (intervals.length == 0) return new int[0][];

        Arrays.sort(intervals, (a, b) -> a[0] - b[0]);

        List<int[]> merged = new ArrayList<>();
        merged.add(intervals[0]);

        for (int i = 1; i < intervals.length; i++) {
            int[] last = merged.get(merged.size() - 1);
            if (intervals[i][0] > last[1]) {
                merged.add(intervals[i]);
            } else {
                last[1] = Math.max(last[1], intervals[i][1]);
            }
        }

        return merged.toArray(new int[merged.size()][]);
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 按左端点排序 | 不排序无法直接判断重叠 |
| 无重叠加新数组 | `interval[0] > last[1]` 说明不重叠 |
| 重叠只更新右端点 | 左端点不变，右端点取较大值 |
