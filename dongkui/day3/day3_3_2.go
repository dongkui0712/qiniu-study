package day3
import "fmt"

func WhichFirst() int {
	return getResult(0)
}

func getResult(num int) int {
	defer func() {
		num ++
		fmt.Println("3rd:",num)
	}()

	defer func() {
		num ++
		fmt.Println("2nd:",num)
	}()

	defer func() {
		num ++
		fmt.Println("1st:",num)
	}()

	return num
}