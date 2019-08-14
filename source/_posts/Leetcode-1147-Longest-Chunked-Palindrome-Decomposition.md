---
title: Leetcode 1147. Longest Chunked Palindrome Decomposition
toc: true
date: 2019-08-15 01:59:36
categories:
- Algorithm
tags:
---
## 题目
https://leetcode.com/problems/longest-chunked-palindrome-decomposition/

给一个字符串，要求将其分割成a1, a2...ak，每一份长度>0，并且a[i] = a[k + 1 - i]。例如"ghiabcdefhelloadamhelloabcdefghi"可以分割为"(ghi)(abcdef)(hello)(adam)(hello)(abcdef)(ghi)"。求最大的k。字符串长度<=1000。
## 做法
本来看n<=1000觉得是O($n^2$)的做法，枚举中间一块的长度，然后从中间向两边贪心，判断字符串相等使用hash。后来发现，直接从两边往中间贪心就行了。。连hash都不用，直接暴力判断。这也是hard题。。
```C++
class Solution {
public:
    int longestDecomposition(string text) {
        int n = text.length();
        int ans = 0;
        int l = 0, r = 0;
        while (r * 2 + 1 < n) {
            if (text.substr(l, r - l + 1) == text.substr(n - r - 1, r - l + 1)) {
                ans += 2;
                l = r + 1;
                r = l;
                continue;
            }
            r++;
        }
        if (n % 2 == 0 && l == n / 2) return ans;
        return ans + 1;
    }
};
```
