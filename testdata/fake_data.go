package testdata

import (
	"github.com/avani-bit/e-commerce-backend/internal/database"
	"github.com/avani-bit/e-commerce-backend/internal/models"
)

func LoadFakeData() {
	db := database.GetDB()

	// ----- Product -----
	products := []models.Product{
		{
			ID:           "prod_001",
			Brand:        "PUMA",
			BrandTier:    models.BrandTierRegular,
			Category:     "T-shirts",
			BasePrice:    float32(1000),
			CurrentPrice: float32(1000),
		},
		{
			ID:           "prod_002",
			Brand:        "NIKE",
			BrandTier:    models.BrandTierPremium,
			Category:     "Shoes",
			BasePrice:    float32(5000),
			CurrentPrice: float32(5000),
		},
		{
			ID:           "prod_003",
			Brand:        "FASTRACK",
			BrandTier:    models.BrandTierBudget,
			Category:     "Sunglasses",
			BasePrice:    float32(800),
			CurrentPrice: float32(800),
		},
	}
	for _, product := range products {
		db.AddProduct(product)
	}

	// ----- Vouchers -----
	db.Vouchers = map[string]models.Voucher{
		"SUPER69": {
			Code:                 "SUPER69",
			Percent:              float64(69),
			ExcludedBrands:       []string{"Nike"},
			ExcludedCategories:   []string{"Shoes"},
			AllowedCustomerTiers: []models.CustomerTier{models.CustomerTierGold, models.CustomerTierDiamond},
		},
	}

	// ----- Offers -----
	db.Offers["brand_puma"] = models.Offer{
		Name:    "Min 40% off on PUMA",
		Percent: float64(40),
		Target:  "brand",
		Condition: func(p models.Product) bool {
			return p.Brand == "PUMA"
		},
	}

	db.Offers["category_tshirts"] = models.Offer{
		Name:    "Extra 10% off on T-shirts",
		Percent: float64(10),
		Target:  "category",
		Condition: func(p models.Product) bool {
			return p.Category == "T-shirts"
		},
	}

	db.Offers["bank_icici"] = models.Offer{
		Name:    "ICICI Bank 10% instant discount",
		Percent: 10.0,
		Target:  "bank",
		Condition: func(p models.Product) bool {
			// This is generic; actual check will be in payment method logic
			return true
		},
	}
}
