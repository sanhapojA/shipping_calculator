package calculator

// Weight-based base fees (THB)
const (
	BaseFeeUnder1Kg  = 30.0
	BaseFee1To5Kg    = 50.0
	BaseFee5To10Kg   = 80.0
	BaseFeeOver10Kg  = 80.0
	ExcessWeightRate = 10.0 // THB per kg over 10kg
)

// Distance thresholds (km)
const (
	ShortDistanceLimit  = 50.0
	MediumDistanceLimit = 200.0
)

// Distance multipliers
const (
	ShortDistanceMultiplier  = 1.0
	MediumDistanceMultiplier = 1.5
	LongDistanceMultiplier   = 2.0
)

// Express delivery surcharge
const (
	ExpressSurchargeRate = 1.5 // +50%
)

// Order amount discount
const (
	LargeOrderThreshold = 1000.0
	LargeOrderDiscount  = 0.8 // 20% off
)

// VIP membership discounts
const (
	VIPGoldDiscount = 0.9 // 10% off
)

// Promotion codes and their discount amounts (THB)
var PromotionDiscounts = map[string]float64{
	"FREESHIP50":  50.0,
	"FREESHIP100": 100.0,
}
