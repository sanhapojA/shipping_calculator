package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"

// 	"shipping-calculator/internal/calculator"
// 	"shipping-calculator/internal/domain"
// )

// func main() {
// 	// Define command-line flags
// 	weight := flag.Float64("weight", 0, "Package weight in kg")
// 	distance := flag.Float64("distance", 0, "Delivery distance in km")
// 	deliveryType := flag.String("delivery-type", "STANDARD", "Delivery type: STANDARD or EXPRESS")
// 	promoCode := flag.String("promo-code", "", "Promotion code (optional)")
// 	orderAmount := flag.Float64("order-amount", 0, "Order amount in THB (optional)")
// 	membership := flag.String("membership", "NONE", "Membership tier: NONE, VIP_GOLD, or VIP_PLATINUM")

// 	flag.Parse()

// 	// Validate required parameters
// 	if *weight <= 0 {
// 		fmt.Println("Error: weight must be greater than 0")
// 		flag.Usage()
// 		os.Exit(1)
// 	}

// 	if *distance < 0 {
// 		fmt.Println("Error: distance cannot be negative")
// 		flag.Usage()
// 		os.Exit(1)
// 	}

// 	// Create shipment
// 	shipment := &domain.Shipment{
// 		Weight:       *weight,
// 		Distance:     *distance,
// 		DeliveryType: domain.DeliveryType(*deliveryType),
// 		OrderAmount:  *orderAmount,
// 		PromoCode:    *promoCode,
// 		Membership:   domain.MembershipTier(*membership),
// 	}

// 	// Calculate shipping fee
// 	calc := calculator.NewShippingCalculator()
// 	fee, err := calc.Calculate(shipment)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// Display result
// 	fmt.Println("=================================")
// 	fmt.Println("  Shipping Fee Calculation")
// 	fmt.Println("=================================")
// 	fmt.Printf("Weight:         %.2f kg\n", *weight)
// 	fmt.Printf("Distance:       %.2f km\n", *distance)
// 	fmt.Printf("Delivery Type:  %s\n", *deliveryType)
// 	if *promoCode != "" {
// 		fmt.Printf("Promo Code:     %s\n", *promoCode)
// 	}
// 	if *orderAmount > 0 {
// 		fmt.Printf("Order Amount:   %.2f THB\n", *orderAmount)
// 	}
// 	if *membership != "NONE" {
// 		fmt.Printf("Membership:     %s\n", *membership)
// 	}
// 	fmt.Println("=================================")
// 	fmt.Printf("Shipping Fee:   %.0f THB\n", fee)
// 	fmt.Println("=================================")
// }
