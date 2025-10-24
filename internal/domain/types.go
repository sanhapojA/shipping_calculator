package domain

// DeliveryType represents the type of delivery service
type DeliveryType string

const (
	DeliveryTypeStandard DeliveryType = "STANDARD"
	DeliveryTypeExpress  DeliveryType = "EXPRESS"
)

// MembershipTier represents the customer's membership level
type MembershipTier string

const (
	MembershipNone     MembershipTier = "NONE"
	MembershipGold     MembershipTier = "VIP_GOLD"
	MembershipPlatinum MembershipTier = "VIP_PLATINUM"
)
