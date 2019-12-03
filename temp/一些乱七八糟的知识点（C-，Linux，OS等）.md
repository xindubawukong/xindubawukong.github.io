---
title: 一些乱七八糟的知识点（C++，Linux，OS等）
date: 2019-10-22 01:46:31
---

### 静态库和动态库的区别

简单来说，静态库会在编译出的可执行文件中包含库代码的一份完整拷贝。动态库是程序运行的时候由系统动态的加载到内存中供程序调用。

参考：

- https://blog.csdn.net/jnu_simba/article/details/9569107
- https://blog.csdn.net/qq_34199383/article/details/80308782

### 进程和线程的区别

进程是具有一定功能的程序关于某个数据集合上的一次运行活动，进程是系统进行资源调度和分配的一个独立单位。

线程是进程的实体，是CPU调度和分派的基本单位，它是比进程更小的能独立运行的基本单位。

一个进程可以有多个线程，多个线程也可以并发执行。一个进程中的不同线程共享同一块内存空间。

### 常量指针和指针常量

`const int *p`或`int const *p`：不能通过指针引用修改内存中的数据。

`int* const p`：本身是一个常量，但是指针所指向的内容可以改变。

参考：
- https://blog.csdn.net/jackystudio/article/details/11519817

### 网络的环形拓扑结构

注意数据是单向流动的。故一个节点出问题，网络就会崩溃。

参考：

- https://zh.wikipedia.org/wiki/%E7%92%B0%E7%8B%80%E6%8B%93%E6%92%B2

### C++ iomanip中的流控制函数

参考：

- https://www.cnblogs.com/byteHuang/p/9968167.html

### 不同网络层次中的数据传输单位

- 物理层：比特（Bit）
- 数据链路层：数据帧（Frame）
- 网络层：分组数据包（Packet）
- 传输层：数据段（Segment）或报文

### C++ `std::thread`, `std::mutex`等相关知识

`#include <thread>`：包括`thread`类及相关函数。

`#include <mutex>`：包括`std::mutex`, `std::lock_guard`, `std::unique_lock`等类，用来保证线程同步的，防止不同的线程同时操作同一个共享数据。

参考：
- https://www.runoob.com/w3cnote/cpp-std-thread.html
- https://blog.csdn.net/weixin_40087851/article/details/82685510
