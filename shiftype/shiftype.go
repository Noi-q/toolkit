package shiftype

import (
	"fmt"
	"strconv"
)

func Shiftype() {
	fmt.Println("Hello World")
}

// string数字 转 int
func ShiftStrInt(s string) (int, error) {
	si, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return si, nil
}

// string数字 转 int64
func ShiftStrInt64(s string) (int64, error) {
	si64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return si64, nil
}

// int 转 string
func ShiftIntStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

// int64 转 string
func ShiftInt64Str(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}
