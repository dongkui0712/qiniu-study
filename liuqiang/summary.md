#2016-01-04 学习总结

>* Mac OS X系统中配置go环境
>* 将代码托管到github上
>* 利用testing包进行单元测试

##1.在Mac OS X操作系统中安装go环境，并设置GOPATH，GOROOT
  ```
  打开终端（Terminal），输入以下代码：
  cd ~    #进入当前用户下的根目录
  ls -a    # -a可将隐藏文件显示，
  找到.bash_profile文件
  如果没有，就输入：vim .bash_profile   #创建很简单
  进入vim后，按下a才能编辑，输入一下代码：
  export GOROOT=/usr/local/go
  export PATH=/usr/local/go/bin:$PATH
  export GOPATH=你自己平时将go代码放置的地方
  GOROOT就是pkg包默认安装到的地方，从官网上也可以看到，默认是安装到/usr/local/go
  PATH很重要，系统自带的源码要运行必须有这个路径，默认安装路径时，在/usr/local/go/bin
  GOPATH可以是用户任意喜欢的地方，放置自己写的go程序
  此时，.bash_profile文件编写完毕，按esc，敲入:wq，回车，搞定。
  退回到终端后,输入source  .bash_profile 使编辑生效
  验证一下路径配置是否成功：
  go env 就可以显示刚才设置的环境变量
  环境变量配置完成。
  ```
##2.在Mac OS X操作系统中安装Intellij IDEA，并在其中安装Go插件
  ```
  Preference-->Plugins中搜索go插件,然后进行安装,安装完成后重启IDEA就可以了
  ```
##3.在go程序中引入包
  ```
    import "包名"
    如要引入errors包,则输入如下内容:
    import "errors"
  ```
##4.在go中如何定义函数,如何调用另一个包中的函数,什么是go语言中的多返回值
  ```
  func (p myType ) funcName ( a, b int , c string ) ( r , s int ) {
      return
  }
  关键字——func
  方法名——funcName
  入参——— a,b int,b string
  返回值—— r,s int
  函数体—— {}
  ```
##6.使用github托管代码常见操作
  ```
  例如要创建一个develop分支,则输入:
  git checkout -b develop
  例如要在分支中添加一个updateFile文件,则输入:
  git add updateFile文件 
  例如要在在分之内中添加updateInformation更新信息
  git commit -m updateInformation
  执行commit命令后,只是提交到了本地仓库,还未提交到远程仓库
  例如要将其提交到develop分支,则输入:
  git push origin <branchName>
  ```
##7.利用testing包编写单元测试
  ```
  首先需要引入testing包,然后使用如下方式编写测试函数
  func Test_要测试的函数的名字(t *testing.T){
     if 测试不通过{
         t.Error("输出不通过信息")
     }
     t.Log("输出通过信息")
  }
  ```
##8.练习题
  ```
  1.1. 编写第一个 Hello World 程序，将打印功能放在另一个包中实现（涉及 $GOPATH、包管理、函数调用）。
  1.2. 给第一题中的包加上单元测试，并将其发布到 Github 上，本地安装后调用。
  ```
  
#2016-01-05 学习总结

>* 变量定义
>* 指针
>* 结构体
>* rune
>* 切片
>* 映射map

##1.变量定义
  ```
  在go语言中,变量定义跟其他语言不太一样,它的变量类型是写在后面的,如:
  var a int 
  初始化变量时,可以直接给变量赋值,如:
  var a=10  //编译器会自动推导a的类型
  也可以不声明,直接使用 := 进行赋值,如:
  a :=3.14
  变量类型既可以是内置类型,也可以是接口类型,如:
  var x interface{}
  ```
##2.指针
  ```
  定义指针
  var p *int即可定义一个指针
  ```
##3.结构体
  ```
  type Person struc{
           Name string
           Age  int
           parent string
      }
  p :=new(Person)
  p.Name="aa"
  p.Age=23
  p.parent="AA"
  ```
##4.什么是rune及其作用
  ```
     int32的别名,可以用来存储unicode的值
  ```
