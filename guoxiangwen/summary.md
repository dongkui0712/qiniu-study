# 2016/1/04
 1.学习Mac上配置go开发环境的技巧
 2.go语言编写test文件
 3.github 上对自己包的管理
 
# 2016/1/05
## go语言类型系统
    字符串
    指针
    数字类型
    切片
    map
    结构体
## 常量定义
    const i int = 1
    const(
      
    )
## 变量定义
    var i int = 10
    i = 20
    j:=10  
    
    var x= interface{}// init nil
    var p *int //指针 为空 nil
    var x int //init 0
    
## 字符串
    var s string = 'hello'
    s[0] = 's' //err ,not modify
    fmt.println()
## 数字类型
    rune 
    var a rune = 72
    fmt.Println(string(a)) //h
## 数组
    var a [5]int
    len(a)
## 切片
    底层是驻足,可变长 实用
    s:=make([]string ,3)
## map
    var m map[string]int
    m["k1"] = 7//error
    v,ok := m["k2"]检查 k2
## make 内建方法
## 指针
    go 指针只是一种应用
    取地址操作符&获取变量存放地址,可以赴给其他指针
## 结构体
    type Person struc{
      Name string
      Age int
      parent string
    }
## 方法定义
    func (p *Person)
        
# 2015/01/06学习总结

## 控制语句和错误处理

goto if else

if else 无小括号
var i = 10
if i == 10{
    //todo
}
i的作用与只限于循环内

switch

switch value{
    case 0://默认break
    case 1,2 //默认只要一个匹配
}
for range 可以迭代 string array 等

错误处理

func ()

# 2015/01/07
## go面向对象的语法
    1.学习了go面向对象的写法
      在go中，采用了一种算是规定的方法：使用大小写来确定，大写是包外可见的，小写的struct或函数只能在包内使用  
      

    type Door struct {
        opened bool
    }
     
    func (d *Door) Open() {
        d.opened = true
    }
     
    func (d *Door) Close() {
        d.opened = false
    }
    2.接口的概念.侵入和非侵入的区别
    any类型是空接口 interface{}
    func UseCoder(c Coder)(err error){
            return
        }
        type Coder interface{
            WriteCode()(err error)
        }
    interface{}
    
    
    