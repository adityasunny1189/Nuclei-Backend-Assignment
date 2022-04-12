package models

type Bill struct {
	ItemName     string
	ItemPrice    float64
	SalesTax     float64
	FinalPrice   float64
	ItemQuantity int
}

func (b *Bill) Setter(i Item, q int) {
	b.ItemName = i.Name
	b.ItemPrice = i.Price
	b.ItemQuantity = q
	b.SalesTax, b.FinalPrice = CalculateTaxAndFinalPrice(i.ItemType, i.Price, b.ItemQuantity)
}

func TwelvePointFivePercentOfItemCost(p float64) float64 {
	return (p * 12.5) / 100
}

func CalculateTaxAndFinalPrice(t string, p float64, iq int) (float64, float64) {
	tax, finalPrice := 0., 0.
	switch t {
	case "raw":
		tax = TwelvePointFivePercentOfItemCost(p)
	case "manufactured":
		tpf := TwelvePointFivePercentOfItemCost(p)
		tax = tpf + (2*(p+tpf))/100
	case "imported":
		importDuty := (10 * p) / 100
		finalCost := importDuty + TwelvePointFivePercentOfItemCost(p) + p
		surchage := 0.
		if finalCost <= 100. {
			surchage = 5
		} else if finalCost <= 200. {
			surchage = 10
		} else {
			surchage = (5 * finalCost) / 100
		}
		tax = importDuty + surchage
	}
	tax *= float64(iq)
	finalPrice = ((p + tax) * float64(iq))
	return tax, finalPrice
}