##5.匿名字段
  ```
   当匿名字段是一个struct的时候，那么这个struct
   所拥有的全部字段都被隐式地引入了当前定义的这个struct
   type S struct{
          a string
          b string
      }
   type B struct{
          S    //匿名字段,只有类型S,字段名是S
          b int //字段名是b
          int//匿名字段,只有类型int,无意义
      }
      var b B
      b.S.a="a"//给匿名字段s赋值
      b.a="a"  //同上
      fmt.Println(b)//输出结果是{{a}} not {a}
      b.S.b="b"给匿名字段S赋值
      b.b=10//名字冲突,不同上,不能这样给匿名字段中相同字段赋值
  ```
##6.切片及其使用方法
  ```
  Slice可以通过两种方式定义，一种是从源数组中衍生，一种是通过make函数定义，本质上来说都一样，
  都是在内存中通过数组的初始化的方式开辟一块内存，将其划分为若干个小块用来存储数组元素，然后
  Slice就去引用整个或者局部数组元素。
  *从数组中切片构建Slice：
  a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
  s := a[2:8]
  fmt.Println(s)	//输出：[3 4 5 6 7 8]
  定义一个数组a，截取下标为2到8之间部分（包括2不包括8），构建一个Slice。
  *通过make函数定义：
  s := make([]int, 10, 20)
  fmt.Println(s) //输出：[0 0 0 0 0 0 0 0 0 0]
  make函数第一个参数表示构建的数组的类型，第二个参数为数组的长度，第三个参数可选，是slice的容量，默认为第二个参数值。 
  ```

##7.map及其使用方法
  ```
  map 是一种无序的键值对的集合。map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
  map 是一种集合，所以我们可以像迭代数组和 slice 那样迭代它。不过，map 是无序的，我们无法决定它的返回顺序，这是因为 
  map 是使用 hash 表来实现的。
  Go 语言中有多种方法创建和初始化 map。我们可以使用内建函数 make 也可以使用 map 字面值：
  其格式为:
  map,m := map[KeyType]ValueType
  // 通过 make 来创建
  dict := make(map[string]int)
  // 通过字面值创建
  dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
  ```
##8.练习题:
  ```
  2.1. 建立一个程序统计字符串里的字符数量，同时输出这个字符串的字节数。
  2.2. 实现一个对字符串进行反转的程序。
  2.3. 编写一个计算类型是 float64 的 slice 的平均值的代码。
  2.4. 实现一个程序，保存多个用户输入的姓名、年龄和学号，并在输入结束后将所有用户信息按学号排序后打印出来。
  2.5. 设计一个字符串匹配算法，找出字符串 s1 在 s2 中出现的所有位置，不存在的情况返回 -1。
  ```

#2016-01-06 学习总结

>  今天主要的学习内容是控制语句和错误处理以及对昨天所做题目进行了交流,

>* if-else语句
>* switch case语句
>* for循环
>* defer关键字
>* panic,recover关键字

##1.if-else语句
  ```
  Go的if有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，
  其他地方就不起作用了，如下所示：
  // 计算获取值x,然后根据x返回的大小，判断是否大于10。
  if x := computedValue(); x > 10 {
      fmt.Println("x is greater than 10")
  } else {
      fmt.Println("x is less than 10")
  }
  //这个地方如果这样调用就编译出错了，因为x是条件里面的变量
  fmt.Println(x)
  多个条件的时候如下所示：  
  if integer == 3 {
      fmt.Println("The integer is equal to 3")
  } else if integer < 3 {
      fmt.Println("The integer is less than 3")
  } else {
      fmt.Println("The integer is greater than 3")
  }
  ```
##2.switch case语句
  ```
  它的语法如下
  switch sExpr {
  case expr1:
      some instructions
  case expr2:
      some other instructions
  case expr3:
      some other instructions
  default:
      other code
  }
  sExpr和expr1、expr2、expr3的类型必须一致。Go的switch非常灵活，表达式不必是常量或整数，
  执行的过程从上至下，直到找到匹配项；而如果switch没有表达式，它会匹配true。
  i := 10
  switch i {
  case 1:
      fmt.Println("i is equal to 1")
  case 2, 3, 4:
      fmt.Println("i is equal to 2, 3 or 4")
  case 10:
      fmt.Println("i is equal to 10")
  default:
      fmt.Println("All I know is that i is an integer")
  }
  在第5行中，我们把很多值聚合在了一个case里面，同时，Go里面switch默认相当于每个case最后带有break，
  匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
  ```
