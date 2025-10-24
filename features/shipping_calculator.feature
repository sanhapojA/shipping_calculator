# language: en
Feature: Shipping Fee Calculator
  As a customer
  I want to know the shipping fee before placing an order
  So that I can make informed purchasing decisions

  Background:
    Given the shipping calculator is ready

  # Scenario Group 1: Weight-Based Pricing
  Scenario Outline: Calculate shipping fee based on weight
    Given a shipment with weight <weight> kg
    And distance 30 km
    And delivery type "STANDARD"
    When I calculate the shipping fee
    Then the fee should be <fee> THB

    Examples:
      | weight | fee |
      | 0.5    | 30  |
      | 1.5    | 50  |
      | 5.5    | 80  |
      | 12.0   | 100 |
      | 15.0   | 130 |

  # Scenario Group 2: Distance Multiplier
  Scenario: Calculate shipping fee for medium distance
    Given a shipment with weight 1.5 kg
    And distance 150 km
    And delivery type "STANDARD"
    When I calculate the shipping fee
    Then the fee should be 75 THB

  Scenario: Calculate shipping fee for long distance
    Given a shipment with weight 1.5 kg
    And distance 250 km
    And delivery type "STANDARD"
    When I calculate the shipping fee
    Then the fee should be 100 THB

  # Scenario Group 3: Express Delivery
  Scenario: Calculate express delivery fee
    Given a shipment with weight 1.5 kg
    And distance 30 km
    And delivery type "EXPRESS"
    When I calculate the shipping fee
    Then the fee should be 75 THB

  # Scenario Group 4: Promotions
  Scenario: Apply FREESHIP50 promotion code
    Given a shipment with weight 5.5 kg
    And distance 30 km
    And delivery type "STANDARD"
    And promotion code "FREESHIP50"
    When I calculate the shipping fee
    Then the fee should be 30 THB

  Scenario: Apply FREESHIP100 on small fee results in zero
    Given a shipment with weight 0.5 kg
    And distance 30 km
    And delivery type "STANDARD"
    And promotion code "FREESHIP100"
    When I calculate the shipping fee
    Then the fee should be 0 THB

  Scenario: Order amount over 1000 gets 20% discount
    Given a shipment with weight 5.5 kg
    And distance 30 km
    And delivery type "STANDARD"
    And order amount 1500 THB
    When I calculate the shipping fee
    Then the fee should be 64 THB

  # Scenario Group 5: VIP Membership
  Scenario: VIP Gold member gets 10% discount
    Given a shipment with weight 5.5 kg
    And distance 30 km
    And delivery type "STANDARD"
    And membership "VIP_GOLD"
    When I calculate the shipping fee
    Then the fee should be 72 THB

  Scenario: VIP Platinum member gets free standard shipping
    Given a shipment with weight 5.5 kg
    And distance 30 km
    And delivery type "STANDARD"
    And membership "VIP_PLATINUM"
    When I calculate the shipping fee
    Then the fee should be 0 THB

  Scenario: VIP Platinum member pays only difference for express
    Given a shipment with weight 5.5 kg
    And distance 30 km
    And delivery type "EXPRESS"
    And membership "VIP_PLATINUM"
    When I calculate the shipping fee
    Then the fee should be 40 THB

  # Scenario Group 6: Edge Cases
  Scenario: Zero weight returns error
    Given a shipment with weight 0 kg
    And distance 30 km
    When I calculate the shipping fee
    Then I should get an error "Weight must be greater than 0"

  Scenario: Negative weight returns error
    Given a shipment with weight -1 kg
    And distance 30 km
    When I calculate the shipping fee
    Then I should get an error "Weight must be greater than 0"

  Scenario: Negative distance returns error
    Given a shipment with weight 1.5 kg
    And distance -10 km
    When I calculate the shipping fee
    Then I should get an error "Distance cannot be negative"

  Scenario: Invalid promotion code returns error
    Given a shipment with weight 5.5 kg
    And distance 30 km
    And delivery type "STANDARD"
    And promotion code "INVALID"
    When I calculate the shipping fee
    Then I should get an error "Invalid promotion code"

  # Scenario Group 7: Complex Scenarios
  Scenario: Complex calculation with all factors
    Given a shipment with weight 12.0 kg
    And distance 150 km
    And delivery type "EXPRESS"
    And order amount 1500 THB
    And membership "VIP_GOLD"
    When I calculate the shipping fee
    Then the fee should be 162 THB

  Scenario: Boundary value - exactly 1 kg
    Given a shipment with weight 1.0 kg
    And distance 30 km
    And delivery type "STANDARD"
    When I calculate the shipping fee
    Then the fee should be 30 THB

  Scenario: Boundary value - exactly 5 kg
    Given a shipment with weight 5.0 kg
    And distance 30 km
    And delivery type "STANDARD"
    When I calculate the shipping fee
    Then the fee should be 50 THB

  Scenario: Boundary value - exactly 50 km
    Given a shipment with weight 1.5 kg
    And distance 50 km
    And delivery type "STANDARD"
    When I calculate the shipping fee
    Then the fee should be 50 THB
