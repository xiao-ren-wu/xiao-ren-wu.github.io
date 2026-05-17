# 串联所有单词的子串

## 题目描述

给定一个字符串 `s` 和一个字符串数组 `words`。`words` 中所有字符串 **长度相同**。

`s` 中的 **串联子串** 是指一个包含 `words` 中所有字符串以任意顺序排列连接起来的子串。

例如，如果 `words = ["ab","cd","ef"]`，那么 `"abcdef"`，`"abefcd"`，`"cdabef"`，`"cdefab"`，`"efabcd"` 和 `"efcdab"` 都是串联子串。`"acdbef"` 不是串联子串，因为他不是任何 `words` 排列的连接。

返回所有串联子串在 `s` 中的开始索引。你可以以 **任意顺序** 返回答案。

**示例 1：**

```
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar"。
```

**示例 2：**

```
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[8]
解释：从索引 8 开始的子串是 "goodgoodbestword"，是 ["word","good","best","word"] 的一个排列。
```

**示例 3：**

```
输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]
```

**提示：**

- `1 <= s.length <= 10^4`
- `1 <= words.length <= 5000`
- `1 <= words[i].length <= 30`
- `words[i]` 和 `s` 由小写英文字母组成

---

## 解题思路

### 核心观察

`words` 中所有字符串 **长度相同**，这是本题最关键的条件。它意味着我们可以按单词粒度（`wordLen`）来切割 `s`，从而将问题转化为 **词级别的滑动窗口**。

设 `wordLen = words[0].length()`，`wordCount = words.length`，串联子串的总长度固定为 `wordLen * wordCount`。

### 算法：词级滑动窗口

#### 为什么需要枚举 offset？

串联子串可以从任意位置开始，但所有单词长度相同，所以起始位置对 `wordLen` 取模的结果决定了它落在哪个偏移类中。我们只需枚举 `[0, wordLen)` 这 `wordLen` 个偏移量，在每个偏移类内以 `wordLen` 为步长滑动窗口。

```
s = "barfoothefoobarman"

offset=0: bar | foo | the | foo | bar | man   匹配 [bar,foo]
offset=1: arf | oot | hef | oob | arm         乱序，不可能匹配
offset=2: rfo | oth | efo | oba | rm          乱序，不可能匹配
```

#### 滑动窗口规则

对每个 offset：

1. **遇到不在 `needMap` 中的词** → 整个窗口作废，清空 window，重置 left
2. **遇到合法词** → 加入 window，更新 matched
3. **如果当前词超量**（`window.get(cur) > needMap.get(cur)`）→ 从左侧逐个移出单词，同时更新 window 和 matched，直到当前词数量合法
4. **如果 matched == needMap.size()** → 记录 left，然后左移一个单词，继续滑动

### 为什么 matched 用词种类数而非总词数？

`matched` 统计的是「有多少种单词的数量已经恰好达标」。当 `matched == needMap.size()` 时，说明每种词的数量都正好满足要求，**总词数也必然等于 `wordCount`**。这种方式在收缩窗口时只需比较 `equals` 就能知道是否需要减少 matched，比维护总词数更简洁。

### 复杂度分析

- **时间复杂度**：`O(n)` — 每个字符最多被访问两次（进入窗口和移出窗口），`n = s.length()`
- **空间复杂度**：`O(m)` — `m = words` 中不同单词的数量

---

## 图解示例

```
s = "barfoothefoobarman", words = ["foo","bar"]

needMap = {"foo":1, "bar":1}, wordLen=3, wordCount=2, totalLeft=6
```

**offset = 0：**

```
right=0: "bar" → window={bar:1}    matched=1
right=3: "foo" → window={bar:1, foo:1}  matched=2
  matched(2)==needMap.size(2) → 记录 left=0
  左移: window={bar:0, foo:1}, matched=1

right=6: "the" → 非法词，重置
  matched=0, window={}, left=9

right=9:  "foo" → window={foo:1}    matched=1
right=12: "bar" → window={foo:1, bar:1}  matched=2
  matched(2)==needMap.size(2) → 记录 left=9
  左移: window={foo:0, bar:1}, matched=1

结果: [0, 9]
```

---

## 题解代码

```java
import java.util.*;

class Solution {
    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> resList = new ArrayList<>();
        int wordLen = words[0].length();
        if (wordLen * words.length > s.length()) {
            return resList;
        }
        Map<String, Integer> needMap = new HashMap<>();
        for (String word : words) {
            needMap.put(word, needMap.getOrDefault(word, 0) + 1);
        }
        for (int offset = 0; offset < wordLen; offset++) {
            int left = offset;
            int matched = 0;
            Map<String, Integer> window = new HashMap<>();
            for (int right = offset; right <= s.length() - wordLen; right += wordLen) {
                String cur = s.substring(right, right + wordLen);
                if (!needMap.containsKey(cur)) {
                    matched = 0;
                    window.clear();
                    left = right + wordLen;
                    continue;
                }
                window.put(cur, window.getOrDefault(cur, 0) + 1);
                if (Objects.equals(window.get(cur), needMap.get(cur))) {
                    matched++;
                }
                while (window.get(cur) > needMap.get(cur)) {
                    String leftWord = s.substring(left, left + wordLen);
                    if (Objects.equals(window.get(leftWord), needMap.get(leftWord))) {
                        matched--;
                    }
                    window.put(leftWord, window.get(leftWord) - 1);
                    left += wordLen;
                }
                if (matched == needMap.size()) {
                    resList.add(left);
                    String leftWord = s.substring(left, left + wordLen);
                    if (Objects.equals(window.get(leftWord), needMap.get(leftWord))) {
                        matched--;
                    }
                    window.put(leftWord, window.get(leftWord) - 1);
                    left += wordLen;
                }
            }
        }
        return resList;
    }
}
```

---

## 关键细节总结

| 细节                   | 说明                                                      |
| -------------------- | ------------------------------------------------------- |
| 枚举 `[0, wordLen)` 偏移 | 所有串联子串的起始位置必然落在这 wordLen 个余数类中                          |
| `matched` 统计词种类数     | 无需维护总词数，减少 equals 判断的次数                                 |
| while 收缩取 `leftWord` | 必须读 `s.substring(left, left+wordLen)` 拿到实际词，不能假定是 `cur` |
| 收缩时联动更新 `matched`    | 只有当 `Objects.equals` 成立（从达标变成不达标）时才递减 matched           |
| 遇到非法词清空重置            | 非法词之前积累的单词全部无效，不能继续累加                                   |
| 找到结果后左移一个单词          | 保持窗口继续滑动，注意 leftWord 的取值要在 left 更新之前                    |

## 常见错误

| 错误                                               | 后果                   |
| ------------------------------------------------ | -------------------- |
| while 只 `left += wordLen` 不更新 window             | 死循环                  |
| while 默认 left 位置 = cur 直接 `window.put(cur, ...)` | 删词错误，window 数据错乱     |
| 找到结果后不更新 matched                                 | 下一轮 matched 偏大，漏判或误判 |
| 遇到非法词只 continue 不清空窗口                            | 非法词前数据污染后续匹配         |
| 循环条件 `right < s.length()`                        | 少处理最后一个单词位置          |

---
