# 删除链表的倒数第 N 个结点

## 题目描述

给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。

**示例 1：**

```
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
```

**示例 2：**

```
输入：head = [1], n = 1
输出：[]
```

**示例 3：**

```
输入：head = [1,2], n = 1
输出：[1]
```

**提示：**

- 链表中结点的数目为 `sz`
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

## 解题思路

### 核心思想：快慢指针

快指针先走 n 步，然后快慢指针同时走。当快指针到达末尾时，慢指针正好在倒数第 n 个节点的前一个位置。

```
dummy = new ListNode(0, head)
fast = slow = dummy

for i in [0..n]:
  fast = fast.next

while fast != null:
  fast = fast.next
  slow = slow.next

slow.next = slow.next.next
return dummy.next
```

### 复杂度分析

- **时间复杂度**：O(L)，链表长度
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dummy = new ListNode(0, head);
        ListNode fast = dummy, slow = dummy;

        for (int i = 0; i <= n; i++) fast = fast.next;

        while (fast != null) {
            fast = fast.next;
            slow = slow.next;
        }

        slow.next = slow.next.next;
        return dummy.next;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 用 dummy 节点 | 避免删除头节点时需要特殊处理 |
| fast 先走 n+1 步 | 因为 dummy 在 head 之前，要走到目标节点的前一个 |
| n 保证有效 | 题目保证 1 ≤ n ≤ sz |
