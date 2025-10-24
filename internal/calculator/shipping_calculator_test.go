package calculator

import (
	"shipping-calculator/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate_WeightUnder1Kg_Returns30Baht(t *testing.T) {
	// Given
	calc := NewShippingCalculator()
	shipment := &domain.Shipment{
		Weight:       0.5,
		Distance:     30,
		DeliveryType: domain.DeliveryTypeStandard,
	}

	// When
	fee, err := calc.Calculate(shipment)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 30.0, fee)
}

func TestCalculate_Weight1To5Kg_Returns50Baht(t *testing.T) {
	// Given
	calc := NewShippingCalculator()
	shipment := &domain.Shipment{
		Weight:       1.5,
		Distance:     30,
		DeliveryType: domain.DeliveryTypeStandard,
	}

	// When
	fee, err := calc.Calculate(shipment)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 50.0, fee)
}

func TestCalculate_Weight5To10Kg_Returns80Baht(t *testing.T) {
	// Given
	calc := NewShippingCalculator()
	shipment := &domain.Shipment{
		Weight:       5.5,
		Distance:     30,
		DeliveryType: domain.DeliveryTypeStandard,
	}

	// When
	fee, err := calc.Calculate(shipment)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 80.0, fee)
}

func TestCalculate_Distance51To200Km_MultiplyBy1Point5(t *testing.T) {
	// Given: น้ำหนัก 1.5kg (50 บาท), ระยะทาง 150km
	calc := NewShippingCalculator()
	shipment := &domain.Shipment{
		Weight:       1.5,
		Distance:     150,
		DeliveryType: domain.DeliveryTypeStandard,
	}

	// When
	fee, err := calc.Calculate(shipment)

	// Then: 50 * 1.5 = 75
	assert.NoError(t, err)
	assert.Equal(t, 75.0, fee)
}

func TestCalculate_DistanceOver200Km_MultiplyBy2(t *testing.T) {
	// Given: น้ำหนัก 1.5kg (50 บาท), ระยะทาง 250km
	calc := NewShippingCalculator()
	shipment := &domain.Shipment{
		Weight:       1.5,
		Distance:     250,
		DeliveryType: domain.DeliveryTypeStandard,
	}

	// When
	fee, err := calc.Calculate(shipment)

	// Then: 50 * 2.0 = 100
	assert.NoError(t, err)
	assert.Equal(t, 100.0, fee)
}

func TestCalculate_ExpressDelivery_AddFiftyPercent(t *testing.T) {
	// Given: น้ำหนัก 1.5kg, ระยะ 30km, Express
	calc := NewShippingCalculator()
	shipment := &domain.Shipment{
		Weight:       1.5,
		Distance:     30,
		DeliveryType: domain.DeliveryTypeExpress,
	}

	// When
	fee, err := calc.Calculate(shipment)

	// Then: 50 * 1.5 = 75
	assert.NoError(t, err)
	assert.Equal(t, 75.0, fee)
}
