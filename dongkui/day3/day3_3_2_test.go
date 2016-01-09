package day3
import "testing"

func TestWhichFirst(t *testing.T) {

	if result := WhichFirst(); result ==0 {
		t.Log("return first")
	}else if result == 3{
		t.Log("defer first",result)
	}else {
		t.Error("failed")
	}
}
