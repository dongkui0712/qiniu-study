package email

import (
	"testing"
	"fmt"

)


func Test_MyEml(email1 *testing.T){
//	mycontent := " my dear令"

//	email := NewEmail("wanbao.emailyang@changhong.com;598145534@qq.com;",
//		"test golang email", mycontent)

	err := SendEmail(email1)

	fmt.Println(err)
}
