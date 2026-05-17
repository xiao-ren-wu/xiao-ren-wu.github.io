# 设计跳表

## 题目描述

不使用任何库函数，设计一个 **跳表**（Skiplist）。

跳表是一种能够在 O(log n) 时间复杂度内完成插入、删除和搜索操作的数据结构。

跳表中有很多层，每一层是一个有序的链表。最底层的链表包含所有元素。每个节点除了指向同一层的下一个节点之外，还可能包含指向下一层或上一层的指针。

实现 `Skiplist` 类：

- `Skiplist()` — 初始化跳表对象
- `bool search(int target)` — 返回 target 是否存在于跳表中
- `void add(int num)` — 向跳表中插入一个元素
- `bool erase(int num)` — 在跳表中删除一个元素，如果存在则删除并返回 true，否则返回 false

**示例：**

```
Skiplist skiplist = new Skiplist();
skiplist.add(1);
skiplist.add(2);
skiplist.add(3);
skiplist.search(0);   // false
skiplist.add(4);
skiplist.search(1);   // true
skiplist.erase(0);    // false（0 不在跳表中）
skiplist.erase(1);    // true
skiplist.search(1);   // false（1 已被擦除）
```

**提示：**

- `0 <= num, target <= 2 * 10^4`
- 最多调用 `5 * 10^4` 次 `search`、`add` 和 `erase` 操作

## 解题思路

### 核心思想：多层链表 + 概率平衡

跳表是一种通过 **随机化** 实现近似平衡的数据结构。它用多层有序链表来加速查找：

```
Level 3: 1 ──────────────→ 9
Level 2: 1 ─────→ 5 ────→ 9
Level 1: 1 ─→ 3 → 5 ─→ 7 → 9
Level 0: 1 → 2 → 3 → 5 → 7 → 9
```

查找时从最高层开始，每层向右找到最后一个 ≤ target 的节点，然后向下到下一层继续。

### 数据结构设计

```java
class Node {
    int val;
    Node[] next; // next[i] 指向同一层的下一个节点
    Node(int val, int level) {
        this.val = val;
        this.next = new Node[level + 1];
    }
}
```

### 算法步骤

```
search(target):
  从 highestLevel 开始向下遍历：
    while next.val < target: 向右移动
    然后向下移动一层
  最终在 level 0 检查 next.val == target

add(num):
  1. 找到每层的前驱节点（类似 search，但记录所有前驱）
  2. 随机决定新节点的层数 randomLevel()
  3. 将新节点插入每一层的前驱之后

erase(num):
  1. 找到每层的前驱节点
  2. 如果 level 0 的 next.val != num，返回 false
  3. 对每一层，将前驱的 next 指向被删除节点的 next
```

### 节点层数随机化

```java
int randomLevel() {
    int level = 0;
    while (Math.random() < 0.5 && level < MAX_LEVEL) {
        level++;
    }
    return level;
}
```

### 复杂度分析

- **时间复杂度**：O(log n) 期望，最坏 O(n)
- **空间复杂度**：O(n log n) 期望

## 代码

```java
class Skiplist {
    private static final int MAX_LEVEL = 16;
    private final Node head = new Node(-1, MAX_LEVEL);
    private int currentLevel = 0;

    static class Node {
        int val;
        Node[] next;

        Node(int val, int level) {
            this.val = val;
            this.next = new Node[level + 1];
        }
    }

    public boolean search(int target) {
        Node cur = head;
        for (int i = currentLevel; i >= 0; i--) {
            while (cur.next[i] != null && cur.next[i].val < target) {
                cur = cur.next[i];
            }
        }
        cur = cur.next[0];
        return cur != null && cur.val == target;
    }

    public void add(int num) {
        Node[] prev = new Node[MAX_LEVEL + 1];
        Node cur = head;

        for (int i = currentLevel; i >= 0; i--) {
            while (cur.next[i] != null && cur.next[i].val < num) {
                cur = cur.next[i];
            }
            prev[i] = cur;
        }

        int level = randomLevel();
        if (level > currentLevel) {
            for (int i = currentLevel + 1; i <= level; i++) {
                prev[i] = head;
            }
            currentLevel = level;
        }

        Node newNode = new Node(num, level);
        for (int i = 0; i <= level; i++) {
            newNode.next[i] = prev[i].next[i];
            prev[i].next[i] = newNode;
        }
    }

    public boolean erase(int num) {
        Node[] prev = new Node[MAX_LEVEL + 1];
        Node cur = head;

        for (int i = currentLevel; i >= 0; i--) {
            while (cur.next[i] != null && cur.next[i].val < num) {
                cur = cur.next[i];
            }
            prev[i] = cur;
        }

        cur = cur.next[0];
        if (cur == null || cur.val != num) return false;

        for (int i = 0; i <= currentLevel; i++) {
            if (prev[i].next[i] == cur) {
                prev[i].next[i] = cur.next[i];
            }
        }

        while (currentLevel > 0 && head.next[currentLevel] == null) {
            currentLevel--;
        }

        return true;
    }

    private int randomLevel() {
        int level = 0;
        while (Math.random() < 0.5 && level < MAX_LEVEL) {
            level++;
        }
        return level;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 每层前驱都要记录 | add 和 erase 需要知道每一层前驱，才能插入或删除节点 |
| 新节点层数可能超过当前最高层 | 超过的部分前驱设为 head，并更新 currentLevel |
| 删除后更新最高层 | 如果最高层空了，currentLevel 要降低 |
| randomLevel 概率 0.5 | 标准实现，平均每层节点数是下层的 1/2 |
| MAX_LEVEL 选择 | 16 对于最多 50000 次操作足够了，log₂(50000) ≈ 16 |
