package test

import (
	"nuclei-assignment-1/utils"
	"testing"
)

var tests = []struct {
	testid    int
	iname     string
	itype     string
	iprice    string
	iquantity string
	expected  bool
}{
	{
		testid:    1,
		iname:     "-name tshirt",
		itype:     "raw",
		iprice:    "12.2",
		iquantity: "2",
		expected:  true,
	},
	{
		testid:    2,
		iname:     "-name tshirt",
		itype:     "manufactured",
		iprice:    "192.2",
		iquantity: "4",
		expected:  true,
	},
	{
		testid:    3,
		iname:     "-name tshirt",
		itype:     "raw",
		iprice:    "12.2",
		iquantity: "-quantit 2",
		expected:  false,
	},
	{
		testid:    4,
		iname:     "-name t shirt",
		itype:     "imported",
		iprice:    "12.2",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    5,
		iname:     "-name tshirt",
		itype:     "funny",
		iprice:    "12.2",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    6,
		iname:     "-name tshirt",
		itype:     "-tye raw",
		iprice:    "12.2",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    7,
		iname:     "-name tshirt",
		itype:     "raw imported",
		iprice:    "12.2",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    8,
		iname:     "-name tshirt",
		itype:     "raw",
		iprice:    "one hundred",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    9,
		iname:     "-name pant",
		itype:     "raw",
		iprice:    "19",
		iquantity: "2",
		expected:  true,
	},
	{
		testid:    10,
		iname:     "-name pant",
		itype:     "",
		iprice:    "19",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    11,
		iname:     "-name",
		itype:     "-type",
		iprice:    "-price",
		iquantity: "-quantity",
		expected:  false,
	},
	{
		testid:    12,
		iname:     "-name",
		itype:     "manufactured",
		iprice:    "192.2",
		iquantity: "4",
		expected:  false,
	},
	{
		testid:    13,
		iname:     "",
		itype:     "-type",
		iprice:    "12.2",
		iquantity: "-quantit 2",
		expected:  false,
	},
	{
		testid:    14,
		iname:     "-name",
		itype:     "imported",
		iprice:    "12.2",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    15,
		iname:     "-name jacket",
		itype:     "-type",
		iprice:    "12.2",
		iquantity: "-quantity",
		expected:  false,
	},
	{
		testid:    16,
		iname:     "-name jacket",
		itype:     "-tye raw",
		iprice:    "12.2",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    17,
		iname:     "-name jacket",
		itype:     "",
		iprice:    "12.2",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    18,
		iname:     "-name jacket",
		itype:     "raw",
		iprice:    "",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    19,
		iname:     "-name pant",
		itype:     "raw",
		iprice:    "19",
		iquantity: "one",
		expected:  false,
	},
	{
		testid:    20,
		iname:     "",
		itype:     "",
		iprice:    "",
		iquantity: "",
		expected:  false,
	},
	{
		testid:    21,
		iname:     "-name jacket",
		itype:     "12",
		iprice:    "",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    22,
		iname:     "-name jacket",
		itype:     "raw",
		iprice:    "-type raw",
		iquantity: "2",
		expected:  false,
	},
	{
		testid:    23,
		iname:     "-name jacket",
		itype:     "raw",
		iprice:    "manufactured",
		iquantity: "-type raw",
		expected:  false,
	},
}

func TestInputString(t *testing.T) {
	for _, test := range tests {
		expectedVal := test.expected
		_, errname := utils.ValidateName(test.iname)
		_, errtype := utils.ValidateType(test.itype)
		_, errprice := utils.ValidatePrice(test.iprice)
		_, errquant := utils.ValidateQuantity(test.iquantity)
		actualValue := false
		if errname == nil && errtype == nil && errprice == nil && errquant == nil {
			actualValue = true
		}
		if expectedVal != actualValue {
			t.Errorf("error in test case %d", test.testid)
		}
	}
}
