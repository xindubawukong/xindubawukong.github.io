---
title: Leetcode 1172. Dinner Plate Stacks
date: 2019-08-27 19:54:13
categories:
- Algorithm
tags:
- leetcode
---

## 题目

https://leetcode.com/problems/dinner-plate-stacks/

## 做法

用一个stack的数组维护所有栈，一个set维护所有非空的栈，一个set维护所有没有满的栈即可。每次操作时间复杂度O(logn)。

```C++
#include <iostream>
#include <set>
#include <stack>
#include <vector>

using namespace std;

const int kMaxIndex = 200011;

class DinnerPlates {
public:
  int capacity;
  stack<int> stacks[kMaxIndex];
  set<int> not_empty_stacks;
  set<int> not_full_stacks;
  int now_max;

  DinnerPlates(int capacity) {
    this->capacity = capacity;
    now_max = -1;
  }

  void push(int val) {
    if (not_full_stacks.empty()) {
      now_max++;
      not_full_stacks.insert(now_max);
    }
    int id = *not_full_stacks.begin();
    if (stacks[id].empty()) {
      not_empty_stacks.insert(id);
    }
    stacks[id].push(val);
    if (stacks[id].size() == capacity) {
      not_full_stacks.erase(id);
    }
  }

  int pop() {
    if (not_empty_stacks.empty()) return -1;
    int id = *not_empty_stacks.rbegin();
    int val = stacks[id].top();
    if (stacks[id].size() == capacity) {
      not_full_stacks.insert(id);
    }
    stacks[id].pop();
    if (stacks[id].empty()) {
      not_empty_stacks.erase(id);
    }
    return val;
  }

  int popAtStack(int index) {
    if (stacks[index].empty()) return -1;
    int val = stacks[index].top();
    if (stacks[index].size() == capacity) {
      not_full_stacks.insert(index);
    }
    stacks[index].pop();
    if (stacks[index].empty()) {
      not_empty_stacks.erase(index);
    }
    return val;
  }
};
```

不过不知道为啥跑得好慢，交了几次过了，faster than 5%。。