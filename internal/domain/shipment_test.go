package domain

// func TestNewShipment_ValidData_Success(t *testing.T) {
// 	// Given: valid shipment data
// 	weight := 5.5
// 	distance := 100.0
// 	deliveryType := DeliveryTypeStandard

// 	// When: creating a new shipment
// 	shipment, err := NewShipment(weight, distance, deliveryType)

// 	// Then: shipment should be created successfully
// 	assert.NoError(t, err)
// 	assert.NotNil(t, shipment)
// 	assert.Equal(t, weight, shipment.Weight)
// 	assert.Equal(t, distance, shipment.Distance)
// 	assert.Equal(t, deliveryType, shipment.DeliveryType)
// 	assert.Equal(t, MembershipNone, shipment.Membership)
// }

// func TestShipment_Validate(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		shipment      Shipment
// 		expectedError error
// 	}{
// 		{
// 			name: "valid standard shipment",
// 			shipment: Shipment{
// 				Weight:       1.5,
// 				Distance:     30,
// 				DeliveryType: DeliveryTypeStandard,
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name: "valid express shipment",
// 			shipment: Shipment{
// 				Weight:       5.0,
// 				Distance:     100,
// 				DeliveryType: DeliveryTypeExpress,
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name: "zero weight",
// 			shipment: Shipment{
// 				Weight:       0,
// 				Distance:     30,
// 				DeliveryType: DeliveryTypeStandard,
// 			},
// 			expectedError: ErrWeightInvalid,
// 		},
// 		{
// 			name: "negative weight",
// 			shipment: Shipment{
// 				Weight:       -1,
// 				Distance:     30,
// 				DeliveryType: DeliveryTypeStandard,
// 			},
// 			expectedError: ErrWeightInvalid,
// 		},
// 		{
// 			name: "negative distance",
// 			shipment: Shipment{
// 				Weight:       1.5,
// 				Distance:     -10,
// 				DeliveryType: DeliveryTypeStandard,
// 			},
// 			expectedError: ErrDistanceNegative,
// 		},
// 		{
// 			name: "invalid delivery type",
// 			shipment: Shipment{
// 				Weight:       1.5,
// 				Distance:     30,
// 				DeliveryType: "INVALID",
// 			},
// 			expectedError: ErrDeliveryTypeInvalid,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// When: validating shipment
// 			err := tt.shipment.Validate()

// 			// Then: check expected error
// 			if tt.expectedError == nil {
// 				assert.NoError(t, err)
// 			} else {
// 				assert.ErrorIs(t, err, tt.expectedError)
// 			}
// 		})
// 	}
// }

// func TestShipment_WithPromoCode(t *testing.T) {
// 	// Given: a shipment
// 	shipment := &Shipment{
// 		Weight:       1.5,
// 		Distance:     30,
// 		DeliveryType: DeliveryTypeStandard,
// 	}

// 	// When: setting promo code
// 	result := shipment.WithPromoCode("FREESHIP50")

// 	// Then: promo code should be set
// 	assert.Equal(t, "FREESHIP50", result.PromoCode)
// 	assert.Equal(t, shipment, result) // fluent interface
// }

// func TestShipment_WithOrderAmount(t *testing.T) {
// 	// Given: a shipment
// 	shipment := &Shipment{
// 		Weight:       1.5,
// 		Distance:     30,
// 		DeliveryType: DeliveryTypeStandard,
// 	}

// 	// When: setting order amount
// 	result := shipment.WithOrderAmount(1500)

// 	// Then: order amount should be set
// 	assert.Equal(t, 1500.0, result.OrderAmount)
// 	assert.Equal(t, shipment, result) // fluent interface
// }

// func TestShipment_WithMembership(t *testing.T) {
// 	// Given: a shipment
// 	shipment := &Shipment{
// 		Weight:       1.5,
// 		Distance:     30,
// 		DeliveryType: DeliveryTypeStandard,
// 	}

// 	// When: setting membership
// 	result := shipment.WithMembership(MembershipGold)

// 	// Then: membership should be set
// 	assert.Equal(t, MembershipGold, result.Membership)
// 	assert.Equal(t, shipment, result) // fluent interface
// }
