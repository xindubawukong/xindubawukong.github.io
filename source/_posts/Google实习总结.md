---
title: Google实习总结
date: 2019-08-18 02:35:07
categories:
- Summary
tags:
- google
---

知乎回答：https://www.zhihu.com/question/336859540/answer/761912465

我的Google账号：xiangyunding@google.com，但check out之后就不能用了。

感谢以下个人/组织：

- host: wenjies，co-host: chaoyuel
- recruiter：irene, rosy, yun, etc.
- 帮我内推的sunyq和帮我介绍内推的jzh
- 帮我review代码，还有其他帮助我的Googler
- 一起入职的小伙伴们
- Dataz group
- Google

## 申请过程

关于正式的招聘流程和最新的招聘信息，请关注：

- 微信公众号：Google招聘包打听
- 知乎：<a href="https://www.zhihu.com/org/googlegu-ge-zhong-guo-xiao-yuan-zhao-pin/activities">Google谷歌中国校园招聘</a>
- 其他官方消息渠道

我大概的流程是这样的：
- 2018.12.23 在群里看到转发的招实习生的通知，就发了简历申请。
- 2018.12.26 收到官方邮件，填写申请表格，主要是：
    - 基本信息，简历
    - 期望的面试时间，我选的是寒假之后的。
- 2019.1.4 通过recruiter review
- 2019.2.26 第一轮电话面试。两道算法题，都做对了。
- 2019.3.7 第二轮电话面试。还是两道算法题，也都做对了。
- 2019.3.13 收到offer。与hr商量实习时间，还考虑到期末可能事情比较多，最后确定了5.6-8.30。因为Google的实习生要求实习天数60-70天（不能少也不能超）。
- 2019.5.6 入职

说实话，最开始只是随便投了一下，都没想到真的能过。几点经验总结：
- 一定要事先调试好设备和网络（需要用google docs）。
- 多与面试官交流。Google不喜欢闷着头做题的人。

Google的申请流程很正规，一步一步很清楚，不过流程也比较漫长。每周recruiter都会发邮件sync进度，有问题也可以发邮件咨询。Google对实习生招聘要求的很严格，每位FTE想招实习生，都要写一份详细的报告，实习生总数也控制的很严。今年实习生与往年相比招了很多，北京上海两个办公室一共140人左右（包括RS，SWE，ML SWE和EP）。另外Google招聘对女生很照顾，即使不算EP，感觉男女数量也差不多（假设男生女生水平一样，招聘的男女比应该与高校计算机专业的男女比一样才对，看看贵系男女比我就不好说什么了）。学校和专业虽然说不考虑，但我认识的intern还是清北居多，并且硕博比较多。

有想申请的同学请关注的官方消息。我虽然离职了应该也可以帮忙推简历的。至于实习待遇方面，可以说工资、工作环境没有其他公司能跟Google比了。具体薪资应该不让说，但可以给个参考价，比头条多不少，而且包三餐，而且发好多房补。很多人跟我一样可以住学校，房子根本没花钱，但Google就是这么大方。

## 实习

实习项目是WikiTrust Revamp，维基百科的恶意修改检测。因为签了保密协议，这部分就不详细讲了。只是做了一点微小的工作。据说这个project的由来是，去年特朗普发现维基百科经常被恶搞，就在twitter上@了一下Google。WikiTrust很久之前有人搞过，不过效果非常差，于是Google就安排了几个人做新的WikiTrust，他们就取了个名字叫WikiTrust Revamp，修补的意思。当然整个Dataz组做的事情就很多了。

Google作为世界顶级的软件公司，内部工具和文档、开发流程、代码质量控制都做的太成熟了，之前上软工的时候还觉得这些东西好麻烦，为什么要写单元测试，为什么要搞文档和注释，来到Google才发现这些东西是真的游泳。在Google提交每一份代码，首先都要通过g4 fix自动调整一下风格，比如每行不超过80字符，include顺序等等。然后至少两人review，包括代码逻辑、代码风格、可读性、可维护性检查，觉得不好的地方可以写comments。然后，针对每一个comments修改代码，重复上面的流程。在这之间系统会运行presubmit（包括很多静态检查、运行所有相关的单元测试等）。直到所有的reviewer都认为可以提交了，并且presubmit也通过了，这时才可以点击submit按钮，将这份代码提交到代码库里。我交的最慢的一个cl交了两个多星期，是临时找了一位美国那边的精通C++的同事review，给我提了五十多条comments，改的时候整个页面comments比code还多。注释和单元测试就更不必说了，没有这些不可能让你提交的。Google内部的工具和文档真的太全面了，比如moma、codesearch、critique、piper、cider、flume等等，而且都有十分详细的文档或codelab帮助使用。除了这些开发用的工具之外，其他用的也几乎都是Google自家的产品：查资料用Google search，邮件用gmail，浏览器用Chrome，文件用Google docs，表格用Google sheets，ppt用Google slides，开会用meetings，英文看不懂的用Google translate，有问题上moma（内部的search），发消息用hangout chat，等等。Google的基础设施真的让人很震撼。

工作环境方面，真的是太太太太太好了。一日三餐，零食水果尽情享受，还有健身房、按摩椅、乒乓球台、游戏机……还有好多我没有尝试过的。据我观察，正式员工大多9-10点到公司，下午5-6点就走，晚饭后公司就基本没人了。正式员工都是采用<a href="https://en.wikipedia.org/wiki/OKR">OKR工作法</a>，同事之间、上下级之间交流也非常轻松愉悦。由于Google是美国企业，入职时还给我们宣传Google的价值观。印象最深的是比较注重个人隐私，然后开放自由之类的。

Intern offsite去的是这个地方：http://www.dianping.com/shop/97821752 。我个人不是很喜欢这种室内娱乐场所，不过这是大家投票决定的。午饭每人100元budget随意吃。晚餐去的这里：http://www.dianping.com/shop/69719009 ，人均270的自助。Offsite这天工资还照发，舒服。

## 总结

其实开始找内推的时候真的没想到能来。并且后面找工作的时候才发现，Google的实习经历真的管用，简历几乎没有被卡过（我清华同学本科无实习的简历有时就会被卡）。

TBD