package day2
import (
	"fmt"
//	"log"
//	"errors"
//	"sort"
)


type Student struct {
	Name string
	Age  int
	Num  string
}

var Students []Student

func (stu *Student)SaveInf()  {

	Students = append(Students,*stu)

	Students = Sort(Students)
}

func Sort(stus []Student) []Student {

	c := cap(stus)
	if c > 1 {
		for i := 1;stus[c-1].Num < stus[c-i-1].Num;i++ {
			if i < c-1 {
				stus[c-1],stus[c-i-1] = stus[c-i-1],stus[c-1]
			}
		}
	}
	return stus
}

func GetStu() Student {
	var stu Student
	fmt.Println("Please input information of student:")
	fmt.Scanf("%s,%i,%s",&stu.Name,&stu.Age,&stu.Num)
	return stu
}

/*
func Sort2(stus []Student) ([]Student,error) {
	var nums []string
	var err error
	nums,_ = getSortedNums(stus)

	stusSorted := make([]Student,cap(stus)-1)

	var stu Student
	for _,c := range nums {
		stu,err = getStuByNum(c)
		stusSorted = append(stusSorted,stu)
	}
	return stusSorted,err
}

func getStuByNum(num string) (Student,error){
	var stu Student
	for _,c := range Students {
		if c.Num == num {
			stu = c
		}
	}
//	if stu == nil {
//		return nil,errors.New("failed to getStuByNum")
//	}
	fmt.Println("stu:",stu)
	return stu,nil
}

func getSortedNums(stus []Student) ([]string,error){
	var nums []string
	var err error
	for _,c := range stus {
		nums = append(nums,c.Num)
	}
	if nums == nil {
		err = errors.New("nums is nil")
		return nil,err
	}
	log.Println("nums before sorting:",nums)

	//升序排序
	sort.Strings(nums)

	log.Println("nums after sorting:",nums)
	return nums,err
}
*/