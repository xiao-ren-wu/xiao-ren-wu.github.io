# LRU 缓存

## 题目描述

请你设计并实现一个满足 LRU（最近最少使用）缓存约束的数据结构。

```
class LRUCache(int capacity)
int get(int key)      // 如果 key 存在返回 value，否则 -1
void put(int key, int value)  // 插入或更新。若缓存满，移除最久未使用的 key
```

函数 `get` 和 `put` 必须以 O(1) 的平均时间复杂度运行。

**示例：**

```
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1);           // 缓存是 {1=1}
lRUCache.put(2, 2);           // 缓存是 {1=1, 2=2}
lRUCache.get(1);              // 返回 1
lRUCache.put(3, 3);           // 删除 key 2，缓存是 {1=1, 3=3}
lRUCache.get(2);              // 返回 -1
lRUCache.put(4, 4);           // 删除 key 1，缓存是 {4=4, 3=3}
lRUCache.get(1);              // 返回 -1
lRUCache.get(3);              // 返回 3
lRUCache.get(4);              // 返回 4
```

## 解题思路

### 核心思想：哈希表 + 双向链表

哈希表实现 O(1) 查找，双向链表维护访问顺序：最近使用的移到头部，缓存满时移除尾部。

```
class Node:
  key, value, prev, next

HashMap<Integer, Node> map
head, tail（哨兵节点）

get(key):
  if map 无 key: return -1
  moveToHead(map.get(key))
  return node.value

put(key, value):
  if map 有 key:
    更新值，moveToHead
  else:
    创建新节点，加到头部，map 添加
    if size > capacity: 移除尾部节点
```

### 复杂度分析

- **时间复杂度**：O(1)
- **空间复杂度**：O(capacity)

## 代码

```java
class LRUCache {
    class Node {
        int key, value;
        Node prev, next;
        Node(int k, int v) { key = k; value = v; }
    }

    private int capacity;
    private Map<Integer, Node> map;
    private Node head, tail;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        map = new HashMap<>();
        head = new Node(-1, -1);
        tail = new Node(-1, -1);
        head.next = tail;
        tail.prev = head;
    }

    public int get(int key) {
        if (!map.containsKey(key)) return -1;
        Node node = map.get(key);
        moveToHead(node);
        return node.value;
    }

    public void put(int key, int value) {
        if (map.containsKey(key)) {
            Node node = map.get(key);
            node.value = value;
            moveToHead(node);
        } else {
            Node node = new Node(key, value);
            map.put(key, node);
            addToHead(node);
            if (map.size() > capacity) {
                Node removed = removeTail();
                map.remove(removed.key);
            }
        }
    }

    private void moveToHead(Node node) {
        removeNode(node);
        addToHead(node);
    }

    private void addToHead(Node node) {
        node.next = head.next;
        node.prev = head;
        head.next.prev = node;
        head.next = node;
    }

    private void removeNode(Node node) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
    }

    private Node removeTail() {
        Node node = tail.prev;
        removeNode(node);
        return node;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 哨兵节点 | head 和 tail 作为虚拟节点，简化边界处理 |
| 节点存 key | 移除尾部时需通过 key 删除 map 中的键 |
| 每次访问移动 | get 和 put 更新都已存在的 key 都要 moveToHead |
