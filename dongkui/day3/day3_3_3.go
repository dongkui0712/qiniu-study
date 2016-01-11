package day3

func WhichFirst3(num int) int {
	return getResult3(num)
}

func getResult3(num int) int {
	defer func() {
		num --
	}()

	return 2 / num
}