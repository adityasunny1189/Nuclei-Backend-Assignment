package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateName(t string) (string, bool) {
	textfields := strings.Fields(t)
	if len(textfields) != 2 || textfields[0] != "-name" {
		InputNameError()
		return "", false
	}
	name := textfields[1]
	_, err := strconv.Atoi(name)
	if err == nil {
		StringError()
		InputNameError()
		return "", false
	} else {
		fmt.Printf("Name accepted✅\n")
		return name, true
	}
}

func ValidateType(t string) (string, bool) {
	switch t {
	case "raw", "manufactured", "imported":
		fmt.Printf("Type accepted✅\n")
		return t, true
	default:
		fmt.Printf("❌type can be only <raw || imported || manufactured>\n")
		return "", false
	}
}

func ValidatePrice(t string) (float64, bool) {
	ip, err := strconv.ParseFloat(t, 64)
	if err != nil {
		ParseFloatError(err)
		return 0., false
	} else if ip < 1 {
		fmt.Printf("❌Price can't be less than 1\n")
		return 0., false
	} else {
		fmt.Printf("Price accepted✅\n")
		return ip, true
	}
}

func ValidateQuantity(t string) (int, bool) {
	iq, err := strconv.Atoi(t)
	if err != nil {
		ParseIntError(err)
		return 0, false
	} else if iq < 1 {
		fmt.Printf("❌quantity can't be less than 1\n")
		return 0, false
	} else {
		fmt.Printf("Quantity accepted✅\n")
		return iq, true
	}
}
