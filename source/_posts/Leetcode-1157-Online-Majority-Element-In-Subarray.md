---
title: Leetcode 1157. Online Majority Element In Subarray
toc: true
date: 2019-08-14 23:50:41
categories:
- Algorithm
tags:
- leetcode
---
## 题目
https://leetcode.com/problems/online-majority-element-in-subarray/

给定一个长度为n的数组，m次查询，每次查询[l, r]区间内出现次数>=threshold的数是哪个。保证 threshold * 2 > r - l + 1 。强制在线。n, m <= 20000。

## 做法
暴力做法：对于每个query，扫描l到r区间中的数，用hash表存下每个数的出现次数。时间复杂度O(n*m)。

这个题中要求的数是majority，也就是出现次数大于总次数的一半。这是一个很强的限制。我们可以发现，暴力算法处理慢的地方在于区间较长的时候，也就是threshold较大的时候。而threshold较大时，满足条件的数的个数就会很少！

于是可以取$\sqrt{n}$作为阈值，分两种情况讨论：
- 若threshold<$\sqrt{n}$，则区间长度<$2\sqrt{n}$，按照暴力做法求解，每次询问时间复杂度O($\sqrt{n}$)。
- 若threshold>$\sqrt{n}$，满足这样条件的数不超过$\frac{n}{\sqrt{n}}=\sqrt{n}$个。因此在开始时进行预处理，求出这$\sqrt{n}$个数，那么只需要在查询时知道每一个数在[l, r]中出现了多少次。我的方法是对每个数维护一个长度为n的前缀和数组，这样O(1)的时间即可得到其在[l, r]中的出现次数，每次询问的时间复杂度为O($\sqrt{n}$)。

总时间复杂度O($n\sqrt{n}$)。代码如下：
```C++
class MajorityChecker {
public:
    
    vector<int> arr;
    vector<int> frequent;
    vector<vector<int>> sum;
    
    MajorityChecker(vector<int>& arr) {
        this->arr = arr;
        map<int, int> count;
        for (int x : arr) {
            count[x]++;
        }
        for (auto it = count.begin(); it != count.end(); it++) {
            if (it->second > 100) {
                frequent.push_back(it->first);
            }
        }
        sum.resize(frequent.size());
        for (int i = 0; i < frequent.size(); i++) {
            int x = frequent[i];
            sum[i].resize(arr.size());
            for (int j = 0; j < arr.size(); j++) {
                if (x == arr[j]) {
                    sum[i][j] = 1;
                }
                if (j > 0) {
                    sum[i][j] += sum[i][j-1];
                }
            }
        }
    }
    
    int work1(int left, int right, int threshold) {
        unordered_map<int, int> count;
        for (int i = left; i <= right; i++) {
            if (++count[arr[i]] >= threshold) {
                return arr[i];
            }
        }
        return -1;
    }
    
    int work2(int left, int right, int threshold) {
        for (int i = 0; i < frequent.size(); i++) {
            int cnt = sum[i][right] - (left == 0 ? 0 : sum[i][left - 1]);
            if (cnt >= threshold) {
                return frequent[i];
            }
        }
        return -1;
    }
    
    int query(int left, int right, int threshold) {
        if (threshold <= 100) {
            return work1(left, right, threshold);
        }
        else {
            return work2(left, right, threshold);
        }
    }
};
```
在Google被人逼着养成了良好的代码风格，好处就是代码比较容易看懂，坏处就是写的过程非常痛苦。
## 其他人的做法
交了之后看了一下discuss（这个题暂时还没有solution），发现有些人用线段树做的。这种做法基于<a href="https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm">摩尔投票算法</a>。这个算法专门用来求数组中的majority，也就是出现次数大于一半的数。伪代码如下：
```
Initialize an element m and a counter i with i = 0
For each element x of the input sequence:
    If i = 0, then assign m = x and i = 1
    else if m = x, then assign i = i + 1
    else assign i = i − 1
Return m
```
看完大概就能明白，这个算法只能做找majority这么一件事情，并且如果没有majority它也会返回一个值，需要再check一下结果是不是majority。

这个题正好是要找majority。不过，我没怎么想清楚他们是怎么用线段树维护这个东西的。<a href="https://leetcode.com/problems/online-majority-element-in-subarray/discuss/358338/SegTree-C%2B%2B-O(lg2)">他们的代码</a>写的我觉得有点问题，但是竟然能过。这种方法的时间复杂度是O($n\log_{2}^{2}{n}$)，实际跑起来比我的做法还要慢一些。