package utils

import "strconv"

func ValidateInput(ch string) (op int, err error) {
	res, err := strconv.Atoi(ch)
	if err != nil {
		return
	}
	op = res
	return
}
