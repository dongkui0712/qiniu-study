# Summary of Golang Study

>  此文档用于纪录golang的学习历程

### 2016-01-04 初识Golang
   ------
   1.搭建golang开发环境
   
   （1）安装http://golangtc.com/download
   
   （2）配置环境变量
        vim ~/.bash_profile
        
        export GOPATH=$HOME/workspace_go
    保存后，source ~/.bash_profile
    
   （3）安装idea的golang开发插件
   
   2.学习golang优缺点，并开始进行第一个go程序开发，及测试用例
   
   3.熟悉整套开发流程，重点熟悉对github的操作
   
    （1）git clone  远程仓库
    （2）将自己新写的加到本地仓库
    （3）git checkout -b 新分支名 
    （4）git remote
    （5）git remote show origin
    （6）git add .（欲提交文件夹)
    （7）git status（查看分支）
    （8）git commit -m "add my code" （添加备注）
    （9）git push origin develop（提交）
    
   ### 2016-01-05  Golang  类型系统
   -----
   **字符串、数字类型、数组、切片、map、指针、结构体**
   
   1. 字符串
   
      *字符串赋值*
	
    	（1）var str string //声明一个字符串变量
    	（2）str = "Hi CHANGHONG"// 字符串赋值
    	（3）ch := str[0] //取字符串第一个
    	
      *字符串操作之遍历*
      
      ``````
      str := "Hello,CHANGHONG"
      n := len(str)
      for i:= 0;i<n;i++{
      ch := str[i]
      fmt.Println(i,ch)
      }
      ``````
      
   2.数字类型
   
   3.数组
   
      ``````
      [32]byte // 长度为32的数组，每个元为一个字符
      [2*n]struct {x,y int32} // 复杂类型数组
      [1000]*float64 //指针数组
      [3][5]int //二维数组
      
      ``````
      
   4.map
   
   （1）结构体
   
      ``````
      type PersonInfo struct{
         ID string
         Name string
         Address string
         }
      var personDB map [string] PersonInfo   //定义一个map
         personDB = make(map[string] personInfo)//生成空map
		//往map插入几条数
		personDB[“12345”] = PsersonInfo{“12345”,”jam”,”beijing”}
		//从map中查找键为“1234”的信息
		person, ok := personDB[“1234”]
		//ok是一个返回的bool型，返回true表示找到对应数
		if ok{
			fmt.Println(“Founf person”,person.Name,”with ID 1234”)
			}else{
			fmt.Println("Did not found person with ID 1234.”)
	``````

   5.结构体
   
      ``````
      type Person struct{
         Id int
         Name string
         Age int
         } 

   p := new(Person)
   p.id=1
   p.Name=“guagua"
   p.Age=24
   ``````

   6.匿名字段
	
      ``````
      type S struct{
      	a string 
	      b string
         }
	
	type B struct{
      		S  //匿名字段，只有类型S，字段名是S
      		b int // 字段名是b
      		int //匿名字段，只有类型int  无意义
         	}

      	var b B
      	b. S.a=“a”//  给匿名字段赋值
      	b.a=“a”  //同上

   	fmt.Ptintln(b)// 输出结果是｛｛a｝｝not ｛a｝

   	b.S.b=“b”给匿名字段S 赋值
   	b.b=10//名字冲突，不同上，不能这样给匿名字段中相同字段赋值
	``````
### Golang 控制语句和错误处理

   **goto：**

  调整执行位置、支持函数内部跳转、跳转到当前函数内定义标签

   **if else**   
   
   *注意：* 条件表达式没有括号

   switch :
   
   支持初始化表达式、默认break、多个分支逗号执行

   **fallthrough ／／进入后面的case进行处理**

   for :
   ``````
   三种形式：
   （1）类似while
   （2） for int;condition;边界｛
	｝
    (3)for range一起使用，完成迭代器的iterator操作 
       string、array、slice、数组都可以使用
       某个值不想要时可以选择丢弃 _ 例如 for _,n := range 数组｛

     （迭代、返回值）两种情况用得多 
    ``````
    **错误处理**

    defer

    error: 抽象了Error()接口方法
    panic


    ####附加 
    
    文件复制、文件修改、限速传输
    
    
