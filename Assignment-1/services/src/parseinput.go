package src

import (
	"bufio"
	"errors"
	"fmt"
	"nuclei-assignment-1/utils"
	"os"
	"strings"
)

const (
	InputValueErrContent = `too less argument`
	Type                 = "-type"
	Price                = "-price"
	Quantity             = "-quantity"
	AlreadyDefined       = "%s field is already defined\n"
)

func ParseInputName() (name string) {
	reader := bufio.NewReader(os.Stdin)
	PrintInputNameDetails()
	for name == "" {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err, "error occured while reading string")
		} else {
			name, err = utils.ValidateName(text)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}

func ParseInputString() (itype string, iprice float64, iquantity int) {
	PrintInputDetails()
	reader := bufio.NewReader(os.Stdin)
naming:
	for itype == "" || iprice == 0 || iquantity == 0 {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err, "error occured while reading string")
		}
		textfields := strings.Fields(text)
		l := len(textfields)
		if l != 2 {
			err := errors.New(InputValueErrContent)
			utils.InputValueError(err)
			continue naming
		}
		t := textfields[0]
		switch t {
		case Type:
			if itype != "" {
				fmt.Printf(AlreadyDefined, Type)
				break
			}
			out, err := utils.ValidateType(textfields[1])
			if err != nil {
				fmt.Println(err)
				break
			}
			itype = out
		case Price:
			if iprice != 0. {
				fmt.Printf(AlreadyDefined, Price)
				break
			}
			out, err := utils.ValidatePrice(textfields[1])
			if err != nil {
				fmt.Println(err)
				break
			}
			iprice = out
		case Quantity:
			if iquantity != 0 {
				fmt.Printf(AlreadyDefined, Quantity)
				break
			}
			out, err := utils.ValidateQuantity(textfields[1])
			if err != nil {
				fmt.Println(err)
				break
			}
			iquantity = out
		default:
			err := errors.New("undefined input")
			utils.InputValueError(err)
			continue naming
		}
	}
	return itype, iprice, iquantity
}
