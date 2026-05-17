# 复制带随机指针的链表

## 题目描述

给你一个长度为 n 的链表，每个节点除了 `next` 指针外，还有一个 `random` 指针指向链表中的任意节点或 null。请深拷贝这个链表。

**示例 1：**

```
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
```

**示例 2：**

```
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
```

## 解题思路

### 核心思想：原地复制 + 拆分

三步走：在每个原节点后面创建新节点 → 设置新节点的 random 指针 → 拆分为原链表和新链表。

```
// 第一步：复制节点
cur = head
while cur != null:
  newNode = new Node(cur.val)
  newNode.next = cur.next
  cur.next = newNode
  cur = newNode.next

// 第二步：设置 random
cur = head
while cur != null:
  cur.next.random = cur.random == null ? null : cur.random.next
  cur = cur.next.next

// 第三步：拆分
old = head, newHead = head.next
while old != null:
  newNode = old.next
  old.next = newNode.next
  newNode.next = newNode.next == null ? null : newNode.next.next
  old = old.next
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public Node copyRandomList(Node head) {
        if (head == null) return null;

        // 1. 复制节点
        Node cur = head;
        while (cur != null) {
            Node newNode = new Node(cur.val);
            newNode.next = cur.next;
            cur.next = newNode;
            cur = newNode.next;
        }

        // 2. 设置 random
        cur = head;
        while (cur != null) {
            cur.next.random = cur.random == null ? null : cur.random.next;
            cur = cur.next.next;
        }

        // 3. 拆分
        Node newHead = head.next;
        cur = head;
        while (cur != null) {
            Node newNode = cur.next;
            cur.next = newNode.next;
            newNode.next = newNode.next == null ? null : newNode.next.next;
            cur = cur.next;
        }

        return newHead;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| random 可能为空 | 设置 random 和拆分时都要判空 |
| 拆分的顺序 | 先恢复原链表 next，再设置新链表 next |
| 哈希表方法 | 也可用 Map<原节点,新节点>，O(n) 空间，但容易理解 |
