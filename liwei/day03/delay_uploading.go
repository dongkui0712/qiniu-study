package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func main() {

	//srcfile
	srcf, err := os.Open("a.csv")
	tof, err := os.OpenFile("b.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)

	//Get file size
	fstat, err := srcf.Stat()

	if err != nil {
		panic(err)
	}

	fmt.Printf("source file length: %d\n", fstat.Size())

	srcBuf := make([]byte, fstat.Size())
	srcBufLen, err := srcf.Read(srcBuf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if 0 == srcBufLen {
		fmt.Print("length ==0")
	}

	defer srcf.Close()
	defer tof.Close()

	var step int = 100

	var count int = 0

	ticker := time.NewTicker(time.Second * 1)
	quit := make(chan bool)
	var mutex = &sync.Mutex{}

	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				mutex.Lock()
				tof.Write(srcBuf[count*step : step*(count+1)])
				var pro = float64(step*count) / float64(srcBufLen) * 100
				s := fmt.Sprintf("%.5f", pro)
				fmt.Println("uploaded...", s, "%")
				count++
				mutex.Unlock()

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(time.Duration(srcBufLen/step) * time.Second)
	quit <- true
}
