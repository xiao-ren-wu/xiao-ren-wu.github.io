# 两两交换链表中的节点

## 题目描述

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部值的情况下完成本题（即只能进行节点交换）。

**示例 1：**

```
输入：head = [1,2,3,4]
输出：[2,1,4,3]
```

**示例 2：**

```
输入：head = []
输出：[]
```

**示例 3：**

```
输入：head = [1]
输出：[1]
```

**提示：**

- 链表中节点的数目在范围 `[0, 100]` 内
- `0 <= Node.val <= 100`

## 解题思路

### 核心思想：迭代（三指针）

```
dummy → 1 → 2 → 3 → 4
         ↑   ↑
        first second

交换后：
dummy → 2 → 1 → 3 → 4

三指针跟踪：
prev → 交换后的一对的前一个节点
first → 第一个节点
second → 第二个节点

交换步骤：
  prev.next = second
  first.next = second.next
  second.next = first
  移动 prev 到 first
```

### 复杂度分析

- **时间复杂度**：O(n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public ListNode swapPairs(ListNode head) {
        ListNode dummy = new ListNode(0, head);
        ListNode prev = dummy;

        while (prev.next != null && prev.next.next != null) {
            ListNode first = prev.next;
            ListNode second = prev.next.next;

            prev.next = second;
            first.next = second.next;
            second.next = first;

            prev = first;
        }

        return dummy.next;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 节点数为奇数时最后一个不交换 | while 条件检查 next 和 next.next 都不为空 |
| 不能只交换值 | 题目明确要求只进行节点交换 |
| dummy 节点的作用 | 统一处理头节点的交换 |
