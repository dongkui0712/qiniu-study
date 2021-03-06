2016-01-08

学习go并发编程基础
* 并发关键
  * 调度 充分使用资源
  * 状态同步 锁
“Concurrency is not parallelism
并发不是同时运行，但是能做到同时运行” （时间片）
* Go并发特性
  * goroutine
  Go 可被调度的最小单元 比线程轻量，有自己的栈信息
  * channel
  Go 并发消息同步机制 向channel 里写入和读取消息都是原子操作，不可覆盖。
  当channel不为空，但是被close了，此时channel的状态应该是处于 c不为空 
  只有channel为空并且被close了才算是 c被close
  * select 用法类似switch语句，用做channel的同步控制 没个case必须为一个对channel的操作


--------------------------------------------------------------
2016-01-07

学习go接口和面向对象编程

- 了解go的起源 与c++的对比 
- go很多方面是传统语言进一步的优化 但在面向对象编程确是全新的设计
     * 新特性     
       * 垃圾回收
       * 异常处理 panic/defer/recover
       * 项目管理和编译速度快 模块化
       * 并发编程 goroutin & channel
       * ...
     * 面向对象全新设计   
       * 1. 没有类的概念、只有结构体
            * 结构体=成员变量 + 成员方法
            * 结构体之间无继承、无构造/析构函数
        * 2. 非侵入式接口：
--------------------------------------------------------------

2016-01-06

学习golang 分支语句、循环控制语句、文件操作、异常处理等基础知识
* goto 忽略标签而不是后面的语句

* if

作用域

* switch

支持初始化

默认 break

fallthough

- 错误处理 panic recover

- 限速传输、文件读写、无写权限

--------------------------------------------------------------

2016-01-05

学习Golang 
- 字符串 操作
- 数字类型
- 数组
     - 无符号数: uint8, uint16, uint32, uint64
     - 有符号数: int8, int16, int32, int64
     - 浮点数: float32, float64
     - 复数: complex64, complex128
     - 其它: byte(uint8 的别名), rune(int32 的别名，用来存储 Unicode)
- 切片
     - 底层为数组，但是可变长
     - 比数组更为实用
- map
     - make() 内建方法只能用于创建 slice, map 和 channel 
     - slice, map, channel 是指向底层数据结构的引用，使用前必须被初始化
- 指针
     - Go 有指针，但是没有指针运算。你不能 用指针变量遍历字符串的各个字节。
     - Go 指针只是一种引用。
     - 取址操作符 & 获取变量存放地址，可以赋给一个指针。
- 结构体
     * 创建自己定义的类型
     * 方法定义
     * 匿名字段
     * 方法继承
     * 方法覆盖
     * 给类型 B 添加一个方法覆盖类型 S 中的 String()
     * 字节对齐
---------------------------------------------------

2016-01-04

学习Golang 编程基础:
- Golang 编程单元测试用例学习
- Golang 包引用学习
- github工作流使用、git基础命令学习掌握
- 发布Golang library 到 github
- SublimeText golang 开发环境搭建
- 学习Golang 单元测试库 assert的用法








