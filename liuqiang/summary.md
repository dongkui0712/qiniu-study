##2016-01-04 学习总结

1.学习了如何在Mac OS操作系统中安装go环境，并设置GOPATH，GOROOT，GOBIN

2.学习了如何在Mac OS操作系统中安装Intellij IDEA，并在其中安装Go插件

3.学习了如何在go程序中引入go常用的包,如引入errors包

4.学习了如何写go函数并进行调用

5.学习了什么是go函数的多返回值

6.学习了如何在git中创建分支,使用命令:
  git checkout -b branchName

7.学习了如何将文件添加和提交到分支,使用命令:
  git add fileName,git commit -m addContent

8.学习了如何提交到远程分支,使用命令:
  git push origin branchName

9.学习了如何克隆项目到本地，如:
  git clone https://github.com/IT20193476/qiniu-study.git
  
  
##2016-01-05 学习总结

1.学习了如何定义一个指针,如var p *int,以及如何使用指针

2.学习了什么是结构体,以及如何访问结构体中的字段

3.学习什么是rune,知道了它是int32的别名,可以用来存储unicode的值

4.学习了什么是匿名字段,当匿名字段是一个struct的时候，那么这个struct
  所拥有的全部字段都被隐式地引入了当前定义的这个struct

5.学习了什么是切片及它的使用方法

6.学习了什么是映射,如何使用map,
  m := map[KeyType]ValueType

7.通过以上的学习,会利用以上知识完成求字符串的个数,完成字符串的反转,
  完成在一串字符串中查找字串并返回其起始位置,完成求一个切片数组中的数据的平均值

##2016-01-06 学习总结

>今天主要的学习内容是控制语句和错误处理以及对昨天所做题目进行了交流,主要学习内容如下:

1.学习了如何使用if else语句,其中需要注意的是条件表达式没有括号
  
2.学习了如何使用switch case语句,在golang中的switch case语句中,默认是加了break的,如果不进行break,可以添加fallthrough关键字
  
3.学习了如何使用for循环,可以通过range返回序号和值,在返回值的时候,如果不想要某个返回值,可以用_
  
4.学习了如何使用defer关键字,defer在声明时不会立即执行，而是在函数 return 后，再按照 FILO（先进后出）的原则依次执行每一个defer，一般用于异常处理、释放资源、清理数据、记录日志等。这有点像面向对象语言的析构函数，优雅又简洁，是 Golang 的亮点之一
  
5.学习了如何使用panic,recover关键字来进行错误处理
