package src

func ReadInput() (string, string, float64, int) {
	iname := ParseInputName()
	itype, iprice, iquantity := ParseInputString()
	return iname, itype, iprice, iquantity
}
