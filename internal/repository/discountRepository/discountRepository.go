package discountRepository

import (
	"github.com/avani-bit/e-commerce-backend/internal/models"
)

func ApplyProductLevelDiscounts(p models.Product, qty int, offers map[string]models.Offer) (float32, map[string]float32) {
	baseTotal := p.BasePrice * float32(qty)
	current := baseTotal
	applied := map[string]float32{}

	for _, offer := range offers {
		if (offer.Target == "brand" && offer.Condition(p)) || (offer.Target == "category" && offer.Condition(p)) {
			discount := baseTotal * float32(offer.Percent) / 100
			current -= discount
			applied[offer.Name] = discount
		}
	}
	return current, applied
}

func ApplyVoucherDiscount(code string, price float32, vouchers map[string]models.Voucher) (float32, bool) {
	voucher, ok := vouchers[code]
	if !ok {
		return 0, false
	}

	discount := price * float32(voucher.Percent) / 100
	return discount, true
}

func ApplyBankDiscount(bank string, price float32, offers map[string]models.Offer) float32 {
	offer, ok := offers["bank_icici"]
	if !ok || offer.Target != "bank" {
		return 0
	}
	return price * float32(offer.Percent) / 100
}
