---
title: Codeforces Global Round 9 D. Replace by MEX
date: 2020-07-09 18:44:42
categories:
- Algorithm
tags:
- codeforces
---

## 题目

http://codeforces.com/contest/1375/problem/D

有一个长度为n的数组a，其中a[i]$\in$[0, n]。你可以进行一种操作：将某个位置的数字替换为整个数组的MEX。

MEX定义为：The MEX (minimum excluded) of an array is the smallest non-negative integer that does not belong to the array.

因此，n个0到n之间的数，其MEX也一定在0到n之间。

需要找到一个长度不超过2n的操作序列，使数组变成非降的。

## 做法

我们的目标就是让数组变为0, 1, 2... n-1。显然，我们可以求MEX，然后如果MEX在[0, n-1]之间，那么直接将对应位置的数替换，这样就多了一个位置的数字是正确的。如果MEX是n怎么办呢？由于要求的是长度不超过2n，所以可以加一次操作，将任意一个不正确的位置替换为n，然后再重复上面的流程。

如果MEX是x，位置x一定是不正确的吗？是的，因为x一定不再当前数组里。

如果每次暴力求MEX，那么时间复杂度是O(n^2)。

其实，如果MEX是n的话，那么此时数组是一个0到n-1的排列。于是，我们可以找到其中所有的置换圈。对于每一个圈，设大小为s，那么我们只需用n进去换一遍，就能得到正确的位置，所需次数为s+1。这样最坏只需1.5n次即可（即所有的圈大小为2）。

## 代码

http://codeforces.com/contest/1375/submission/86274403