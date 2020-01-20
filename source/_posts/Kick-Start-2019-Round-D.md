---
title: Kick Start 2019 Round D
toc: true
date: 2019-09-01 15:18:46
categories:
- Algorithm
tags:
- google
---

https://codingcompetitions.withgoogle.com/kickstart/archive/2019

这场的题比较好玩。

## T1: X or What? (8pts, 19pts)

### 题目

给定长度为N的初始int数组A。定义有趣区间[L, R]：A[L] xor A[L+1] xor ... xor A[R]结果的二进制具有偶数个1。M个操作，每次修改A中的一个数。求每次操作后A中最长的有趣区间的长度。N, M <= 10^5。

### 做法

一开始我想用前缀和维护区间异或值，但是貌似没法做。稍微思考一下可以发现，我们要求的有趣区间比较有趣。如果a xor b有偶数个1，因为xor的性质，要么a和b都有偶数个1，要么都有奇数个1。换句话说，我们先把所有数都替换成0和1即可。那么有趣区间中一定包含偶数个1。如果A中有偶数个1，那么[0, N-1]就是答案。否则答案就是0到倒数第二个1，或者第二个1到N-1。用一个set维护所有1的位置即可。时间复杂度O(nlogn)。

我做的时候比较sb，用一个线段树维护0和1的位置。代码就不贴了。

## T2: Latest Guests (12pts, 23pts)

### 题目

这个题比较好玩。有N个在一个圆圈上的城市，编号1..N，N与1相邻。有G个人，第i个人从Hi城市出发，方向是Di（顺时针或逆时针），每一分钟就走到下一个城市。这样过了M分钟。问每一个人是多少个城市的最后访问者。如果一个城市最后有好多人访问，那么这些人都是最后访问者。N, G <= 10^5，M <= 10^9。

### 做法

我们先考虑顺时针的情况，显然可以求出每个店最后访问的是谁（只会有一个，从最后的位置倒着找即可），然后求出这个时间。对逆时针的人也同样处理，那么比较两个时间哪个靠后即可。时间复杂度O(nlogn)。

```C++
#include <iostream>
#include <algorithm>
#include <vector>
#include <set>

using namespace std;

int GetSum(const vector<int>& sum, int l, int r) {
  if (l <= r) {
    return sum[r] - (l == 0 ? 0 : sum[l - 1]);
  }
  return sum[r] + sum[sum.size() - 1] - sum[l - 1];
}

void Work(int N, int G, int M, const vector<pair<int, bool>>& guests, vector<int>* ans) {
  set<int> clock_end, anti_clock_end;
  for (int i = 0; i < G; i++) {
    if (guests[i].second) {
      int end = (guests[i].first + M) % N;
      clock_end.insert(end);
    }
    else {
      int end = (guests[i].first - M) % N;
      if (end < 0) end += N;
      anti_clock_end.insert(end);
    }
  }
  vector<int> clock_last_time(N), anti_clock_last_time(N);
  for (int i = 0; i < N; i++) {
    clock_last_time[i] = -1;
    anti_clock_last_time[i] = -1;
  }
  for (auto it = clock_end.begin(); it != clock_end.end(); it++) {
    int last;
    if (it == clock_end.begin()) {
      last = *clock_end.rbegin();
    }
    else {
      auto tmp = it;
      tmp--;
      last = *tmp;
    }
    int now = M;
    for (int pos = *it; pos != (last + 1) % N; pos = (pos + N - 1) % N) {
      clock_last_time[pos] = now;
      now--;
      if (now < 0) break;
    }
    if (now >= 0) clock_last_time[(last + 1) % N] = now;
  }
  for (auto it = anti_clock_end.begin(); it != anti_clock_end.end(); it++) {
    int next;
    if ((*it) == (*anti_clock_end.rbegin())) {
      next = *anti_clock_end.begin();
    }
    else {
      auto tmp = it;
      tmp++;
      next = *tmp;
    }
    int now = M;
    for (int pos = *it; pos != (next - 1 + N) % N; pos = (pos + 1) % N) {
      anti_clock_last_time[pos] = now;
      now--;
      if (now < 0) break;
    }
    if (now >= 0) anti_clock_last_time[(next - 1 + N) % N] = now;
  }
  vector<int> clock_sum(N), anti_clock_sum(N);
  for (int i = 0; i < N; i++) {
    int res = max(clock_last_time[i], anti_clock_last_time[i]);
    if (res != -1 && res == clock_last_time[i]) clock_sum[i] = 1;
    if (res != -1 && res == anti_clock_last_time[i]) anti_clock_sum[i] = 1;
    if (i > 0) {
      clock_sum[i] += clock_sum[i - 1];
      anti_clock_sum[i] += anti_clock_sum[i - 1];
    }
  }
  ans->resize(G);
  for (int i = 0; i < G; i++) {
    if (guests[i].second) {
      int end = (guests[i].first + M) % N;
      auto it = clock_end.find(end);
      int last;
      if (it == clock_end.begin()) {
        last = *clock_end.rbegin();
      }
      else {
        auto tmp = it;
        tmp--;
        last = *tmp;
      }
      int len = min(end == last ? N : (end - last + N) % N, M + 1);
      int start = (end - len + 1 + N) % N;
      (*ans)[i] = GetSum(clock_sum, start, end);
    }
    else {
      int end = (guests[i].first - M) % N;
      if (end < 0) end += N;
      auto it = anti_clock_end.find(end);
      int next;
      if ((*it) == (*anti_clock_end.rbegin())) {
        next = *anti_clock_end.begin();
      }
      else {
        auto tmp = it;
        tmp++;
        next = *tmp;
      }
      int len = min(end == next ? N : (next - end + N) % N, M + 1);
      int to = (end + len - 1 + N) % N;
      (*ans)[i] = GetSum(anti_clock_sum, end, to);
    }
  }
}

int main() {
  int T;
  cin >> T;
  for (int test = 1; test <= T; test++) {
    int N, G, M;
    cin >> N >> G >> M;
    vector<pair<int, bool>> guests(G);
    for (int i = 0; i < G; i++) {
      int x;
      cin >> x;
      guests[i].first = x - 1;
      char ch = getchar();
      while (ch != 'C' && ch != 'A') ch = getchar();
      guests[i].second = (ch == 'C');
    }
    vector<int> ans;
    Work(N, G, M, guests, &ans);
    cout << "Case #" << test << ": ";
    for (int x : ans) cout << x << " ";
    cout << endl;
  }
  return 0;
}
```

