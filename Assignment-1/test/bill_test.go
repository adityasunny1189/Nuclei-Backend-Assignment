package test

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
	{
		ittype:      "raw",
		itprice:     100,
		itquantity:  1,
		salesTTax:   12.5,
		finalTPrice: 112.5,
	},
	{
		ittype:      "manufactured",
		itprice:     120,
		itquantity:  2,
		salesTTax:   35.4,
		finalTPrice: 310.8,
	},
}

func TestBill(t *testing.T) {
	for _, test := range billTests {
		actualSalesTax, actualFinalPrice := models.
			CalculateTaxAndFinalPrice(test.ittype, test.itprice, test.itquantity)
		if actualSalesTax != test.salesTTax || actualFinalPrice != test.finalTPrice {
			t.Errorf("expected %f and %f got %f and %f",
				test.salesTTax, test.finalTPrice, actualSalesTax, actualFinalPrice)
		}
	}
}
