package practice03

import (
	"fmt"
	"io/ioutil"
)

func CopyFile() {
	b, _ := ioutil.ReadFile("test.txt")
	err := ioutil.WriteFile("write.txt", b, 0644)
	if err == nil {
		fmt.Println("write success:")
	}
}
