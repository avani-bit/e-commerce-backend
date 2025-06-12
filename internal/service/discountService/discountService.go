package discountService

import (
	"context"
	"errors"
	"fmt"
	"github.com/avani-bit/e-commerce-backend/internal/database"
	"github.com/avani-bit/e-commerce-backend/internal/models"
	"github.com/avani-bit/e-commerce-backend/internal/repository/discountRepository"
	"slices"
)

type DiscountService interface {
	CalculateCartDiscounts(ctx context.Context, cartItems []models.CartItem,
		customer models.CustomerProfile, paymentInfo *models.PaymentInfo) (*models.DiscountedPrice, error)

	ValidateDiscountCode(ctx context.Context, code string, cartItems []models.CartItem,
		customer models.CustomerProfile) (bool, error)
}

// DiscountServiceImpl is a placeholder struct for actual logic
type DiscountServiceImpl struct{}

func NewDiscountService() DiscountService {
	return &DiscountServiceImpl{}
}

func (d *DiscountServiceImpl) CalculateCartDiscounts(
	_ context.Context,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
	paymentInfo *models.PaymentInfo,
) (*models.DiscountedPrice, error) {
	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	db := database.GetDB()

	var originalPrice float32
	var finalPrice float32
	appliedDiscounts := make(map[string]float32)

	// Step 1: Brand & Category Discounts
	for _, item := range cartItems {
		product := item.Product
		totalQtyPrice := product.BasePrice * float32(item.Quantity)

		discountedPrice, discounts := discountRepository.ApplyProductLevelDiscounts(product, item.Quantity, db.Offers)

		originalPrice += totalQtyPrice
		finalPrice += discountedPrice

		for name, amount := range discounts {
			appliedDiscounts[name] += amount
		}
	}

	// Step 2: Voucher
	if voucherDiscount, ok := discountRepository.ApplyVoucherDiscount("SUPER69", finalPrice, db.Vouchers); ok {
		appliedDiscounts["Voucher: SUPER69"] = voucherDiscount
		finalPrice -= voucherDiscount
	}

	// Step 3: Bank Offer
	if paymentInfo != nil && paymentInfo.BankName != nil && *paymentInfo.BankName == "ICICI" {
		bankDiscount := discountRepository.ApplyBankDiscount(*paymentInfo.BankName, finalPrice, db.Offers)
		appliedDiscounts["Bank: ICICI"] = bankDiscount
		finalPrice -= bankDiscount
	}

	return &models.DiscountedPrice{
		OriginalPrice:    originalPrice,
		FinalPrice:       finalPrice,
		AppliedDiscounts: appliedDiscounts,
		Message:          "Discounts applied successfully",
	}, nil
}

func (d *DiscountServiceImpl) ValidateDiscountCode(
	_ context.Context,
	code string,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
) (bool, error) {
	db := database.GetDB()
	voucher, ok := db.Vouchers[code]
	if !ok {
		return false, fmt.Errorf("voucher %s not found", code)
	}

	if len(voucher.AllowedCustomerTiers) > 0 && !slices.Contains(voucher.AllowedCustomerTiers, customer.Tier) {
		return false, fmt.Errorf("voucher is not valid for customer tier: %v", customer.Tier)
	}

	// 2. Check brand/category exclusions
	for _, item := range cartItems {
		brand := item.Product.Brand
		category := item.Product.Category

		if slices.Contains(voucher.ExcludedBrands, brand) {
			return false, fmt.Errorf("voucher not valid for brand: %s", brand)
		}
		if slices.Contains(voucher.ExcludedCategories, category) {
			return false, fmt.Errorf("voucher not valid for category: %s", category)
		}
	}

	return true, nil
}
