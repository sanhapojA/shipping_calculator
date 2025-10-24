package domain

import "errors"

// Validation errors
var (
	ErrWeightInvalid       = errors.New("weight must be greater than 0")
	ErrDistanceNegative    = errors.New("distance cannot be negative")
	ErrDeliveryTypeInvalid = errors.New("invalid delivery type")
	ErrInvalidPromoCode    = errors.New("invalid promotion code")
)
