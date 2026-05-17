# 排序链表

## 题目描述

给你链表的头结点 `head`，请将其按升序排列并返回排序后的链表。要求 O(n log n) 时间复杂度和 O(1) 额外空间复杂度。

**示例 1：**

```
输入：head = [4,2,1,3]
输出：[1,2,3,4]
```

**示例 2：**

```
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
```

## 解题思路

### 核心思想：归并排序（自底向上）

不能用数组的快排（需要 O(1) 空间且是链表），自底向上归并排序最适合。

```
step = 1
while step < n:
  dummy.next = merge passes（每 step 个节点一组归并）
  step *= 2

// 每次归并：取两个长度为 step 的子链表，合并排序
function merge(l1, l2): 标准双指针归并
function split(head, step): 切出前 step 个节点，返回后半部分
```

### 复杂度分析

- **时间复杂度**：O(n log n)
- **空间复杂度**：O(1)

## 代码

```java
class Solution {
    public ListNode sortList(ListNode head) {
        if (head == null || head.next == null) return head;

        int n = 0;
        ListNode cur = head;
        while (cur != null) { n++; cur = cur.next; }

        ListNode dummy = new ListNode(0);
        dummy.next = head;

        for (int step = 1; step < n; step <<= 1) {
            ListNode prev = dummy, curr = dummy.next;
            while (curr != null) {
                ListNode left = curr;
                ListNode right = split(left, step);
                curr = split(right, step);
                prev = merge(left, right, prev);
            }
        }

        return dummy.next;
    }

    private ListNode split(ListNode head, int step) {
        for (int i = 1; head != null && i < step; i++) head = head.next;
        if (head == null) return null;
        ListNode next = head.next;
        head.next = null;
        return next;
    }

    private ListNode merge(ListNode l1, ListNode l2, ListNode prev) {
        ListNode cur = prev;
        while (l1 != null && l2 != null) {
            if (l1.val <= l2.val) {
                cur.next = l1;
                l1 = l1.next;
            } else {
                cur.next = l2;
                l2 = l2.next;
            }
            cur = cur.next;
        }
        cur.next = l1 != null ? l1 : l2;
        while (cur.next != null) cur = cur.next;
        return cur;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 自底向上而非递归 | 递归归并需要 O(log n) 栈空间，不满足 O(1) |
| split 操作 | 切到第 step 个节点时断链，返回后半部分头 |
| merge 返回尾节点 | 方便 prev 衔接下一次归并 |
