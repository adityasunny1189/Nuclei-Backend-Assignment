package src

import (
	"bufio"
	"fmt"
	"nuclei-assignment-1/utils"
	"os"
	"strings"
)

func ParseInputName() string {
	name := ""
	reader := bufio.NewReader(os.Stdin)
	InputNameDetails()
	for name == "" {
		text, _ := reader.ReadString('\n')
		name, _ = utils.ValidateName(text)
	}
	return name
}

func ParseInputString() (string, float64, int) {
	itype, iprice, iquantity := "", 0., 0
	InputDetails()
	reader := bufio.NewReader(os.Stdin)
naming:
	for itype == "" || iprice == 0. || iquantity == 0 {
		text, _ := reader.ReadString('\n')
		textfields := strings.Fields(text)
		l := len(textfields)
		if l != 2 {
			utils.InputValueError()
			continue naming
		}
		t := textfields[0]
		switch t {
		case "-type":
			if itype != "" {
				fmt.Printf("already mentioned\n")
				break
			}
			itype, _ = utils.ValidateType(textfields[1])
		case "-price":
			if iprice != 0. {
				fmt.Printf("already mentioned\n")
				break
			}
			iprice, _ = utils.ValidatePrice(textfields[1])
		case "-quantity":
			if iquantity != 0 {
				fmt.Printf("already mentioned\n")
				break
			}
			iquantity, _ = utils.ValidateQuantity(textfields[1])
		default:
			utils.InputValueError()
			continue naming
		}
	}
	return itype, iprice, iquantity
}
