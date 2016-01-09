package test0105
import (
	"strings"
)

//字符串匹配算法
//func StrMatchA(s1, s2 string) (index []int, err error) {
//	if !strings.Contains(s1, s2) {
//		err = errors.New("字符串1不包含字符串2!")
//		return nil, err
//	}
//
//	count := int(strings.Count(s1, s2))
//	if count > 0 {
//		fmt.Println(s1, "包含", s2, "的位置个数有:", count)
//
//		inFirst := strings.Index(s1, s2)
//		index = append(index, inFirst)
//		count = count-1
//
//		inLast := strings.LastIndex(s1, s2)
//		if inLast != inFirst {
//			index = append(index, inLast)
//			count=count-1
//
//		}
//
//		return index, nil
//	} else {
//		return index, err
//	}
//}

func StrMatchB(str, subStr string, result *[]int) {
	first := strings.Index(str, subStr)
	if len(*result) > 0 {
		first = (*result)[len(*result)-1] + first + 1
	}
	*result = append(*result, first)

	strs := strings.Split(str, "")
	strslice := strs[first+1:]
	strs1 := strings.Join(strslice, "")
	if strings.Index(strs1, subStr) != -1 {
		StrMatchB(strs1, subStr, result)
	}

}

//func FindSubStrIndex(str, subStr string, result *[]int) {
//
//	first := strings.Index(str, subStr)
//
//	strs := strings.Split(str, "")
//	strslice := strs[first+1:]
//	strs1 := strings.Join(strslice, "")
//	if len(*result) > 0 {
//		first = (*result)[len(*result)-1] + first + 1
//	}
//	*result = append(*result, first)
//	if strings.Index(strs1, subStr) != -1 {
//		FindSubStrIndex(strs1, subStr, result)
//	}
//
//}