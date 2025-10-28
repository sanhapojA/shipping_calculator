package calculator

import "shipping-calculator/internal/domain"

type ShippingCalculator struct{}

func NewShippingCalculator() *ShippingCalculator {
	return &ShippingCalculator{}
}

func (c *ShippingCalculator) Calculate(shipment *domain.Shipment) (float64, error) {
	weight := shipment.Weight

	if weight > 0 && weight <= 1 {
		return 30.0, nil
	} else if weight > 1 && weight <= 5 {
		return 50.0, nil
	} else if weight > 5 && weight <= 10 {
		return 80.0, nil
	} else if weight > 10 {
		return 80.0 + ((weight - 10) * 10), nil
	}

	return 0, nil
}
