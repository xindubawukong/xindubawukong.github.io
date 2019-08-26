---
title: Kick Start 2019 Round E
toc: true
date: 2019-08-26 21:05:27
categories:
- Algorithm
tags:
- google
---

https://codingcompetitions.withgoogle.com/kickstart/archive/2019

这场最后40分钟才想起来参加。。。过了T1，然后T3写错了只过了小数据。太菜了。这轮的题相对来说很简单，AK的人也不少。

## T1 Cherries Mesh (9pts, 13pts)

### 题意

有一个n个点的完全图，其中给定的m条边的长度为1，其他边长度为2。要删掉一些边把这个图变成一棵树。求树上边的长度之和的最小值。n, m <= 10^5。

### 做法

显然，只需知道m条边时点的联通情况，联通块内的用长度为1的连接，联通块之间用长度为2的连接即为答案。用并查集维护即可。时间复杂度O(n+m)。

```C++
#include <iostream>
#include <vector>
#include <map>

using namespace std;

vector<int> fa(101111);

int GetFather(int x) {
  if (fa[x] == x) return x;
  return fa[x] = GetFather(fa[x]);
}

void Union(int x, int y) {
  x = GetFather(x);
  y = GetFather(y);
  fa[x] = y;
}

int main() {
  int T;
  cin >> T;
  for (int test = 1; test <= T; test++) {
    int N, M;
    cin >> N >> M;
    for (int i = 1; i <= N; i++) fa[i] = i;
    for (int i = 1; i <= M; i++) {
      int x, y;
      cin >> x >> y;
      Union(x, y);
    }
    map<int, bool> go;
    for (int i = 1; i <= N; i++) go[GetFather(i)] = true;
    int num = go.size();
    int ans = N - num + (num - 1) * 2;
    cout << "Case #" << test << ": " << ans << endl;
  }
  return 0;
}
```

## T2 Code-Eat Switcher (11pts, 25pts)

### 题意

有n个时间段，每个长度为1，在时间段i做x可以获得ai的价值（事情x的价值），做y可以获得bi的价值（事情y的价值）。每个时间段可以按在[0, 1]之间划分，即可以一段做x，一段做y，获得的价值即乘以相应的比率。m个询问A, B，每次询问是否存在一种安排方式使得x的价值不少于A，且y的价值不少于B。n,m<=10^5。

### 做法

显然，一般来说，某个时间段最好只干一种事情，获得的收益才最大。除了最后一个事情需要a和b划分一下之外。按照ai/bi排序，那么越考前的时间段做x比较好，越靠后的时间段做y比较好。因此维护一个a的前缀和，一个b的后缀和，每次查询A和B，在其中二分查找A的位置，将最后一个按比例划分，然后判断后面的B够不够即可。时间复杂度O(nlogn)。

```C++
#include <cstdio>
#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

const double eps = 1e-6;

int D, S;
vector<pair<int, int>> slots, days;

string Work() {
  sort(slots.begin(), slots.end(), [](const pair<int, int>& a, const pair<int, int>& b) {
    //  a.x / a.y > b.x / b.y
    return a.first * b.second > b.first * a.second;
  });
  vector<int> sum_x(S);
  sum_x[0] = slots[0].first;
  for (int i = 1; i < S; i++) {
    sum_x[i] = sum_x[i - 1] + slots[i].first;
  }
  vector<int> rev_sum_y(S);
  rev_sum_y[S - 1] = slots[S - 1].second;
  for (int i = S - 2; i >= 0; i--) {
    rev_sum_y[i] = rev_sum_y[i + 1] + slots[i].second;
  }
  string ans;
  for (const pair<int, int>& day : days) {
    int X = day.first, Y = day.second;
    int l = 0, r = S - 1, res = -1;
    while (l <= r) {
      int mid = (l + r) / 2;
      if (sum_x[mid] >= X) {
        res = mid;
        r = mid - 1;
      }
      else {
        l = mid + 1;
      }
    }
    if (res == -1) {
      ans += 'N';
      continue;
    }
    int x_now = (res == 0) ? 0 : sum_x[res - 1];
    int y_now = (res == S - 1) ? 0 : rev_sum_y[res + 1];
    double x_ratio = double(X - x_now) / slots[res].first;
    double y_ratio = 1 - x_ratio;
    if (y_now + y_ratio * slots[res].second > Y - eps) {
      ans += 'Y';
    }
    else {
      ans += 'N';
    }
  }
  return ans;
}

int main() {
  int T;
  cin >> T;
  for (int test = 1; test <= T; test++) {
    cin >> D >> S;
    slots.resize(S);
    for (int i = 0; i < S; i++) {
      cin >> slots[i].first >> slots[i].second;
    }
    days.resize(D);
    for (int i = 0; i < D; i++) {
      cin >> days[i].first >> days[i].second;
    }
    string ans = Work();
    cout << "Case #" << test << ": " << ans << endl;
  }
  return 0;
}
```

