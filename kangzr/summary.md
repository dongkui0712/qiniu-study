------
康钟荣
2016-01-04 学习总结:
     今天是在七牛培训的第一天,主要内容是熟悉使用mac,配置开发环境与工具,Go语言入门基础.
     上午是熟悉Mac,并安装开发环境,包括Go SDK,IDEA集成开发工具及Go语言插件的安装与调试,以及postman等其他开发与测试辅助工具;
     下午是Go语言入门培训,包括类型,函数,包导入等基础技能,讲师讲的时间较短,主要时间是题目练习,学员展示,及GitHub工具的命令行使用方法,并将
  所做的作业提交到GitHub上,学会使用了代码在Git命令行下的提交上传等.练习题目比较开放,主要要求是完成函数功能,并要设计开发完整的单元测试,
  初步掌握函数式开发及Go语言内置的单元测试组件.
  
2016-01-05 学习总结:
     今日培训主要内容是Go语言类型系统,以及相关的5道编程练习题.
     培训内容,类型系统,包括:字符串,数字类型,数组,数组切片,映射map,指针,结构体等,并阐述了相应的使用方法及示例.
     5道练习题的主要内容是类型系统的操作技巧,基本使用方法,以及字符串匹配算法等,对于个人来说,今天的练习题对学习Go的益处较为明显.
     
2016-01-06 学习总结:
     今天培训的主要内容是Go语言的流程控制语句与错误处理,包括了流程控制关键字goto,if else,switch,for等,其中,for循环有三种形式;错误处理
  关键字包括log,panic,defer,error,其中defer类似于栈,有先进后出的特点;附加讲解内容有文件复制,文件修改,限速传输等文件操作技巧.
     本日练习题1,是求小于输入整数N的最大素数,并统计if的总执行数;练习题2是,实现限速拷贝的程序.
     
2016-01-07 学习总结:
     今天的培训学习内容是,Go语言中的借口及面向对象编程,go时一种面向对象的全新设计,C++,java面向对象的两大机制是,强大的类,以及侵入式地接口机制,
  go面向对象的语法是,结构体,成员变量,成员方法,接口,Go语言提供一种any类型,即interface{},任何类型的变量,都可以赋值给interface{},并提供了一个
  接口查询的方法,用来判断结构体是否实现了该接口中的方法,Go语言的编程哲学是实用派,非学院派.
     今日练习题,是一个关于老A创业,获取订单并招聘程序员完成订单的过程,反复循环,直到满足预先设定的条件,退出循环.
     
2016-01-08 学习总结:
     今天培训学习的主题是Go语言并发基础.并发的目的是在有限的服务器资源下,使得CPU能够完成更多地任务,单核也需要并发.
     并发的特点是调度与锁,其中调度的作用是为了充分使用资源;锁即是状态同步,任务需要按照顺序来执行,类似于先闭眼再睡觉,并发不是同时运行,
     但是能够做到同时运行.
     Go的并发机制中有三个关键字,gorountine,channel,select.其中goruntine是go中调度的一个最小单元,比线程轻量;channel通过消息同步实现
     内存共享,往channle里写消息与读消息是一个原子操作;select配合channel选择执行的路径,用来做channel的同步控制.
     本日练习题是并发相关的编程题.
     
2016-01-11 学习总结:
     今天培训学习的内容主要是,Go语言WEB表单处理,以及上周练习题的review.
     培训内容,表单处理部分,为我们讲述Go Web开发的基础知识与技能,并展示了一些单间示例.
     练习题的review部门,分别请来了上周为我们培训流程控制错误处理,与Go面向对象特性的两位讲师,为我们详细解析了这两天培训练习题,同时为我们分别指出了
     我们练习题答案中的各种不足,及待优化的地方.
     
     
     
     
     
     
     
     
     
     
     
     
     
     
     
     
     
     