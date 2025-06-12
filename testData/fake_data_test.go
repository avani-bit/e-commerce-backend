package testdata

import (
	"testing"

	"github.com/avani-bit/e-commerce-backend/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestLoadFakeData(t *testing.T) {
	// Init DB and load
	database.InitDB()
	LoadFakeData()
	db := database.GetDB()

	// Products
	assert.Len(t, db.Products, 3, "Expected 3 products to be loaded into DB")
	assert.NotNil(t, db.Products["prod_001"], "Expected product with ID prod_001 to exist")
	assert.Equal(t, "PUMA", db.Products["prod_001"].Brand, "Expected brand of prod_001 to be PUMA")

	// Vouchers
	assert.Len(t, db.Vouchers, 1, "Expected 1 voucher to be loaded")
	voucher, ok := db.Vouchers["SUPER69"]
	assert.True(t, ok, "Expected voucher SUPER69 to exist")
	assert.Equal(t, float64(69), voucher.Percent, "Expected SUPER69 voucher to offer 69%")

	// Offers
	assert.Len(t, db.Offers, 3, "Expected 3 offers to be loaded")
	offer, ok := db.Offers["brand_puma"]
	assert.True(t, ok, "Expected brand_puma offer to exist")
	assert.Equal(t, "Min 40% off on PUMA", offer.Name, "Offer name mismatch for brand_puma")
}
