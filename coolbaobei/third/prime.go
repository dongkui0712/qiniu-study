package third
import (
		"math")
func JudgPrime(value int) bool{
//	fmt.Println("input number is :",a)
//	k := (int)(math.Sqrt(float64(a)))
//	fmt.Println(k)
//	var n int
//	for i := 2;i<k;i++{
//		if a%i == 0{
//			fmt.Println(n)
//			n++
//			break
//		}
//		if i > k{
//			fmt.Println(a,"%d is prime")
//		}else{
//			fmt.Println(a,"%d isn't prime")
//		}
//	}
//	fmt.Println(n)
//	return err

	if value <= 1 {
		return false
	}
	if value == 2 || value == 3 || value == 5 || value == 7 {
		return true
	}
	if value%2 == 0 || value%3 == 0 || value%5 == 0 || value%7 == 0 {
		return false
	}
	factor := 7
	c := []int{4, 2, 4, 2, 4, 6, 2, 6}
	max := int(math.Sqrt(float64(value)))
	if max*max == value {
		return false
	}
	for factor < max {
		for i := 0; i < len(c); i++ {
			factor += c[i]
			if value%factor == 0 {
				return false
			}
		}
	}
	return true
}