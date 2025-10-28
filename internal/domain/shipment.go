package domain

import "time"

// Shipment represents a shipment order with all relevant details for fee calculation
type Shipment struct {
	Weight       float64        // Weight in kilograms
	Distance     float64        // Distance in kilometers
	DeliveryType DeliveryType   // Type of delivery service
	OrderAmount  float64        // Total order amount in THB
	PromoCode    string         // Promotion code (optional)
	Membership   MembershipTier // Customer membership tier
	CreatedAt    time.Time      // When the shipment was created
}

// NewShipment creates a new Shipment with validation
func NewShipment(weight, distance float64, deliveryType DeliveryType) (*Shipment, error) {
	shipment := &Shipment{
		Weight:       weight,
		Distance:     distance,
		DeliveryType: deliveryType,
		Membership:   MembershipNone,
		CreatedAt:    time.Now(),
	}

	if err := shipment.Validate(); err != nil {
		return nil, err
	}

	return shipment, nil
}

// Validate checks if the shipment data is valid
func (s *Shipment) Validate() error {
	if s.Weight <= 0 {
		return ErrWeightInvalid
	}

	if s.Distance < 0 {
		return ErrDistanceNegative
	}

	if s.DeliveryType != DeliveryTypeStandard && s.DeliveryType != DeliveryTypeExpress {
		return ErrDeliveryTypeInvalid
	}

	return nil
}

// WithPromoCode sets the promotion code for the shipment
func (s *Shipment) WithPromoCode(code string) *Shipment {
	s.PromoCode = code
	return s
}

// WithOrderAmount sets the order amount for the shipment
func (s *Shipment) WithOrderAmount(amount float64) *Shipment {
	s.OrderAmount = amount
	return s
}

// WithMembership sets the membership tier for the shipment
func (s *Shipment) WithMembership(tier MembershipTier) *Shipment {
	s.Membership = tier
	return s
}
