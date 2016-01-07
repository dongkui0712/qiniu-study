package day3
import "errors"

func GetPri(num rune) (rune,int,error) {
	var result rune
	var err error
	var no int
	if num < 3  {
		no++
		err = errors.New("num is not allowed")
		return -1,no,err
	}

	no++
	for i := num;i > 2;i-- {
		no++
		var temp int
		for j := rune(2);j < i;j++ {
			no++
			if i%j == 0 {
				no++
				temp = 1
			}
			no++
		}
		no++
		if temp == 0 {
			no++
			result = i
			break
		}
		no++
	}

	return result,no,err
}
