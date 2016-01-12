package day3
import (
	"fmt"
	"log"
	"time"
)

func WhichFaster() {
	fmt.Println("start")
	log.Println("=== returnFunc():\n",returnFunc(1),"\nreturnFunc() end time:",time.Now())

	for i := 0;i<3;i++ {
		time.Sleep(time.Second)
		fmt.Println(".")
	}
	fmt.Println("end of WhichFaster()\n",time.Now())
}

func returnFunc(num int) string {

	defer func(){
		log.Println("<defer> === printTime():", printTime(), "\n\n")
	}()
	fmt.Println("    ========sleep 3s",time.Now())
	for i := 0;i<3;i++ {
		time.Sleep(time.Second)
		fmt.Println("    .")
	}

	log.Println("====log middle",time.Now())
	for i := 0;i<3;i++ {
		time.Sleep(time.Second)
		fmt.Println("    .")
	}

	return "\ntime of return:\n"+time.Now().String()+"\n\n<string returned>\n"
}

func printTime() string {
	return  fmt.Sprintf("\nprintTime:",time.Now())
}