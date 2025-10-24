package calculator

type ShippingCalculator struct{}

func NewShippingCalculator() *ShippingCalculator {
	return &ShippingCalculator{}
}

// func (c *ShippingCalculator) Calculate(shipment *domain.Shipment) (float64, error) {
// 	if shipment.Weight <= 1 {
// 		return 30, nil
// 	}

// 	return 0, nil
// }