## T3: Food Stalls (7pts, 31pts)

### 题目

x轴上有N个点，每个点有坐标Xi和代价Ci。现在要选一个点i建仓库，代价Ci，然后再选K个点建杂货店，每个点代价Cj+|Xi-Xj|。求最小代价。N <= 10^5。

### 做法

首先可以发现，实际上就是要选出K+1个点，将它们的Ci都算上，然后选一个中心点。这个中心点显然一定是中间的那个。所以枚举每个点作为中心点，左边求Ci-Xi的前K/2小，右边求Ci+Xi的前K/2(+1)小即可。用set维护即可。时间复杂度O(NlogN)。

```C++
#include <iostream>
#include <algorithm>
#include <set>
#include <vector>

using namespace std;

const long long inf64 = 1ll << 60;

long long Work(int K, int N, const vector<int>& X, const vector<int>& C) {
  vector<pair<int, int>> spots(N);
  for (int i = 0; i < N; i++) {
    spots[i] = make_pair(X[i], C[i]);
  }
  sort(spots.begin(), spots.end());
  int left_num = K / 2;
  int right_num = K - left_num;
  vector<long long> left(N), right(N);
  multiset<int> now;
  long long sum = 0;
  for (int i = 0; i < N; i++) {
    if (now.size() == left_num) left[i] = sum;
    else left[i] = inf64;
    sum += spots[i].second - spots[i].first;
    now.insert(spots[i].second - spots[i].first);
    if (now.size() > left_num) {
      auto it = --now.end();
      sum -= *it;
      now.erase(it);
    }
  }
  now.clear();
  sum = 0;
  for (int i = N - 1; i >= 0; i--) {
    if (now.size() == right_num) right[i] = sum;
    else right[i] = inf64;
    sum += spots[i].second + spots[i].first;
    now.insert(spots[i].second + spots[i].first);
    if (now.size() > right_num) {
      auto it = --now.end();
      sum -= *it;
      now.erase(it);
    }
  }
  long long res = inf64;
  for (int i = 0; i < N; i++) {
    long long left_cost = left[i] + (long long)left_num * spots[i].first;
    long long right_cost = right[i] - (long long)right_num * spots[i].first;
    res = min(res, left_cost + right_cost + spots[i].second);
  }
  return res;
}

int main() {
  int T;
  cin >> T;
  for (int test = 1; test <= T; test++) {
    int K, N;
    cin >> K >> N;
    vector<int> X(N), C(N);
    for (int i = 0; i < N; i++) cin >> X[i];
    for (int i = 0; i < N; i++) cin >> C[i];
    auto ans = Work(K, N, X, C);
    cout << "Case #" << test << ": " << ans << endl;
  }
  return 0;
}
```