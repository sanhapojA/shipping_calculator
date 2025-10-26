Feature: คำนวณค่าจัดส่งสินค้าเพื่อให้ลูกค้าทราบค่าใช้จ่ายก่อนสั่งซื้อ

Background:
Given ระบบคำนวณค่าจัดส่งพร้อมใช้งาน

# Scenario Group 1: Weight-based pricing

Scenario Outline: คำนวณค่าจัดส่งตามน้ำหนัก (Standard, ระยะทางสั้น)
Given พัสดุมีน้ำหนัก <weight> กิโลกรัม
And ระยะทาง <distance> กิโลเมตร
And เลือกการจัดส่งแบบ "Standard"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ <fee> บาท

    Examples:
      | weight | distance | fee |
      | 0.5    | 30       | 30  |
      | 1.5    | 30       | 50  |
      | 5.5    | 30       | 80  |
      | 12.0   | 30       | 100 |

# Scenario Group 2: Distance multiplier

Scenario: คำนวณค่าจัดส่งระยะไกล
Given พัสดุมีน้ำหนัก 1.5 กิโลกรัม
And ระยะทาง 150 กิโลเมตร
And เลือกการจัดส่งแบบ "Standard"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 75 บาท # Calculation: 50 × 1.5 = 75

Scenario: คำนวณค่าจัดส่งระยะไกลมาก
Given พัสดุมีน้ำหนัก 1.5 กิโลกรัม
And ระยะทาง 250 กิโลเมตร
And เลือกการจัดส่งแบบ "Standard"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 100 บาท # Calculation: 50 × 2.0 = 100

# Scenario Group 3: Express delivery

Scenario: จัดส่งแบบ Express
Given พัสดุมีน้ำหนัก 1.5 กิโลกรัม
And ระยะทาง 30 กิโลเมตร
And เลือกการจัดส่งแบบ "Express"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 75 บาท # Calculation: 50 × 1.5 = 75

# Scenario Group 4: Promotions

Scenario: ใช้โค้ดส่วนลด FREESHIP50
Given พัสดุมีน้ำหนัก 5.5 กิโลกรัม
And ระยะทาง 30 กิโลเมตร
And เลือกการจัดส่งแบบ "Standard"
And ใช้โค้ดส่วนลด "FREESHIP50"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 30 บาท # Calculation: 80 - 50 = 30

Scenario: ใช้โค้ดส่วนลดที่มากกว่าค่าจัดส่ง
Given พัสดุมีน้ำหนัก 0.5 กิโลกรัม
And ระยะทาง 30 กิโลเมตร
And เลือกการจัดส่งแบบ "Standard"
And ใช้โค้ดส่วนลด "FREESHIP100"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 0 บาท # Calculation: 30 - 100 = 0 (min 0)

Scenario: สั่งซื้อเกิน 1000 บาทได้ส่วนลด 20%
Given พัสดุมีน้ำหนัก 5.5 กิโลกรัม
And ระยะทาง 30 กิโลเมตร
And เลือกการจัดส่งแบบ "Standard"
And ยอดสั่งซื้อ 1500 บาท
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 64 บาท # Calculation: 80 × 0.8 = 64

# Scenario Group 5: VIP Members

Scenario: สมาชิก VIP Gold
Given พัสดุมีน้ำหนัก 5.5 กิโลกรัม
And ระยะทาง 30 กิโลเมตร
And เลือกการจัดส่งแบบ "Standard"
And เป็นสมาชิก "VIP_GOLD"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 72 บาท # Calculation: 80 × 0.9 = 72

Scenario: สมาชิก VIP Platinum - ส่งฟรี
Given พัสดุมีน้ำหนัก 5.5 กิโลกรัม
And ระยะทาง 30 กิโลเมตร
And เลือกการจัดส่งแบบ "Standard"
And เป็นสมาชิก "VIP_PLATINUM"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 0 บาท

Scenario: สมาชิก VIP Platinum - Express ยังต้องจ่ายส่วนต่าง
Given พัสดุมีน้ำหนัก 5.5 กิโลกรัม
And ระยะทาง 30 กิโลเมตร
And เลือกการจัดส่งแบบ "Express"
And เป็นสมาชิก "VIP_PLATINUM"
When ทำการคำนวณค่าจัดส่ง
Then ค่าจัดส่งต้องเท่ากับ 40 บาท # Calculation: Standard 80 = Free, Express 120, Pay difference 40

# Edge Cases

Scenario: น้ำหนักเป็น 0 หรือค่าลบ
Given พัสดุมีน้ำหนัก 0 กิโลกรัม
When ทำการคำนวณค่าจัดส่ง
Then ระบบแสดงข้อผิดพลาด "Weight must be greater than 0"

Scenario: ระยะทางเป็นค่าลบ
Given พัสดุมีน้ำหนัก 1.5 กิโลกรัม
And ระยะทาง -10 กิโลเมตร
When ทำการคำนวณค่าจัดส่ง
Then ระบบแสดงข้อผิดพลาด "Distance cannot be negative"

Scenario: ใช้โค้ดส่วนลดที่ไม่มีอยู่
Given พัสดุมีน้ำหนัก 1.5 กิโลกรัม
And ระยะทาง 30 กิโลเมตร
And ใช้โค้ดส่วนลด "INVALID_CODE"
When ทำการคำนวณค่าจัดส่ง
Then ระบบแสดงข้อผิดพลาด "Invalid promotion code"
