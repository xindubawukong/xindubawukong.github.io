---
title: Leetcode 1074. Number of Submatrices That Sum to Target
date: 2019-08-19 02:06:03
categories:
- Algorithm
tags:
- leetcode
---

## 题目

https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/

给定一个m * n的矩阵，求有多少个子矩阵，满足其中元素的和等于target。m, n <= 300。

## 做法

经典题。首先区间求和一定要用前缀和。枚举子矩阵的左右边界[l, r]，那么我们只需要处理[l, r]中间的这一块。这部分的前缀和可以预处理出来。然后用一个hash map统计即可。如果hash map是O(1)的，那么算法时间复杂度为O($n^3$)。

```C++
class Solution {
public:
    int Count(const vector<int>& sum, int target) {
        unordered_map<int, int> record;
        record[0] = 1;
        int res = 0;
        for (int i = 0; i < sum.size(); i++) {
            res += record[sum[i] - target];
            record[sum[i]]++;
        }
        return res;
    }
    
    int numSubmatrixSumTarget(vector<vector<int>>& matrix, int target) {
        int m = matrix.size();
        int n = matrix[0].size();
        vector<vector<int>> row_sum(m);
        for (int i = 0; i < m; i++) {
            row_sum[i].resize(n);
            row_sum[i][0] = matrix[i][0];
            for (int j = 1; j < n; j++) {
                row_sum[i][j] = row_sum[i][j - 1] + matrix[i][j];
            }
        }
        int ans = 0;
        for (int l = 0; l < n; l++) {
            for (int r = l; r < n; r++) {
                vector<int> sum(m);
                sum[0] = row_sum[0][r] - (l == 0 ? 0 : row_sum[0][l - 1]);
                for (int i = 1; i < m; i++) {
                    sum[i] = sum[i - 1] + row_sum[i][r] - (l == 0 ? 0 : row_sum[i][l - 1]);
                }
                ans += Count(sum, target);
            }
        }
        return ans;
    }
};
```
