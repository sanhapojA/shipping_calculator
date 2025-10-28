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

	t.Run("error weight < 0", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       -10.0,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		_, err := calc.Calculate(shipment)
		assert.Error(t, err)
		assert.ErrorContains(t, err, domain.ErrWeightInvalid.Error())
	})

}

func TestShippingFeeCalculator(t *testing.T) {
	t.Run("Calculate distance 30 km and weight 0.5 should be fee 30 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       0.5,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 30.0, fee)
	})

	t.Run("Calculate distance 50 km and weight 1.5 should be fee 50 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       1.5,
			Distance:     50,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 50.0, fee)
	})

	t.Run("Calculate distance 90 km and weight 5.5 should be fee 120 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       5.5,
			Distance:     90,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 120.0, fee)
	})

	t.Run("Calculate distance 160 km and weight 12.0 should be fee 150 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       12.0,
			Distance:     160,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 150.0, fee)
	})

	t.Run("Calculate distance 200 km and weight 15.0 should be fee 195 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       15.0,
			Distance:     200,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 195.0, fee)
	})

	t.Run("Calculate distance 500 km and weight 15.0 should be fee 260 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       15.0,
			Distance:     500,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)

		assert.NoError(t, err)
		assert.Equal(t, 260.0, fee)
	})
}

func TestDeliveryType(t *testing.T) {
	t.Run("Calculate fee based on delivery type is standard, weight is 0.5, distance is 30 should be fee 30 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       0.5,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 30.0, fee)
	})

	t.Run("Calculate fee based on delivery type is standard, weight is 5.5, distance is 90 should be fee 120 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       5.5,
			Distance:     90,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 120.0, fee)
	})

	t.Run("Calculate fee based on delivery type is standard, weight is 15.0, distance is 200 should be fee 195 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       15.0,
			Distance:     200,
			DeliveryType: domain.DeliveryTypeStandard,
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 195.0, fee)
	})

	t.Run("Calculate fee based on delivery type is express, weight is 1.5, distance is 50 should be fee 75 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       1.5,
			Distance:     50,
			DeliveryType: domain.DeliveryTypeExpress,
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 75.0, fee)
	})

	t.Run("Calculate fee based on delivery type is express, weight is 12.0, distance is 160 should be fee 225 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       12.0,
			Distance:     160,
			DeliveryType: domain.DeliveryTypeExpress,
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 225.0, fee)
	})

	t.Run("Calculate fee based on delivery type is express, weight is 15.0, distance is 500 should be fee 390 THB", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       15.0,
			Distance:     500,
			DeliveryType: domain.DeliveryTypeExpress,
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 390.0, fee)
	})
}

func TestCalculateWithPromotion(t *testing.T) {
	t.Run("Calculate fee base on promotion code with weight is 0.5 and distance is 30 then fee should be 0", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       0.5,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
			PromoCode:    "FREESHIP50",
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 0.0, fee)
	})

	t.Run("Calculate fee base on promotion code with weight is 1.5 and distance is 50 then fee should be 0", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       1.5,
			Distance:     50,
			DeliveryType: domain.DeliveryTypeStandard,
			PromoCode:    "FREESHIP50",
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 0.0, fee)
	})

	t.Run("Calculate fee base on promotion code with weight is 5.5 and distance is 30 then fee should be 30", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       5.5,
			Distance:     30,
			DeliveryType: domain.DeliveryTypeStandard,
			PromoCode:    "FREESHIP50",
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 30.0, fee)
	})

	t.Run("Calculate fee base on promotion code with weight is 12.0 and distance is 160 then fee should be 50", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       12.0,
			Distance:     160,
			DeliveryType: domain.DeliveryTypeStandard,
			PromoCode:    "FREESHIP100",
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 50.0, fee)
	})

	t.Run("Calculate fee base on promotion code with weight is 15.0 and distance is 200 then fee should be 95", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       15.0,
			Distance:     200,
			DeliveryType: domain.DeliveryTypeStandard,
			PromoCode:    "FREESHIP100",
		}

		fee, err := calc.Calculate(shipment)
		assert.NoError(t, err)
		assert.Equal(t, 95.0, fee)
	})

	t.Run("Calculate fee error with invalid promo code", func(t *testing.T) {
		calc := NewShippingCalculator()
		shipment := &domain.Shipment{
			Weight:       5.5,
			Distance:     90,
			DeliveryType: domain.DeliveryTypeStandard,
			PromoCode:    "FREESHIP250",
		}

		_, err := calc.Calculate(shipment)
		assert.Error(t, err)

	})
}

func TestCalculateWithPromotionOrderAmountDiscount(t *testing.T) {
	t.Run("Calculate with order amount is 1500 and weight 0.5 and distance 30 should be fee 24 thb", func(t *testing.T) {

	})
}
