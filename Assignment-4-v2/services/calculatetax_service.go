package services

import "gorm-orm-connection/models"

const (
	raw          = "raw"
	manufactured = "manufactured"
	imported     = "imported"
)

func CalculateTaxAndFinalPrice(o models.Order) (tax float64, finalPrice float64) {
	t, p, iq := o.ItemType, o.ItemPrice, o.ItemQuantity
	switch t {
	case raw:
		tax = (p * 12.5) / 100
	case manufactured:
		tpf := (p * 12.5) / 100
		tax = tpf + (2*(p+tpf))/100
	case imported:
		importDuty := (10 * p) / 100
		finalCost := importDuty + ((p * 12.5) / 100) + p
		var surchage float64
		if finalCost <= 100 {
			surchage = 5
		} else if finalCost <= 200 {
			surchage = 10
		} else {
			surchage = (5 * finalCost) / 100
		}
		tax = importDuty + surchage
	}
	tax *= float64(iq)
	finalPrice = ((p + tax) * float64(iq))
	return
}
