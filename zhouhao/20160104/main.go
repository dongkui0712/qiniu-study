package main
import (
	"fmt"
	"./division"
)
func main()  {
	result,err:=division.Division(8,1)
	if(err==nil){
		fmt.Println("结果是",result)
	}else {
		fmt.Println("出错鸟,错误原因是:",err)

	}

}
