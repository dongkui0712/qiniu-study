package guoxiangwen

import (
	"fmt"
	"github.com/guoxiangwen/golang-train/add"
)
func main() {
	s, err := add.Add(4, 4)
	if err != nil {
		fmt.Printf("error")
		return
	}
	fmt.Printf("the plus is:",s)
}
