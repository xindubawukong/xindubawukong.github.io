---
title: Kick Start 2019 Round A
toc: true
date: 2019-08-24 22:50:25
categories:
- Algorithm
tags:
- google
---
https://codingcompetitions.withgoogle.com/kickstart/archive/2019

没参加比赛，拿题来练了一下。感觉这题考场上也就能做前两道。

## T1 Training

### 题目

给N个数，你可以花1块钱给一个数加1。要从中选出K个相同的数。求最少花多少钱。$N<=10^5$。

### 做法

显然要先排序，那么我们要选的K个数一定是排好序后连续的K个。枚举并用前缀和即可。时间复杂度O(nlogn)。

```C++
#include <cstdio>
#include <algorithm>

using namespace std;
const int maxN = 100005;

int T, N, P;
int a[maxN], sum[maxN];

int main() {
    scanf("%d", &T);
    for (int tt = 1; tt <= T; tt++) {
        scanf("%d%d", &N, &P);
        for (int i = 1; i <= N; i++) scanf("%d", &a[i]);
        sort(a+1, a+N+1);
        for (int i=1;i<=N;i++)
            sum[i] = sum[i-1] + a[i];
        int ans = 1000000000 + 1;
        for (int i=1;i<=N;i++) {
            int r = i;
            int l = r - P + 1;
            if (l < 1) continue;
            ans = min(ans, a[r] * P - sum[r] + sum[l-1]);
        }
        printf("Case #%d: %d\n", tt, ans);
    }
    return 0;
}
```

## T2 Parcels

### 题目

给定一个$m\times n$的01矩阵。一个矩阵的分数k定义为，矩阵中的所有点到其最近的1的距离的最大值。这里的距离为曼哈顿距离（|r1 - r2| + |c1 - c2|）。现在可以将矩阵中的某个0改成1，求矩阵分数k的最小值。m, n <= 250。

### 做法

最大值最小 -> 二分答案。二分k，那么只需对于每一个位置(i, j)，距离它k以内是否有1，如果没有的话将其放到一个数组里，最后判断数组中的位置是否可以修改一个1搞定。

那么如何判断距(i, j)最近的1的距离呢？注意到这是曼哈顿距离，所以可以用动态规划。设f[i][j]表示(i, j)的左上角这个矩阵中最近的1的距离。那么f[i][j] = min(f[i-1][j]+1, f[i][j-1]+1)。对四个方向各做一次dp即可。

所以问题就只剩下给定一些点，要判断是否能选择一个位置改成1，使得这些点到该位置的最大值<=k。同样dp一遍即可，不过这次是维护最大值。

时间复杂度O($n^2log_2n$)。如果我们暴力枚举k的话是O($n^3$)，也能通过。

