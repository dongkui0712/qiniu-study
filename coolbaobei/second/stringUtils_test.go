package second

import (
	"testing"
	"fmt"
)

func TestCharachars(t *testing.T){
	str := "Hello world"

	m,n,err := Charachars(str)

	if err!=nil {
		t.Errorf("error")
	}
	if m == nil{
		t.Errorf("error")
	}else{
		t.Log("The charachars is %d",m)
		t.Log("The bytes is %d",n)
	}

//	a,err := ReverseByArray(str)
//	if err!=nil {
//		t.Errorf("error")
//	}
//	fmt.Println(a)


}
