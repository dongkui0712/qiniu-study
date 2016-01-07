package day3
import (
	"fmt"
	"log"
	"time"
)

func WhichFaster() {
	fmt.Println("start")
	log.Println("result of returnFunc():\n\n",returnFunc(1),"\n------returnFunc end time:------\ntime of returnFunc::",time.Now())

	for i := 0;i<4;i++ {
		time.Sleep(1000000000)
		fmt.Println(".")
	}
	fmt.Println("end of WhichFaster()\n",time.Now())
}

func returnFunc(num int) string {

	defer func(){
		log.Println("\n<<<<<<<<\ndefer\ntime of defer====", time.Now(), "\n>>>>>>\n\n")
	}()
	fmt.Println("    ========sleep 4s",time.Now())
	for i := 0;i<4;i++ {
		time.Sleep(1000000000)
		fmt.Println("    .")
	}

	log.Println("====log")
	for i := 0;i<4;i++ {
		time.Sleep(1000000000)
		fmt.Println("    .")
	}

	fmt.Println("    middle")
	for i := 0;i<4;i++ {
		time.Sleep(1000000000)
		fmt.Println("    .")
	}

	log.Println("====",time.Now())
	for i := 0;i<4;i++ {
		time.Sleep(1000000000)
		fmt.Println("    .")
	}

	return "<<<<\ntime of return:\n"+time.Now().String()+"\n\n====string returned====\n>>>>\n"
}