```C++
#include <cstdio>
#include <iostream>
#include <vector>
using namespace std;

#define inf 999999999

int m, n;
bool board[255][255], need[255][255];
int f[4][255][255], g[4][255][255];

void Prepare() {
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      f[0][i][j] = board[i][j] ? 0 : inf;
      if (i > 0) f[0][i][j] = min(f[0][i][j], f[0][i - 1][j] + 1);
      if (j > 0) f[0][i][j] = min(f[0][i][j], f[0][i][j - 1] + 1);
    }
  }
  for (int i = 0; i < m; i++) {
    for (int j = n - 1; j >= 0; j--) {
      f[1][i][j] = board[i][j] ? 0 : inf;
      if (i > 0) f[1][i][j] = min(f[1][i][j], f[1][i - 1][j] + 1);
      if (j < n - 1) f[1][i][j] = min(f[1][i][j], f[1][i][j + 1] + 1);
    }
  }
  for (int i = m - 1; i >= 0; i--) {
    for (int j = 0; j < n; j++) {
      f[2][i][j] = board[i][j] ? 0 : inf;
      if (i < m - 1) f[2][i][j] = min(f[2][i][j], f[2][i + 1][j] + 1);
      if (j > 0) f[2][i][j] = min(f[2][i][j], f[2][i][j - 1] + 1);
    }
  }
  for (int i = m - 1; i >= 0; i--) {
    for (int j = n - 1; j >= 0; j--) {
      f[3][i][j] = board[i][j] ? 0 : inf;
      if (i < m - 1) f[3][i][j] = min(f[3][i][j], f[3][i + 1][j] + 1);
      if (j < n - 1) f[3][i][j] = min(f[3][i][j], f[3][i][j + 1] + 1);
    }
  }
}

bool Check(int dist) {
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      int tmp = min(min(f[0][i][j], f[1][i][j]), min(f[2][i][j], f[3][i][j]));
      need[i][j] = tmp > dist;
    }
  }
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      g[0][i][j] = need[i][j] ? 0 : -inf;
      if (i > 0) g[0][i][j] = max(g[0][i][j], g[0][i - 1][j] + 1);
      if (j > 0) g[0][i][j] = max(g[0][i][j], g[0][i][j - 1] + 1);
    }
  }
  for (int i = 0; i < m; i++) {
    for (int j = n - 1; j >= 0; j--) {
      g[1][i][j] = need[i][j] ? 0 : -inf;
      if (i > 0) g[1][i][j] = max(g[1][i][j], g[1][i - 1][j] + 1);
      if (j < n - 1) g[1][i][j] = max(g[1][i][j], g[1][i][j + 1] + 1);
    }
  }
  for (int i = m - 1; i >= 0; i--) {
    for (int j = 0; j < n; j++) {
      g[2][i][j] = need[i][j] ? 0 : -inf;
      if (i < m - 1) g[2][i][j] = max(g[2][i][j], g[2][i + 1][j] + 1);
      if (j > 0) g[2][i][j] = max(g[2][i][j], g[2][i][j - 1] + 1);
    }
  }
  for (int i = m - 1; i >= 0; i--) {
    for (int j = n - 1; j >= 0; j--) {
      g[3][i][j] = need[i][j] ? 0 : -inf;
      if (i < m - 1) g[3][i][j] = max(g[3][i][j], g[3][i + 1][j] + 1);
      if (j < n - 1) g[3][i][j] = max(g[3][i][j], g[3][i][j + 1] + 1);
    }
  }
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      int tmp = max(max(g[0][i][j], g[1][i][j]), max(g[2][i][j], g[3][i][j]));
      if (tmp <= dist) return true;
    }
  }
  return false;
}

int Work() {
  cin >> m >> n;
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      char ch = getchar();
      while (ch != '0' && ch != '1') ch = getchar();
      board[i][j] = (ch == '1');
    }
  }
  Prepare();
  int l = 0, r = m + n, res = m + n;
  while (l <= r) {
    int mid = (l + r) / 2;
    if (Check(mid)) {
      res = mid;
      r = mid - 1;
    }
    else {
      l = mid + 1;
    }
  }
  return res;
}

int main() {
  int T;
  cin >> T;
  for (int i = 1; i <= T; i++) {
    int ans = Work();
    cout << "Case #" << i << ": " << ans << endl;
  }
  return 0;
}
```

## T3 Contention

这个题是真的不会。看了题解才明白。

### 题目

有N个位置，编号1..N。有Q个订单，每个订单是一个区间[Li, Ri]，其将占据[Li, Ri]间所有没有被占据的位置。你可以任意安排订单来的顺序，求最大的k，使得每一个订单都至少占据k个位置。N<=1000000, Q<=30000。

### 做法

任意安排顺序，那该怎么安排呢？比如说我们首先处理第i个订单，那么后面怎么办？这个订单所占据的位置会影响后面的订单。

这里用到了正难则反的思想，我们到这考虑行不行呢？假设最后处理的是第i个订单，那么它就不影响它前面的订单了，并且无论它前面的订单是按什么顺序排的，订单i能占据的位置是一定的，因为已经被占据的位置就是它前面的区间的并集。

考虑从后往前确定订单的顺序。这里还有一个贪心：每次都选那个放到最后能占据的位置最多的订单。

为什么呢？首先，因为我们要求的是所有订单占据的位置的最小值作为k，所以直观上来说，我们先保证最后的这个尽量大就行。下面用交换法证明：假设这样得到的答案顺序是[..., j, i, ...]。显然j之前和i之后的订单占据多少位置与i和j的顺序没有关系。假设订单i能占据x0个位置，订单j能占据y0个位置。那么这样更新是ans = min(ans, min(x0, y0))。如果我们交换i和j的位置，那么顺序变成[..., i, j, ...]，设订单i能占据x1个位置，订单j能占据y1个位置。这样更新答案就变成了ans = min(ans, min(x1, y1))。根据我们贪心的顺序，一定有x0 >= y1。并且，由于下面一种情况j之前多了一个i，故一定有y0 >= y1。因此，min(x0, y0) >= y1 >= min(x1, y1)。综上，交换相邻的i和j不会使答案更优。对于任意序列a[1..N]，将贪心得到的序列使用类似冒泡排序的方法，通过交换相邻的两个，最终得到序列a。因此，任意一个序列都不会比贪心得到的序列更优。

