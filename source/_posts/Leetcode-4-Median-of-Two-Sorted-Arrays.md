---
title: Leetcode 4. Median of Two Sorted Arrays
date: 2019-09-06 18:45:39
categories:
- Algorithm
tags:
- leetcode
---

## 题目

https://leetcode.com/problems/median-of-two-sorted-arrays/

给定两个有序数组，求两个数组中的所有数字的中位数。

## 做法

二分查找，分情况讨论。参见代码注释。

```C++

class Solution {
public:
  // 在一个数组中的第K小
  int GetKthFromArray(const vector<int>& a, int l, int r, int K) {
    return a[l + K - 1];
  }
  // 在一个数组和一个额外的数（另一个数组只有一个数）中的第K小。
  int GetKthFromArrayAndNumber(const vector<int>& a, int l, int r,
                               int extra_number, int K) {
    if (l + K - 1 > r) return max(a[r], extra_number);
    if (a[l + K - 1] < extra_number) return a[l + K - 1];
    if (l + K - 2 >= l) return max(a[l + K - 2], extra_number);
    else return extra_number;
  }

  // 两个数组中的第K小
  int GetKth(const vector<int>& a, int l1, int r1,
             const vector<int>& b, int l2, int r2, int K) {
    if (l1 > r1) return GetKthFromArray(b, l2, r2, K);
    if (l2 > r2) return GetKthFromArray(a, l1, r1, K);
    if (l1 == r1) return GetKthFromArrayAndNumber(b, l2, r2, a[l1], K);
    if (l2 == r2) return GetKthFromArrayAndNumber(a, l1, r1, b[l2], K);
    int m1 = (l1 + r1) / 2, cnt1 = m1 - l1 + 1;
    int m2 = (l2 + r2) / 2, cnt2 = m2 - l2 + 1;
    if (K <= cnt1) return GetKth(a, l1, m1, b, l2, r2, K);  // a的后一半不用考虑
    if (K <= cnt2) return GetKth(a, l1, r1, b, l2, m2, K);  // b的后一半不用考虑
    if (K <= cnt1 + cnt2) {  // 去掉一个数组的后一半
      if (a[m1] > b[m2]) return GetKth(a, l1, m1, b, l2, r2, K);
      else return GetKth(a, l1, r1, b, l2, m2, K);
    }
    else {  // 去掉一个数组的前一半，同时维护K
      if (a[m1] > b[m2]) return GetKth(a, l1, r1, b, m2 + 1, r2, K - cnt2);
      else return GetKth(a, m1 + 1, r1, b, l2, r2, K - cnt1);
    }
  }

  double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
    int n1 = nums1.size();
    int n2 = nums2.size();
    int tot = n1 + n2;
    if (tot % 2 == 1) {
      return GetKth(nums1, 0, n1 - 1, nums2, 0, n2 - 1, tot / 2 + 1);
    }
    else {
      int t1 = GetKth(nums1, 0, n1 - 1, nums2, 0, n2 - 1, tot / 2);
      int t2 = GetKth(nums1, 0, n1 - 1, nums2, 0, n2 - 1, tot / 2 + 1);
      return (t1 + t2) * 0.5;
    }
  }
};
```

时间复杂度O(log(n+m))