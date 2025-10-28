package calculator

import (
	"shipping-calculator/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightBasedPricing(t *testing.T) {
	// Given
	t.Run("Weight less than 1 kg should cost 30 THB", func(t *testing.T) {
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
	})

	t.Run("Weight equal 1.5 and distance 30km should be fee 50 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       1.5,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 50.0, fee)
	})

	t.Run("Weight equal 5.5 and distance 30km should be fee 80 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       5.5,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 80.0, fee)
	})

	t.Run("Weight equal 12.0 and distance 30km should be fee 100 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       12.0,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 100.0, fee)
	})

	t.Run("Weight equal 15.0 and distance 30km should be fee 130 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       15.0,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 130.0, fee)
	})

}
