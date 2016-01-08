package main
import (
	"os"
	"io"
	"fmt"
)


func main() {
	w,err := CopyFile("/Users/Apple_Xuan/Golang/src.txt", "/Users/Apple_Xuan/Golang/des.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(w)
}



func CopyFile(src, des string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}