## T3

### 题意

给定区间[L, R]，求区间中满足以下条件的x的个数：x的偶数因数的个数与x的奇数因数的个数之差的绝对值<=2。L, R<= 10^9 且 R - L<=10^5。

### 做法

因为区间长度只有10^5，故枚举区间中的每一个数判断即可。设$x=p_1^{q_1}\times p_2^{q_2}\times...\times p_k^{q_k}$，那么x的因数个数即为$(q_1+1)\times(q_2+1)\times...\times(q_k+1)$。而奇数因数的个数，只需将2的那个指数项去掉即可。偶数因数的个数用总因数个数减奇数因数个数得到。因此，首先预处理0至$\sqrt{R}$之间的质数，然后分解区间中的每一个x判断即可。

比赛时我这样交上TLE。我们可以计算一下，$\sqrt{10^9}\approx35000$，小于它的质数个数约有$\frac{35000}{ln(35000)}\approx3000$个。由于区间长度为100000，乘以3000确实有点大。

如何优化呢？由于我们只需要判断奇偶因数只差的绝对值。我们把2的次数单独拿出来，即$x=2^a\times p_1^{q_1}\times p_2^{q_2}...\times p_k^{q_k}$。其中$a$可能等于0。设$(q_1+1)\times(q_2+1)\times...\times(q_k+1)=M$，那么总因数个数即为$(a+1)\times M$，奇数因数个数为$M$，偶数因数个数为$a\times M$，两者只差的绝对值为$|a-1|\times M$。

-   如果$a=0$，即x是一个奇数，两者之差的绝对值即为M。
-   如果$a=1$，结果为0，一定满足条件。
-   如果$a\geq 2$，结果为$(a-1)\times M$。

以上结果可以看出，我们并不需要知道$M$是几，只需要知道$M$是不是大于2就可以了！所以在质因数分解时，如果$M>2$就break。加入这个优化后就可以通过了。

```C++
#include <iostream>
#include <vector>

using namespace std;

vector<int> primes;

void Prepare() {
  vector<bool> is_prime(40000);
  is_prime[2] = true;
  for (int i = 3; i < 40000; i++) is_prime[i] = true;
  for (int i = 2; i < 40000; i++) {
    if (is_prime[i]) {
      for (int j = i + i; j < 40000; j += i) {
        is_prime[j] = false;
      }
    }
  }
  for (int i = 3; i < 40000; i++) {
    if (is_prime[i]) primes.push_back(i);
  }
}

void Factorize(int x, int& a, int& M) {
  a = 0;
  while ((x & 1) == 0) {
    x >>= 1;
    a++;
  }
  int now = x;
  M = 1;
  for (int prime : primes) {
    if (prime * prime > x) break;
    if (now % prime == 0) {
      int cnt = 0;
      while (now % prime == 0) {
        now /= prime;
        cnt++;
      }
      M *= (cnt + 1);
      if (M > 2) break;
    }
  }
  if (now > 1) M *= 2;
}

bool Check(int x) {
  int a, M;
  Factorize(x, a, M);
  if (a == 0) {
    return M <= 2;
  }
  else if (a == 1) {
    return true;
  }
  else {
    return (a - 1) * M <= 2;
  }
}

int Work(int L, int R) {
  int ans = 0;
  for (int i = L; i <= R; i++) {
    if (Check(i)) ans++;
  }
  return ans;
}

int main() {
  Prepare();
  int T;
  cin >> T;
  for (int test = 1; test <= T; test++) {
    int L, R;
    cin >> L >> R;
    int ans = Work(L, R);
    cout << "Case #" << test << ": " << ans << endl;
  }
  return 0;
}
```