package main

import (
	"firstGo/second"
	"fmt"
	"sort"
	"strings"
//	"firstGo/third"
//	"os"
//	"time"
//	"os"
//	"io"
//	"time"
//	"fmt"
)

func main() {

	c,b,err := second.Charachars("Hello world")
	if(err != nil){
		fmt.Println("出错了")
	}else{
		fmt.Println("输出字符数是:",c)
		fmt.Println("输出字节数是:",b)
	}

	a,err := second.ReverseByArray("Hello world")
	if(err != nil){
		fmt.Println("出错了")
	}else{
		fmt.Println("hello world 反转结果是:",a)
	}

	var myArray [6]float64 = [6]float64{1.3, 2.1, 3.2, 4.5, 5.21,1.1}
	var mySlice []float64 = myArray[:5]
	result,err := second.Average(mySlice,len(mySlice))
	if(err != nil){
		fmt.Println("出错了")
	}else{
		fmt.Println("slice的平均值是:",result)
	}

	type Person struct{
		ID int
		Name string
		Age string
	}
//	var name string
//	var age string
//
//	fmt.Println("Please input your  name: ")
//	fmt.Scanf(&name)
//	fmt.Printf("Hi %s!\n", name)
//	fmt.Println("Please input your  age: ")
//	fmt.Scanf(&age)
//	fmt.Printf("Wow you %s!\n", age)
//	personDB,err := second.SaveInfo(rand.Intn(20),name,name)


	var personDB map[int] Person
	personDB = make(map[int] Person)
	personDB[3] = Person{3, "Tom", "12"}
	personDB[1] = Person{1, "Jack", "15"}
	personDB[6] = Person{6, "Jim", "15"}

	var ids []int
	for _,n := range personDB{
		ids = append(ids,n.ID)

	}
	if(ids == nil){
		fmt.Println("出错了")
	}
	sort.Ints(ids)

	for _,m := range ids{
		person, ok := personDB[m]
		if ok {
			personDB[m] = Person{m, person.Name, person.Age}
			fmt.Println(personDB[m])
		}
	}
//	fmt.Println("排序后的结果是:")
	str1 := "Hi qiniu Hi qiniu Hi qiniu Hi qiniu Hi qiniu "
	str2 := "qiniu"
	first := strings.Index(str1,str2)
	last :=strings.LastIndex(str1,str2)
	count := strings.Count(str1,str2)
	res := strings.IndexAny(str1,str2)
	fmt.Println("the first:",first)
	fmt.Println("the last :",last)
	fmt.Println("count of index:",count)
	fmt.Println(res)
//
////	err1 := third.JudgPrime(2)
////	if err1 != nil{
////		fmt.Println("程序错误")
////	}
//	var nCount int
//	n := 1000
//	for i := 1; i <= n; i++ {
//		if third.JudgPrime(i) {
//			nCount += 1
//			fmt.Printf("%5d", i)
//			if nCount%8 == 0 {
//				fmt.Println("")
//			}
//		}
//	}
//	fmt.Println("Count =", nCount)
//
//	userFile := "text.txt"
//	fout,err := os.Create(userFile)
//	defer fout.Close()
//	if err != nil{
//		fmt.Println(userFile,err)
//		return
//	}
//	for i := 0;i<10;i++{
//		fout.WriteString("TEST!\n\t")
//		fout.Write([]byte("TEST!\n\t"))
//	}
//
//	fin,err := os.Open(userFile)
//	defer fin.Close()
//	if err != nil{
//		fmt.Println(userFile,err)
//		return
//	}
//	buf := make([]byte, 1024)
//	for{
//		n, _ := fin.Read(buf)
//		if 0== n {
//			break
//		}
//		os.Stdout.Write(buf[:n])
//	}
//
////	third.CpFile(1024);
//
//	var t int64 = time.Now().Unix()
//	var s string = time.Unix(t, 0).Format("2006-01-02 15:04:05")
//	println(s)

//	var t int64 = time.Now().Unix()
//	filename_in := "aa.png"
//	fi,err := os.Open(filename_in)
//	if err != nil{
//		panic(err)
//	}
//	defer fi.Close()
//
//	filename_out := "aabak.png"
//	fo,err := os.Create(filename_out)
//	if err != nil{
//		panic(err)
//	}
//	defer fo.Close()
//
//
//	var buff = make([]byte,10)
//	for{
//		n,err := fi.Read(buff)
//		if err != nil && err != io.EOF{
//			panic(err)
//		}
//
//		if n == 0{
//			break
//		}
//
//		if _,err := fo.Write(buff[:n]); err != nil{
//			panic(err)
//		}
//
//	}
//	var t1 int64 = time.Now().Unix()
//	fmt.Println(t1-t)

}