##3.for循环
  ```
  在for循环里面有两个关键操作break和continue ,break操作是跳出当前循环，continue是跳过本次循环。
  当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置，例如：
  for index := 10; index>0; index-- {
      if index == 5{
          break // 或者continue
      }
      fmt.Println(index)
  }
  // break打印出来10、9、8、7、6
  // continue打印出来10、9、8、7、6、4、3、2、1
  break和continue还可以跟着标号，用来跳到多重循环中的外层循环
  for配合range可以用于读取slice和map的数据：
  for k,v:=range map {
      fmt.Println("map's key:",k)
      fmt.Println("map's val:",v)
  }
  ```
##4.defer关键字
  ```
  Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。
  当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。
  特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，
  不然很容易造成资源泄露等问题
  ```
##5.panic,recover关键字
  ```
  Go使用了panic和recover机制来处理异常。
  *Panic
  是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。当函数F调用panic，函数F的执行被中断，
  但是F中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，F的行为就像调用了panic。
  这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。
  恐慌可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。
  *Recover
  是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。recover仅在延迟函数中有效。
  在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。如果当前的goroutine陷入恐慌，
  调用recover可以捕获到panic的输入值，并且恢复正常的执行。
  ```
##6.练习题
  ```
  3.1. 求最大素数中 if 语句总执行次数
  输入：一个 3 - 2^32 的整数 N
  输出：小于 N 的最大素数，以及程序中所有if语句（含if、switch、for循环条件）的总执行次数。
  要求：
  用 Go 语言；
  不允许调用任何标准库或者第三方库，但允许用 builtin 函数，比如 make、println。
  3.2. 实现一个程序，限速拷贝文件
  输入：速度
  输出：拷贝时间
  3.3. 实现一个程序，验证 defer 语句和 return 语句的执行顺序
  ```
#2016-01-07 学习总结

>  今天主要的学习内容是接口和面向对象编程

>* go的起源及与其他语言的比较
>* go面向对象机制
>* 接口
>* 接口查询

##1.go的起源及与其他语言的比较
  ```
  在很多方面，go提供的编程机制，是传统语言的进一步优化,但是在面向对象机制上，却是全新的设计
  与其他语言相比较,go的语法更简洁.如:
  a.去掉了分号
  b.函数可以返回多个返回值
  c.defer机制
  具备垃圾自动回收机制
  异常处理和控制逻辑相分离,使得异常处理不干扰正常逻辑
  ```
##2.go的面向对象机制
  ```
  在go语言中,
  1.没有类的概念，只有结构体的概念。结构体=成员变量+成员方法。
    结构体之间没有继承，没有构造/析构函数。
  2.非侵入式接口：一个结构体只要实现了某个接口的所有函数，就认为实现了该接口
  3.语法上，go的结构体对应java的类,有成员变量，有成员方法
  4.go的结构体没有构造函数，提供了一种简洁的结构体初始化方式。没有析构函数，结构体对象由go的垃圾回收器自动回收
  5.go的结构体可以匿名组合。组合原则：组合了一个结构体，就拥有该结构体的变量和方法
  6.如果结构体中的变量，首字母大写，那么这个变量能够被外面的pkg访问到，否则，只能在本pkg内访问
  ```
##3.接口:
  ```
  接口的概念：
  和java的接口一样，是一组方法的集合。接口是一种规范，描述了要做什么。
  而实现了该接口的结构体，则描述了怎么做。
  侵入式和非侵入式的区别：
  java的接口是侵入式的，如果一个类要实现该接口，需要显式声明。
  而go语言的接口是非侵入式的，在定义go的结构体时，不需要显式声明要实现哪个接口。只要实现了该接口所有的方法，
  就实现了该接口。
  ```
##4.接口查询
  ```
  1.查询某个接口对象，是否是某个结构体对象
  2.查询某个接口对象指向的结构体对象，是否实现了另外一个接口
  ```
##5.练习题
####虚拟一家软件公司的运作流程流程
  ```
  老Y开了一家软件开发公司，公司的运作规则很简单：
  1.总投资100万
  2.老Y自己跑业务，获取一个订单，订单有价格（如果失败就休息一天然后继续）
  3.通过人才市场，雇佣一个程序员编写代码（如果没找到，就退回订单并赔偿2倍）
  4.该程序员编写好代码（如果开发失败，则公司倒闭，创业失败）
  5.交付订单，获得收入
  6.付给程序员money，然后结束雇佣（money=订单价格/10）
  7.获取下一个订单
  一直到100万投资耗尽，或者公司资产超过1000万，结束创业。
  ```
