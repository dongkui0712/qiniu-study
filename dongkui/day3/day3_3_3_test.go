package day3
import "testing"

func TestWhichFirst3(t *testing.T) {

	num := 3
	if result := WhichFirst3(num); result ==2 / 3 {
		t.Log("return first,result=",result)
	}else if result := WhichFirst3(num); result == 1 {
		t.Log("deger first,result=", result)
	}else {
		t.Error("failed")
	}
}
