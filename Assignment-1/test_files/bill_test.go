package testfiles

import (
	"nuclei-assignment-1/models"
	"testing"
)

var billTests = []struct {
	ittype      string
	itprice     float64
	itquantity  int
	salesTTax   float64
	finalTPrice float64
}{
	{"raw", 100, 1, 12.8, 78},
	{"manufactured", 120, 2, 23, 121},
}

func BillTest(t *testing.T) {
	for _, test := range billTests {
		actualSalesTax, actualFinalPrice := models.
			CalculateTaxAndFinalPrice(test.ittype, test.itprice, test.itquantity)
		if actualSalesTax != test.salesTTax || actualFinalPrice != test.finalTPrice {
			t.Errorf("expected %f and %f got %f and %f",
				test.salesTTax, test.finalTPrice, actualSalesTax, actualFinalPrice)
		}
	}
}
