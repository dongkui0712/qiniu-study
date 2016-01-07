package main

import (
	"fmt"
)

func main() {
	findPrimeNumber(100)
}

func findPrimeNumber(n int) {
	var intIf int = 0
	var primes []int
	for i := 1; i < n; i++ {
		is, cIf := isPrime(i)
		intIf = intIf + cIf
		if is == true {
			intIf = intIf + cIf
			primes = append(primes, i)
		}
	}

	num := int(len(primes) )
	fmt.Println("num:",num)
	fmt.Println("最大素数为:", primes[num-1])
	fmt.Println("if执行次数:", intIf)
	//	for _,v := range primes {
	//
	//	}
}

func isPrime(value int) (is bool, intIf int) {
	intIf=0
	if value <= 3 {
		is = (value >= 2)
		intIf++
		return is, intIf
	}

	if value%2 == 0 || value%3 == 0 {
		is = false
		intIf++
		return is, intIf
	}

	for i := 5; i*i <= value; i += 6 {
		intIf=intIf+1
		if value%i == 0 || value%(i+2) == 0 {
			intIf=intIf+1
			return false, intIf
		}
	}

	return true, intIf
}

//func isPrimeB(value int) bool {
//	if value <= 3 {
//		return value >= 2
//	}
//	if value%2 == 0 || value%3 == 0 {
//		return false
//	}
//	for i := 5; i*i <= value; i += 6 {
//		if value%i == 0 || value%(i+2) == 0 {
//			return false
//		}
//	}
//	return true
//}