以上证明了贪心算法的正确性。那么我们从后往前确定订单顺序。每一次循环，我们需要确定哪一个订单放到最后能占据的位置最多，然后将这个订单放到最后并删除，同时更新答案。这里容易想到用数组num[1..N]来维护每一个位置上剩下的还有多少个订单。开始时先通过num[L]++, num[R+1]--，最后求前缀和的方式预处理出来。这样对于订单i，其能占据的位置就是[Li, Ri]之间1的个数。假设我们选定了订单k，那么我们需要将num[Lk, Rk]之间的数减1。如果num[i]被减成1了，我们就需要把占据i的那个区间上1的个数加1。用一个数组cnt维护每一个订单的区间内有多少个1。总结一下，预处理部分我们要做的有：

1. 求出初始num[1..N]数组。
2. 求出初始cnt[1..Q]数组。

每一次循环我们需要做的事情有：

1. 选择cnt最大的订单i。
2. 将num[Li..Ri]减1。
3. 如果num[j]被减成1了，就将j位置对应订单k的cnt加1。
4. 删除订单i。

暴力做，时间复杂度O(NQ)，可以过小数据。

容易想到用线段树维护。区间修改通过懒标记处理。因为我们要找1，所以我们维护区间最小值。如果发现某个位置j被减成1了，那么将其对应的订单cnt[k]加1即可，然后将其改成inf。这样就可以方便地知道是否有1的位置。因为每个位置只会变成1一次，所以这部分的复杂度为O(nlogn)。怎样找到位置j对应的订单k呢？显然可以在线段树的每个节点上用一个set记录覆盖这个节点的订单id的集合，开始时将区间全部插入，每次询问查一下即可。如果用unordered_set的话这部分就可以不增加复杂度。删除订单i，区间删除unordered_set中的值即可。cnt用一个set维护，可以方便地修改并查询最大值。这样总时间复杂度为O(NlogN)。空间上因为我们用了线段树套set，每一个订单会出现在O(logN)个节点上，故空间复杂度为O(NlogN)。

这样交上去会TLE，因为N比较大，并且我们的算法常数也很大。容易发现Q比较小，因此可以离散化，将位置数量变为O(Q)级别，再使用上面的方法时间复杂度为O(QlogQ)。

