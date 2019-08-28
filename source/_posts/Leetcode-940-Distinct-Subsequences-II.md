---
title: Leetcode 940. Distinct Subsequences II
date: 2019-08-28 22:16:17
categories:
- Algorithm
tags:
- leetcode
---

## 题目

https://leetcode.com/problems/distinct-subsequences-ii/

给定一个小写字母组成的字符串s，求s的不同子序列的个数模P。

## 做法

设f[i]表示s的前i个字符构成的不同子序列个数。那么对于第i个，我们可以选第i个，也可以不选第i个，所以f[i] = f[i-1] * 2。这样肯定是有问题的，会出现重复计算的情况。什么时候会重复呢？首先，选第i个，结果确实是f[i-1]，不选第i个，结果也确实是f[i-1]。问题就在于这两个f[i-1]，选第i个和不选第i个可能会出现同样的子序列。那么重复的最后一定是s[i]。设上一个s[i]字母出现在k位置，那么f[k-1]这些串，后面跟s[k]和s[i]，结果在两个f[i-1]中被算了两次！所以减去f[k-1]即可。时间复杂度O(n)。

```C++
const int P = 1000000007;
int last[26];
int f[2111];

class Solution {
public:
    int distinctSubseqII(string S) {
        for (int i = 0; i < 26; i++) last[i] = -1;
        f[0] = 2;
        last[S[0] - 'a'] = 0;
        for (int i = 1; i < S.length(); i++) {
            f[i] = f[i-1] * 2 % P;
            if (last[S[i] - 'a'] >= 0) {
                f[i] -= last[S[i] - 'a'] == 0 ? 1 : f[last[S[i] - 'a'] - 1];
                if (f[i] < 0) f[i] += P;
            }
            last[S[i] - 'a'] = i;
        }
        int ans = f[S.length() - 1] - 1;
        if (ans < 0) ans += P;
        return ans;
    }
};
```
