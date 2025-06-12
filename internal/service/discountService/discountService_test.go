package discountService

import (
	"context"
	"github.com/avani-bit/e-commerce-backend/internal/database"
	"github.com/avani-bit/e-commerce-backend/internal/models"
	"github.com/avani-bit/e-commerce-backend/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateCartDiscounts(t *testing.T) {
	database.InitDB()
	testdata.LoadFakeData()
	svc := NewDiscountService()

	ctx := context.Background()

	cart := []models.CartItem{
		{
			Product:  database.GetDB().Products["prod_001"],
			Quantity: 1,
			Size:     "M",
		},
	}

	bank := "ICICI"
	payment := &models.PaymentInfo{
		Method:   "CARD",
		BankName: &bank,
	}

	customer := models.CustomerProfile{
		ID:   "cust_123",
		Tier: "gold",
	}

	result, err := svc.CalculateCartDiscounts(ctx, cart, customer, payment)
	assert.NoError(t, err)

	assert.Equal(t, float32(1000), result.OriginalPrice, "Expected original price to be 1000, got %v", result.OriginalPrice)

	assert.True(t, result.FinalPrice < result.OriginalPrice,
		"Expected final price > original price, got final: %v, original: %v",
		result.FinalPrice, result.OriginalPrice)

	assert.Greater(t, len(result.AppliedDiscounts), 0,
		"Expected at least one discount to be applied, got %v", result.AppliedDiscounts)

}

func TestValidateDiscountCode(t *testing.T) {
	// Add table-driven tests here
}
