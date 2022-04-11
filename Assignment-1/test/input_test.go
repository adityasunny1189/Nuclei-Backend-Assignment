package test

import (
	"nuclei-assignment-1/utils"
	"testing"
)

var tests = []struct {
	iname     string
	itype     string
	iprice    string
	iquantity string
	expected  bool
}{
	{
		iname:     "-name tshirt",
		itype:     "raw",
		iprice:    "12.2",
		iquantity: "2",
		expected:  true,
	},
	{
		iname:     "-name tshirt",
		itype:     "manufactured",
		iprice:    "192.2",
		iquantity: "4",
		expected:  true,
	},
	{"-name tshirt", "raw", "12.2", "-quantit 2", false},
	{"-name t shirt", "imported", "12.2", "2", false},
	{"-name tshirt", "funny", "12.2", "2", false},
	{"-name tshirt", "-tye raw", "12.2", "2", false},
	{"-name tshirt", "raw imported", "12.2", "2", false},
	{"-name tshirt", "raw", "one hundred", "2", false},
	{"-name pant", "raw", "19", "2", true},
	{"-name pant", "", "19", "2", false},
	{"-name", "-type", "-price", "-quantity", false},
	{"-name", "manufactured", "192.2", "4", false},
	{"", "-type", "12.2", "-quantit 2", false},
	{"-name", "imported", "12.2", "2", false},
	{"-name jacket", "-type", "12.2", "-quantity", false},
	{"-name jacket", "-tye raw", "12.2", "2", false},
	{"-name jacket", "", "12.2", "2", false},
	{"-name jacket", "raw", "", "2", false},
	{"-name pant", "raw", "19", "one", false},
	{"", "", "", "", false},
	{"-name jacket", "12", "", "2", false},
	{"-name jacket", "raw", "-type raw", "2", false},
	{"-name jacket", "raw", "manufactured", "-type raw", false},
}

func TestInputString(t *testing.T) {
	for _, test := range tests {
		expectedVal := test.expected
		testname := test.iname
		_, actualTestNameVal := utils.ValidateName(testname)
		_, actualTestTypeVal := utils.ValidateType(test.itype)
		_, actualTestPriceVal := utils.ValidatePrice(test.iprice)
		_, actualTestQuantityVal := utils.ValidateQuantity(test.iquantity)
		actualValue := actualTestNameVal && actualTestPriceVal && actualTestQuantityVal && actualTestTypeVal
		if expectedVal != actualValue {
			t.Errorf("expected %v got %v", expectedVal, actualValue)
		}
	}
}
