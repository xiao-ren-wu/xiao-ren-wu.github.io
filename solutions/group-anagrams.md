# 字母异位词分组

## 题目描述

给你一个字符串数组，请你将字母异位词组合在一起。字母异位词指字母相同但排列不同的字符串。

**示例 1：**

```
输入：strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出：[["bat"],["nat","tan"],["ate","eat","tea"]]
```

**示例 2：**

```
输入：strs = [""]
输出：[[""]]
```

**示例 3：**

```
输入：strs = ["a"]
输出：[["a"]]
```

**提示：**

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` 仅包含小写字母

## 解题思路

### 核心思想：哈希表 + 排序

异位词排序后得到相同的字符串，以此作为 key 分组。

```
Map<String, List<String>> map
for str in strs:
  chars = str.toCharArray()
  Arrays.sort(chars)
  key = new String(chars)
  map.computeIfAbsent(key, k -> new ArrayList<>()).add(str)

return new ArrayList(map.values())
```

### 复杂度分析

- **时间复杂度**：O(n × k log k)，k 为字符串最大长度
- **空间复杂度**：O(n × k)

## 代码

```java
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();

        for (String s : strs) {
            char[] chars = s.toCharArray();
            Arrays.sort(chars);
            String key = new String(chars);
            map.computeIfAbsent(key, k -> new ArrayList<>()).add(s);
        }

        return new ArrayList<>(map.values());
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| 计数也可替代排序 | 用长度为 26 的 int[] 计数作为 key，O(n×k) |
| key 唯一性 | 排序后的字符串能唯一标识异位词 |
| 空串处理 | 空串排序后仍是空串，可作为 key |
