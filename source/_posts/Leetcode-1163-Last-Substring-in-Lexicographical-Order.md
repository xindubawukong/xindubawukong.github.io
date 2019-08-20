---
title: Leetcode 1163. Last Substring in Lexicographical Order
date: 2019-08-20 19:11:30
categories:
- Algorithm
tags:
- leetcode
---

## 题目

https://leetcode.com/problems/last-substring-in-lexicographical-order/

给定一个字符串s，求s的所有子串中字典序最大的那一个。s长度<=400000。

## 分析

首先可以发现，s[l, n-1]的字典序一定比s[l, r]大。因此我们只需在s的n个后缀中找字典许最大的即可。

说到后缀自然我当时想到的是后缀数组。倍增求后缀数组可以参考https://blog.csdn.net/Sunshine_cfbsl/article/details/52190433，时间复杂度O(nlogn)。这种方法我交上去超时了。不过代码还是放一下：

```C++
class SuffixArray {
public:
  int n;
  vector<int> rank, sa, height;

  void work(const vector<pair<pair<int, int>, int>>& a) {
    vector<vector<pair<pair<int, int>, int>>> go(n);
    vector<pair<pair<int, int>, int>> now;
    for (auto tmp : a) {
      if (tmp.first.second == -1) now.push_back(tmp);
      else go[tmp.first.second].push_back(tmp);
    }
    for (int i = 0; i < n; i++) {
      for (auto tmp : go[i]) now.push_back(tmp);
    }
    for (int i = 0; i < n; i++) go[i].clear();
    for (auto tmp : now)
      go[tmp.first.first].push_back(tmp);
    now.clear();
    for (int i = 0; i < n; i++) {
      for (auto tmp : go[i]) now.push_back(tmp);
    }
    rank[now[0].second] = 0;
    for (int i = 1, p = 0; i < n; i++) {
      if (now[i].first != now[i-1].first) p++;
      rank[now[i].second] = p;
    }
  }

  void GetSuffixArray(const string& s) {
    n = s.length();
    rank.resize(n);
    sa.resize(n);
    height.resize(n);
    vector<pair<pair<int, int>, int>> a(n);
    for (int i = 0; i < n; i++)  a[i] = make_pair(make_pair(s[i], -1), i);
    sort(a.begin(), a.end());
    rank[a[0].second] = 0;
    for (int i = 1, p = 0; i < n; i++) {
      if (a[i].first != a[i-1].first) p++;
      rank[a[i].second] = p;
    }
    for (int k = 1; k < n; k <<= 1) {
      for (int i = 0; i < n; i++) a[i] = make_pair(make_pair(rank[i], i + k >= n ? -1 : rank[i + k]), i);
      work(a);
    }
    for (int i = 0; i < n; i++) sa[rank[i]] = i;
    for (int i = 0, p = 0; i < n; i++)  {
      if (rank[i] == 0) continue;
      int a = i, b = sa[rank[i] - 1];
      while (a + p < n && b + p < n && s[a + p] == s[b + p]) p++;
      height[rank[i]] = p;
      if (p > 0) p--;
    }
  }
};

class Solution {
public:
    string lastSubstring(string s) {
        SuffixArray sa;
        sa.GetSuffixArray(s);
        int t = sa.sa[sa.n - 1];
        return s.substr(t);
    }
};
```

后缀数组写的很傻，尤其是基数排序那里。实在是没法看。。。

那么有没有O(n)的方法呢？我想了一会琢磨了一个乱七八糟的算法。首先，我们可以把所有第一个字母最大的后缀作为candidates。然后，比较candidates的第二个字母，把第二个字母也最大的放到新的candidates中。这样一直比较，知道candidates中只剩下一个后缀即可。

但是这种方法时间上不能保证，比如"aaaaaa..."，就会变成O(n^2)。这时需要利用一下candidates的性质：candidates中的后缀一定比不在candidates中的后缀大。对于candidates中的两个后缀i和j，i < j，且他们前k个字母都相同，s[i, i+k-1] = s[j, j+k-1]，如果i+k-1=j，也就是说长度k使得两个后缀重叠了，会发生很神奇的现象。由于我们需要比较的是s[i, n-1]和s[j, n-1]这两个后缀，由于其前k个相同，那么等价于比较s[i+k-1, n-1]和s[j+k-1, n-1]这两个后缀，也就是比较s[j, n-1]和s[j+k-1, n-1]。由于前面说过，candidates中的后缀一定比不在candidates中的后缀大，所以当j+k-1不是candidates时，s[j, n-1] > s[j+k-1, n-1]，也就是s[i, n-1] > s[j, n-1]。那么j就可以从candidates中删除。如果j+k-1也在candidates中，那么套用上面的理论，j+k-1也可以删除，从而j也可以删除。也就是说，如果我们在比较所有candidates的第k个字母时，如果发现两个candidates出现了重叠，那么前面的就可以把后面的那个吃掉。这样可以做了。代码如下：

```C++
class Solution {
public:
  string lastSubstring(string s) {
    int n = s.size();
    vector<bool> flag(n);
    vector<int> num(26);
    for (int i = 0; i < n; i++) num[s[i] - 'a']++;
    vector<int> now;
    for (int k = 25; k >= 0; k--) {
      if (num[k] > 0) {
        for (int i = 0; i < n; i++) {
          if (s[i] - 'a' == k) {
            flag[i] =true;
            now.push_back(i);
          }
        }
        break;
      }
    }
    for (int len = 2; len <= n; len++) {
      if (now.size() == 1) break;
      for (int i = 0; i < 26; i++) num[i] = 0;
      for (int pos : now) {
        int t = pos + len - 1;
        if (t >= n) continue;
        num[s[t] - 'a']++;
      }
      for (int k = 25; k >= 0; k--) {
        if (num[k] > 0) {
          for (int pos : now) {
            int t = pos + len - 1;
            if (t >= n) {
              flag[pos] = false;
              continue;
            };
            flag[t] = false;
            if (s[t] - 'a' != k) flag[pos] = false;
          }
          break;
        }
      }
      vector<int> new_now;
      for (int pos : now) {
        if (flag[pos]) new_now.push_back(pos);
      }
      now = new_now;
    }
    return s.substr(now[0]);
  }
};
```

下面分析复杂度（其实做的时候没想复杂度）。对于开始时的所有candidates，它最多会被遍历d次，d为它之前的那个candidates的距离（先不考虑第一个candidates）。这些距离之和不超过n，再加上第一个candidates最多被遍历n次，故时间复杂度为O(n)。