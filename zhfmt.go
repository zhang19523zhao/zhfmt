package util

import (
	"fmt"
	"strconv"
	"unicode"
)

// 判断是否有中文标点符号
func Ispunc(v rune) bool {
	punctuations := "，。！；’ “）（"
	for _, p := range punctuations {
		if p == v {
			return true
		}
	}
	return false
}

// 判断汉字个数, 返回汉字个数
func pd(str string) int {
	count := 0
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
		} else if Ispunc(v) {
			count++
		}
	}
	return count
}

// 计算长度减去汉字个书的值，返回格式化字符串
func getFmtStr(str []string, l int, separator string, direction int) string {
	s := ""
	if direction == 1 {
		for _, str := range str {
			n := pd(str)
			endL := strconv.Itoa(l - n)
			s += "%" + endL + "s" + separator
		}
	} else if direction == -1 {
		for _, str := range str {
			n := pd(str)
			endL := strconv.Itoa(l - n)
			s += "%" + "-" + endL + "s" + separator
		}
	}

	return s
}

// 一个一个的计算格式化字符串
func getOneFmtStr(str []string, separator string, direction int, l ...int) string {
	s := ""
	if direction == 1 {
		for i, str := range str {
			n := pd(str)
			endL := strconv.Itoa(l[i] - n)
			s += "%" + endL + "s" + separator
		}
	} else if direction == -1 {
		for i, str := range str {
			n := pd(str)
			endL := strconv.Itoa(l[i] - n)
			s += "%" + "-" + endL + "s" + separator
		}
	}
	return s
}

// 传递给函数一个数据切片，分隔符，和每个字段分隔的距离，如果只写一个默认为全部字段的分隔距离
// ["id", "姓名", "年龄", "电话", "地址"]
// direction: 1为右对齐，-1位左对齐
func Zhfmt(data []string, separator string, direction int, width ...int) (string, error) {
	fmtStr := ""
	strs := ""

	if len(width) != 1 && len(width) != len(data) {
		//fmt.Printf("有%d个字段，但给了%d个width\n", len(data), len(width))
		return "", fmt.Errorf("有%d个字段，但给了%d个width\n", len(data), len(width))
	}
	if len(width) == 1 {
		fmtStr = getFmtStr(data, width[0], separator, direction)
	} else {
		fmtStr = getOneFmtStr(data, separator, direction, width...)
	}

	args := make([]interface{}, len(data))
	for i := range data {
		args[i] = data[i]
	}
	//fmt.Printf(fmtStr, args...)
	strs = fmt.Sprintf(fmtStr, args...)
	strs = strs + "\n"

	//返回格式化的字符串
	return strs, nil
}
