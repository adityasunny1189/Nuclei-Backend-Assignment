package utils

import (
	"errors"
	"strconv"
	"strings"
)

const (
	InputNameErr = `❌Input Rejected
Usage: [-name <product name>]`
	StringErr           = `❌expected string`
	TypeErr             = `❌type can be only <raw || imported || manufactured>`
	Name                = `-name`
	Raw                 = "raw"
	Manufactured        = "manufactured"
	Imported            = "imported"
	NegativePriceErr    = `❌Price can't be less than 1`
	NegativeQuantityErr = `❌quantity can't be less than 1`
)

func ValidateName(t string) (name string, err error) {
	textfields := strings.Fields(t)
	if len(textfields) != 2 || textfields[0] != Name {
		err = errors.New(InputNameErr)
		return
	}
	out := textfields[1]
	_, er := strconv.Atoi(out)
	if er == nil {
		err = errors.New(StringErr)
		return
	}
	name = out
	return
}

func ValidateType(t string) (tpe string, err error) {
	switch t {
	case Raw, Manufactured, Imported:
		tpe = t
		return
	default:
		err = errors.New(TypeErr)
		return
	}
}

func ValidatePrice(t string) (ip float64, err error) {
	out, er := strconv.ParseFloat(t, 64)
	if er != nil {
		err = er
		return
	} else if out < 1 {
		err = errors.New(NegativePriceErr)
		return
	}
	ip = out
	return
}

func ValidateQuantity(t string) (iq int, err error) {
	out, er := strconv.Atoi(t)
	if er != nil {
		err = er
		return
	} else if out < 1 {
		err = errors.New(NegativeQuantityErr)
		return
	}
	iq = out
	return
}