```C++
#include <cstdio>
#include <iostream>
#include <algorithm>
#include <vector>
#include <set>
#include <unordered_set>
#include <map>

using namespace std;

#define inf 999999999

class Node {
public:
  Node *left, *right;
  int l, r;
  int tag;
  int min_val;
  unordered_set<int> S;

  explicit Node(int l_, int r_) {
    left = right = nullptr;
    l = l_, r = r_;
    tag = 0;
  }

  void Subtract() {
    tag++;
    min_val--;
  }

  void PushDown() {
    if (tag == 0) return;
    if (left != nullptr) {
      left->tag += tag;
      left->min_val -= tag;
    }
    if (right != nullptr) {
      right->tag += tag;
      right->min_val -= tag;
    }
    tag = 0;
  }

  void Update() {
    min_val = min(left->min_val, right->min_val);
  }
};

Node* BuildTree(const vector<int>& a, int l, int r) {
  Node* v = new Node(l, r);
  if (l == r) {
    v->min_val = a[l];
    return v;
  }
  int mid = (l + r) / 2;
  v->left = BuildTree(a, l, mid);
  v->right = BuildTree(a, mid + 1, r);
  v->Update();
  return v;
}

void Insert(Node* v, int ww, int ee, int x) {
  if (ww <= v->l && v->r <= ee) {
    v->S.insert(x);
    return;
  }
  int mid = (v->l + v->r) / 2;
  if (ww <= mid) Insert(v->left, ww, ee, x);
  if (ee > mid) Insert(v->right, ww, ee, x);
}

void Remove(Node* v, int ww, int ee, int x) {
  if (ww <= v->l && v->r <= ee) {
    v->S.erase(x);
    return;
  }
  int mid = (v->l + v->r) / 2;
  if (ww <= mid) Remove(v->left, ww, ee, x);
  if (ee > mid) Remove(v->right, ww, ee, x);
}

int FindId(Node* v, int pos) {
  if (v->l <= pos && pos <= v->r && v->S.size() == 1) {
    return *(v->S.begin());
  }
  int mid = (v->l + v->r) / 2;
  if (pos <= mid) return FindId(v->left, pos);
  else return FindId(v->right, pos);
}

void Subtract(Node* v, int ww, int ee) {
  if (ww <= v->l && v->r <= ee) {
    v->Subtract();
    return;
  }
  v->PushDown();
  int mid = (v->l + v->r) / 2;
  if (ww <= mid) Subtract(v->left, ww, ee);
  if (ee > mid) Subtract(v->right, ww, ee);
  v->Update();
}

int FindMinValuePosition(Node* v) {
  if (v->l == v->r) {
    return v->l;
  }
  v->PushDown();
  if (v->left->min_val == v->min_val) return FindMinValuePosition(v->left);
  else return FindMinValuePosition(v->right);
}

void SetValue(Node* v, int pos, int val) {
  if (v->l == v->r) {
    v->min_val = val;
    return;
  }
  v->PushDown();
  int mid = (v->l + v->r) / 2;
  if (pos <= mid) SetValue(v->left, pos, val);
  else SetValue(v->right, pos, val);
  v->Update();
}

void ClearTrash(Node* v) {
  if (v->left) ClearTrash(v->left);
  if (v->right) ClearTrash(v->right);
  delete v;
}

int Work(int N, int Q, const vector<pair<int, int>>& bookings) {
  set<int> a;
  a.insert(1);
  a.insert(N);
  for (int i = 0; i < Q; i++) {
    if (bookings[i].first > 1) a.insert(bookings[i].first - 1);
    a.insert(bookings[i].first);
    if (bookings[i].first < N) a.insert(bookings[i].first + 1);
    if (bookings[i].second > 1) a.insert(bookings[i].second - 1);
    a.insert(bookings[i].second);
    if (bookings[i].second < N) a.insert(bookings[i].second + 1);
  }
  int M = a.size();
  int last = 0, t = 0;
  vector<int> cover(M + 1);
  map<int, int> go;
  for (int x : a) {
    t++;
    cover[t] = x - last;
    last = x;
    go[x] = t;
  }
  vector<int> num(M + 2);
  vector<pair<int, int>> new_bookings = bookings;
  for (pair<int, int>& booking : new_bookings) {
    booking.first = go[booking.first];
    booking.second = go[booking.second];
    num[booking.first]++;
    num[booking.second + 1]--;
  }
  for (int i = 1; i <= M; i++) {
    num[i] += num[i - 1];
  }
  for (int i = 1; i <= M; i++) {
    if (num[i] == 0) {
      num[i] = inf;
    }
  }
  Node* root = BuildTree(num, 1, M);
  for (int i = 0; i < Q; i++) {
    Insert(root, new_bookings[i].first, new_bookings[i].second, i);
  }
  vector<int> cnt(Q);
  for (int i = 1; i <= M; i++) {
    if (num[i] == 1) {
      int id = FindId(root, i);
      cnt[id] += cover[i];
      SetValue(root, i, inf);
    }
  }
  set<pair<int, int>> cnt_to_id;
  for (int i = 0; i < Q; i++) {
    cnt_to_id.insert(make_pair(cnt[i], i));
  }
  int ans = inf;
  for (int i = 0; i < Q; i++) {
    int id = cnt_to_id.rbegin()->second;
    cnt_to_id.erase(make_pair(cnt[id], id));
    ans = min(ans, cnt[id]);
    Remove(root, new_bookings[id].first, new_bookings[id].second, id);
    Subtract(root, new_bookings[id].first, new_bookings[id].second);
    while (root->min_val == 1) {
      int pos = FindMinValuePosition(root);
      SetValue(root, pos, inf);
      int now_id = FindId(root, pos);
      cnt_to_id.erase(make_pair(cnt[now_id], now_id));
      cnt[now_id] += cover[pos];
      cnt_to_id.insert(make_pair(cnt[now_id], now_id));
    }
  }
  ClearTrash(root);
  return ans;
}

int main() {
  int T;
  cin >> T;
  for (int test_id = 1; test_id <= T; test_id++) {
    int N, Q;
    cin >> N >> Q;
    vector<pair<int, int>> bookings(Q);
    for (int i = 0; i < Q; i++) {
      cin >> bookings[i].first >> bookings[i].second;
    }
    int k = Work(N, Q, bookings);
    cout << "Case #" << test_id << ": " << k << endl;
  }

  return 0;
}
```

## 总结

这三个题一个比一个难，尤其是第三题，考场上有人能做出来也很厉害了，光贪心我就想不出来。还要努力啊。