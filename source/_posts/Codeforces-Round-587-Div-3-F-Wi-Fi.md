---
title: 'Codeforces Round #587 (Div. 3) F. Wi-Fi'
date: 2019-10-01 05:03:37
categories:
- Algorithm
tags:
- codeforces
- DP
---

## 题目

https://codeforces.com/contest/1216/problem/F

有n个房间，编号[1,n]，路由器范围为k。现要给所有房间通wifi。对于房间i，有两种方案：
- 直连网线，话费为i
- 在i建路由器，话费也是i，但这样[i-k, i+k]区间就都连通了
有些房间不能建路由器。求最少花费。

## 做法

首先，能建基站的一定不会直连。由于基站之间会相互覆盖，所以只需决定哪些保留哪些基站。感觉像是贪心？不知道能不能做。我是用动态规划做的。设f[i]表示只用[1,i]之间的基站，将[1,i]都连通的最小话费。那么有两种情况：
- i直连，花费为f[i-1]+i
- 在某位置j建基站(i-k<=j<=i)，这样[j-k,i]之间就都已经覆盖了，只需前面的可以就行。我们需要将前面的也满足。此时的花费为j+min(f[t], j-k-1<=t<j)。

由于k给定，所以这两个最小值（j和t）都可以用单调队列来维护。DP即可，时间复杂度O(n)。

```c++
#include <iostream>
#include <algorithm>
#include <vector>
#include <deque>

using namespace std;

long long Work(int N, int K, const vector<bool>& can) {
  vector<long long> f(N + 1);
  deque<pair<int, long long>> q_f;  // min f in [i - K - 1, i - 1]
  deque<pair<int, long long>> q_g;  // min g in [i - K, i]
  q_f.push_back(make_pair(0, 0));
  
  for (int i = 1; i <= N; i++) {
    while (!q_f.empty() && q_f.front().first < i - K - 1) q_f.pop_front();

    if (can[i]) {
      long long min_f = q_f.front().second;
      long long min_g = min_f + i;
      while (!q_g.empty() && q_g.back().second >= min_g) q_g.pop_back();
      q_g.push_back(make_pair(i, min_g));
    }
    while (!q_g.empty() && q_g.front().first < i - K) q_g.pop_front();

    f[i] = f[i - 1] + i;
    if (!q_g.empty()) f[i] = min(f[i], q_g.front().second);

    while (!q_f.empty() && q_f.back().second >= f[i]) q_f.pop_back();
    q_f.push_back(make_pair(i, f[i]));
  }
  return f[N];
}

int main() {
  int N, K;
  cin >> N >> K;
  vector<bool> can(N + 1);
  for (int i = 1; i <= N; i++) {
    char ch;
    cin >> ch;
    while (ch != '0' && ch != '1') cin >> ch;
    can[i] = (ch == '1');
  }
  cout << Work(N, K, can) << endl;
  return 0;
}
```