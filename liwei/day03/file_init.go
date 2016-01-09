package main

import (
	"fmt"
	"os"
)

func main() {
	//srcfile
	srcf, err := os.OpenFile("a.csv", os.O_RDWR|os.O_CREATE, 0664)

	if err != nil {
		panic(err)
	}

	defer srcf.Close()
	//Get file size
	fstat, err := srcf.Stat()

	if fstat.Size() > 0 {
		fmt.Printf("source file has initialize: %d\n", fstat.Size())
		return
	}

	for i := 0; i < 10000; i++ {
		if _, err = srcf.WriteString("s,中,w\n"); err != nil {
			panic(err)
		}
	}

}
