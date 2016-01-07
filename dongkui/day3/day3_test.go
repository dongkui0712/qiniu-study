package day3
import "testing"

func TestGetPri_2(t *testing.T)  {

	if result,no,err := GetPri(2); err == nil {
		t.Log("TestGetPri_2 passed,result=",result,"no=",no)
	}else {
		t.Error("TestGetPri_2 failed,err:",err,";no=",no)
	}
}

func TestGetPri_3(t *testing.T)  {

	if result,no,err := GetPri(3);err == nil {
		t.Log("TestGetPri_3 passed,result=",result,"no=",no)
	}else {
		t.Error("TestGetPri_3 failed,err:",err,";no=",no)
	}
}

func TestGetPri_100(t *testing.T)  {

	if result,no,err := GetPri(100);err == nil {
		t.Log("TestGetPri_100 passed,result=",result,"no=",no)
	}else {
		t.Error("TestGetPri_100 failed,err:",err,";no=",no)
	}
}

