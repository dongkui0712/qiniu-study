package day3
import (
"fmt"
"io"
"os"
"log"
)

func noCopy() {
	var src, dst *os.File
	src, err := os.Open("README.md")
	if err == nil {
		log.Println("hava opened README.md")
	}

	defer func() {
		src.Close()
		dst.Close()
	}()

	dst, err1 := os.OpenFile("README2.md", os.O_WRONLY | os.O_CREATE, 0644)
	if err1 == nil {
		log.Println("hava opened README2.md")
	}
	//_, err := io.Copy(dst, src)
	//fmt.Println(err)
	buf := make([]byte, 1024)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, _ := dst.Write(buf[0:nr])
			fmt.Println("write size:", nw)
		}
		if er == io.EOF {
			break
		}
	}

}

