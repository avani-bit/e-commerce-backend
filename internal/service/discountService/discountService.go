package discountService

import (
	"context"
	"errors"
	"github.com/avani-bit/e-commerce-backend/internal/models"
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
	_ []models.CartItem,
	_ models.CustomerProfile,
	_ *models.PaymentInfo,
) (*models.DiscountedPrice, error) {
	// TODO: Implement CalculateCartDiscounts
	return nil, errors.New("unimplemented")
}

func (d *DiscountServiceImpl) ValidateDiscountCode(
	_ context.Context,
	_ string,
	_ []models.CartItem,
	_ models.CustomerProfile,
) (bool, error) {
	// TODO: Implement ValidateDiscountCode
	return false, errors.New("unimplemented")
}
