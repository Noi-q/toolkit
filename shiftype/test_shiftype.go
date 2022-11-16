package shiftype

import (
	"strconv"
)

func TestShiftStrInt(s string) (int, error) {
	si, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return si, nil
}
