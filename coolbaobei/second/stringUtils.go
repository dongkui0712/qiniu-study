package second

import ("fmt"
	"unsafe"
)
//统计字符数/字节数
func Charachars(s string) (a int,b int, err error) {
	fmt.Println("统计字符串字符数")
//	str := "Hello world"
	n := len(s)
	fmt.Println("字符长度:",n)
	byteNum := unsafe.Sizeof(s)

	return n,int(byteNum),err
}
//字符串反转
func ReverseByArray(str string)(str1 string,err error){

	newstr := ""
	for _,c := range str {
		newstr = string(c)+ newstr
	}
	return newstr,err

}

//编写一个计算类型是 float64 的 slice 的平均值的代码
func Average(a []float64,n int)(b float64,err error){
	for i, v := range a {
		fmt.Println("mySlice[", i, "] =", v)
		b = b+v;
//		fmt.Println(b)
	}
	b= b/float64(n);
//	fmt.Println(b)
	return b,err
}

//实现一个程序，保存多个用户输入的姓名、年龄和学号，并在输入结束后将所有用户信息按学号排序后打印出来。
type Person struct{
	ID int
	Name string
	Age string
}
func SaveInfo(ID int, name string ,age string) (a map[int] Person,err error){
	var personDB map[int] Person
	personDB = make(map[int] Person)
	personDB[ID] = Person{ID, name, age}
	return personDB,err
}

//func StringIndex(str1 string,str2 string)error{
//	param1 := []rune(str1)
//	param2 := []rune(str2)
//	position :=0
//
//	for i:=0;i<len(param1);i++{
//		if param1[i] == param2[0]{
//			for j:=0;j<len(param2);j++{
//
//			}
//		}
//	}
//	return
//}



