package models

type BrandTier string

const (
	BrandTierPremium BrandTier = "premium"
	BrandTierRegular BrandTier = "regular"
	BrandTierBudget  BrandTier = "budget"
)

type Product struct {
	ID           string    `json:"id"`
	Brand        string    `json:"brand"`
	BrandTier    BrandTier `json:"brand_tier"`
	Category     string    `json:"category"`
	BasePrice    float32   `json:"base_price"`
	CurrentPrice float32   `json:"current_price"`
}

type Offer struct {
	Name      string
	Percent   float64
	Target    string // brand/category/bank
	Condition func(Product) bool
}

type Voucher struct {
	Code               string
	Percent            float64
	ExcludedBrands     []string
	ExcludedCategories []string
	AllowedTiers       []BrandTier
}

type CartItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	Size     string  `json:"size"`
}

type PaymentInfo struct {
	Method   string  `json:"method"`    // CARD, UPI, etc
	BankName *string `json:"bank_name"` // Optional
	CardType *string `json:"card_type"` // Optional: CREDIT, DEBIT
}

type DiscountedPrice struct {
	OriginalPrice    float32            `json:"original_price"`
	FinalPrice       float32            `json:"final_price"`
	AppliedDiscounts map[string]float32 `json:"applied_discounts"`
	Message          string             `json:"message"`
}

type CustomerProfile struct {
	ID   string `json:"id"`
	Tier string `json:"tier"`
}
