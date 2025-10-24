package calculator

import "shipping-calculator/internal/domain"

type ShippingCalculator struct{}

func NewShippingCalculator() *ShippingCalculator {
	return &ShippingCalculator{}
}

func (c *ShippingCalculator) Calculate(shipment *domain.Shipment) (float64, error) {
	// Calculate base fee by weight
	var baseFee float64
	if shipment.Weight <= 1 {
		baseFee = 30
	} else if shipment.Weight <= 5 {
		baseFee = 50
	} else if shipment.Weight <= 10 {
		baseFee = 80
	} else {
		excess := shipment.Weight - 10
		baseFee = 80 + (excess * 10)
	}

	// Apply distance multiplier
	var multiplier float64
	if shipment.Distance <= 50 {
		multiplier = 1.0
	} else if shipment.Distance <= 200 {
		multiplier = 1.5
	} else {
		multiplier = 2.0
	}

	fee := baseFee * multiplier

	// Apply express delivery surcharge
	if shipment.DeliveryType == domain.DeliveryTypeExpress {
		fee = fee * 1.5
	}

	return fee, nil
}
