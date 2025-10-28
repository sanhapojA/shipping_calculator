package calculator

import (
	"shipping-calculator/internal/domain"
)

type ShippingCalculator struct{}

func NewShippingCalculator() *ShippingCalculator {
	return &ShippingCalculator{}
}

func (c *ShippingCalculator) Calculate(shipment *domain.Shipment) (float64, error) {
	weight := shipment.Weight
	distance := shipment.Distance

	var baseFee float64
	switch {
	case weight > 0 && weight <= 1:
		baseFee = 30.0
	case weight > 1 && weight <= 5:
		baseFee = 50.0
	case weight > 5 && weight <= 10:
		baseFee = 80.0
	case weight > 10:
		baseFee = 80.0 + (weight-10.0)*10.0
	default:
		return 0, domain.ErrWeightInvalid
	}

	multiplier := 1.0
	if distance > 0 && distance <= 50 {
		multiplier = 1.0
	} else if distance > 50 && distance <= 200 {
		multiplier = 1.5
	} else {
		multiplier = 2.0
	}
	deliveryType := shipment.DeliveryType
	additionalWithDeliveryType := CalculateWithDeliveryType(deliveryType)

	fee := baseFee * multiplier * additionalWithDeliveryType

	fee, err := CalculateWithPromotionCode(fee, string(shipment.PromoCode))
	if err != nil {
		return 0, err
	}
	return fee, nil
}

func CalculateWithDeliveryType(deliveryType domain.DeliveryType) float64 {
	if deliveryType == domain.DeliveryTypeExpress {
		return 1.5
	}
	return 1.0
}

func CalculateWithPromotionCode(baseFee float64, promoCode string) (float64, error) {
	fee := baseFee
	if promoCode != "" {
		discount := PromotionDiscounts[promoCode]
		if discount == 0 {
			return 0, domain.ErrInvalidPromoCode
		}
		fee = fee - discount
		if fee < 0 {
			fee = 0
		}
	}

	return fee, nil
}
