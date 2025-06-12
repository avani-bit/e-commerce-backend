package discountService

import (
	"context"
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
