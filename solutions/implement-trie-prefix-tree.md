# 实现 Trie（前缀树）

## 题目描述

Trie（发音类似 "try"）或者说前缀树是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。实现 Trie 类：

```
Trie() 初始化前缀树
void insert(String word)  向前缀树中插入单词 word
boolean search(String word)  返回 word 是否在前缀树中
boolean startsWith(String prefix)  返回是否存在以 prefix 为前缀的单词
```

**示例：**

```
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");    // 返回 true
trie.search("app");      // 返回 false
trie.startsWith("app");  // 返回 true
trie.insert("app");
trie.search("app");      // 返回 true
```

## 解题思路

### 核心思想：每个节点包含 26 个子节点指针 + isEnd 标记

```
class TrieNode:
  TrieNode[] children = new TrieNode[26]
  boolean isEnd

insert(word):
  逐字符遍历，不存在则创建新节点，末尾标记 isEnd = true

search(word):
  逐字符遍历，如果字符路径不存在返回 false，最后返回 isEnd

startsWith(prefix):
  逐字符遍历，如果路径不存在返回 false，遍历完返回 true
```

### 复杂度分析

- **时间复杂度**：insert/search/startsWith 均为 O(L)，L 为单词长度
- **空间复杂度**：O(总字符数 × 26)

## 代码

```java
class Trie {
    class TrieNode {
        TrieNode[] children = new TrieNode[26];
        boolean isEnd;
    }

    private TrieNode root;

    public Trie() {
        root = new TrieNode();
    }

    public void insert(String word) {
        TrieNode cur = root;
        for (char c : word.toCharArray()) {
            int idx = c - 'a';
            if (cur.children[idx] == null) cur.children[idx] = new TrieNode();
            cur = cur.children[idx];
        }
        cur.isEnd = true;
    }

    public boolean search(String word) {
        TrieNode cur = root;
        for (char c : word.toCharArray()) {
            int idx = c - 'a';
            if (cur.children[idx] == null) return false;
            cur = cur.children[idx];
        }
        return cur.isEnd;
    }

    public boolean startsWith(String prefix) {
        TrieNode cur = root;
        for (char c : prefix.toCharArray()) {
            int idx = c - 'a';
            if (cur.children[idx] == null) return false;
            cur = cur.children[idx];
        }
        return true;
    }
}
```

## 关键细节/常见错误

| 注意点 | 说明 |
|--------|------|
| children 是 26 长度数组 | 通过 `c - 'a'` 映射到 0-25 |
| isEnd 标记 | 只有 insert 的最后一个节点才标记（区分单词和前缀） |
| 通配符/删除 | 标准 Trie 不要求实现删除；通配符搜索需递归 |
