---
title: Leetcode 1531. String Compression II
date: 2020-07-31 00:38:54
categories:
- Algorithm
tags:
- leetcode
---

## 题目

https://leetcode.com/problems/string-compression-ii/

字符串缩写：比如`"aaabbbcdddddddddd"`可以缩写为`"a3b3cd10"`。只有一个字母的话不需要写1。数字按十进制。

给定一个字符串s。你可以最多删除s中的k个字符，要使删除之后的s经过缩写之后长度最短。求这个最短长度。

## 做法

### 我的智障做法

`f[i][j][p][q]`表示前i个字符，删去j个，最后一个字符是p（连续q个），这样的最短长度。转移时讨论下一个位置保不保留即可。状态有$100\times100\times26\times100$个，转移O(1)。交上去TLE，优化了优化过了。

### 好的做法

显然，我们只需要关心最后一段就行了。设f[i][j]表示前i个字符，删去j个的最短长度。然后枚举k，我们需要让k到j变成同一段。因此记录一下k到j中出现次数最多的那个，其他字符都要删掉。次数最多的那个的话，倒着循环k，顺便维护即可。时间复杂度O(n^3)。

## 我的代码

https://leetcode.com/submissions/detail/373673723/