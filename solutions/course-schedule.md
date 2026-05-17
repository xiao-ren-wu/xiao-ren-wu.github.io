# 课程表

## 题目描述

你这个学期必须选修 `numCourses` 门课程，记为 `0` 到 `numCourses - 1`。在选修某些课程之前需要一些先修课程：`prerequisites[i] = [ai, bi]` 表示想学 `ai` 必须先学 `bi`。请你判断是否可能完成所有课程的学习。

**示例 1：**

```
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
```

**示例 2：**

```
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
```

## 解题思路

### 核心思想：拓扑排序

构建邻接图和入度数组。用队列做 BFS 拓扑排序，能排完所有课程则无环。

```
// 建图
for [a, b] in prerequisites:
  graph[b].add(a)
  inDegree[a]++

// BFS
queue = [所有 inDegree == 0 的节点]
count = 0
while queue 不为空:
  cur = queue.poll()
  count++
  for next in graph[cur]:
    inDegree[next]--
    if inDegree[next] == 0: queue.offer(next)

return count == numCourses
```

### 复杂度分析

- **时间复杂度**：O(V + E)，V=课程数，E=先修关系数
- **空间复杂度**：O(V + E)

## 代码

```java
class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        List<List<Integer>> graph = new ArrayList<>();
        int[] inDegree = new int[numCourses];

        for (int i = 0; i < numCourses; i++) graph.add(new ArrayList<>());

        for (int[] p : prerequisites) {
            graph.get(p[1]).add(p[0]);
            inDegree[p[0]]++;
        }

        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) queue.offer(i);
        }

        int count = 0;
        while (!queue.isEmpty()) {
            int cur = queue.poll();
            count++;
            for (int next : graph.get(cur)) {
                inDegree[next]--;
                if (inDegree[next] == 0) queue.offer(next);
            }
        }

        return count == numCourses;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 入度为 0 的节点 | 表示没有前置课程，可作为起点 |
| 环检测 | count != numCourses 说明存在环，无法完成 |
| DFS 也可 | 用 visited 状态（0=未访问/1=访问中/2=已访问）检测环 |
