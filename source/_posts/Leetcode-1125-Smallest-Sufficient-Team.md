---
title: Leetcode 1125. Smallest Sufficient Team
toc: true
date: 2019-08-16 20:25:32
categories:
- Algorithm
tags:
- leetcode
---
## 题目
https://leetcode.com/problems/smallest-sufficient-team/

有n个数，m个数集，要从中选出k个数集，使得这k个集合的并集包含这n个数。求最小的k。

n <= 16, m <= 60
## 做法
状态压缩动态规划。设f[S]表示包含集合S中的数的最少集合数。枚举S的每一个子集A，那么f[S] = min(f[A] + f[S - A])。通过给定的m个集合初始化即可。代码如下：
```C++
class Solution {
public:
    
    int n;
    
    void getans(const vector<int>& from, const vector<int>& go, int S, vector<int>* ans) {
        if (S == 0) return;
        if (go[S] != -1) {
            ans->push_back(go[S]);
            return;
        }
        int subset = from[S];
        if (subset == 0 || subset == S) return;
        getans(from, go, subset, ans);
        getans(from, go, S ^ subset, ans);
    }
    
    vector<int> smallestSufficientTeam(vector<string>& req_skills, vector<vector<string>>& people) {
        n = req_skills.size();
        vector<int> f(1 << n);
        vector<int> go(1 << n);
        f[0] = 0;
        for (int i = 1; i < f.size(); i++) {
            f[i] = 1e8;
        }
        for (int i = 0; i < go.size(); i++) {
            go[i] = -1;
        }
        for (int i = 0; i < people.size(); i++) {
            int S = 0;
            for (const string& skill : people[i]) {
                for (int k = 0; k < req_skills.size(); k++) {
                    if (req_skills[k] == skill) {
                        S |= 1 << k;
                        break;
                    }
                }
            }
            // 枚举S的子集
            for (int subset = S; subset > 0; subset = (subset - 1) & S) {
                f[subset] = 1;
                go[subset] = i;
            }
        }
        vector<int> from(1 << n);
        for (int S = 0; S < f.size(); S++) {
            for (int subset = S; subset > 0; subset = (subset - 1) & S) {
                int num = f[subset] + f[S ^ subset];
                if (num < f[S]) {
                    f[S] = num;
                    from[S] = subset;
                }
            }
        }
        vector<int> ans;
        getans(from, go, (1 << n) - 1, &ans);
        sort(ans.begin(), ans.end());
        return ans;
    }
};
```
下面分析时间复杂度。显然我们枚举了n个数的所有子集的所有子集。因此时间复杂度为：
$$C_n^0\times2^0+C_n^1\times2^1+C_n^2\times2^2+...+C_n^n\times2^n$$
高中时学过杨辉三角，可以知道上面这个式子就是$(2+1)^n$的展开式。因此总时间复杂度O($3^n$)，$3^{16}=43046721$，四千多万，可以通过。
