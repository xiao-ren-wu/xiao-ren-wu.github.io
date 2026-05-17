# 两数相加

## 题目描述

给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，每个节点只能存储一位数字。请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

**示例 1：**

```
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807
```

**示例 2：**

```
输入：l1 = [0], l2 = [0]
输出：[0]
```

**示例 3：**

```
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
```

**提示：**

- 每个链表中的节点数在范围 `[1, 100]` 内
- `0 <= Node.val <= 9`
- 题目保证列表表示的数字不含前导零

## 解题思路

### 核心思想：逐位相加 + 进位

同时遍历两个链表，逐位相加并处理进位。注意处理长度不一致和最后进位的情况。

```
dummy → 结果链表头
carry = 0

while l1 != null || l2 != null || carry != 0:
  val1 = l1.val if l1 else 0
  val2 = l2.val if l2 else 0
  sum = val1 + val2 + carry
  carry = sum / 10
  digit = sum % 10
  创建新节点 digit
  l1 = l1.next if l1
  l2 = l2.next if l2
```

### 复杂度分析

- **时间复杂度**：O(max(m, n))，遍历较长的链表
- **空间复杂度**：O(1)（不计输出空间）

## 代码

```java
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0);
        ListNode cur = dummy;
        int carry = 0;

        while (l1 != null || l2 != null || carry != 0) {
            int sum = carry;
            if (l1 != null) {
                sum += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                sum += l2.val;
                l2 = l2.next;
            }
            carry = sum / 10;
            cur.next = new ListNode(sum % 10);
            cur = cur.next;
        }

        return dummy.next;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 最后进位不能忘 | 遍历完两个链表后如果 carry>0 还需再加一个节点 |
| 链表长度不一致 | 短链表遍历完后视为 0 继续计算 |
| dummy 节点技巧 | 用哑节点避免处理头节点的特殊情况 |
