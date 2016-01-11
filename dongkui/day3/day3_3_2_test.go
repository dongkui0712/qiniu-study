package day3
import "testing"

func TestWhichFirst(t *testing.T) {

	if result := WhichFirst(); result ==0 {
		t.Log("return first,result=",result)
	}else if result == 3{
		t.Log("defer first,result=",result)
	}else {
		t.Error("failed")
	}
}
