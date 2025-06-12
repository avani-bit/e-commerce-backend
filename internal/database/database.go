package database

import (
	"errors"
	"sync"

	"github.com/avani-bit/e-commerce-backend/internal/models"
)

type InMemoryDB struct {
	Products map[string]models.Product
	Vouchers map[string]models.Voucher // VoucherCode -> Discount %
	Offers   map[string]models.Offer   // Brand/Category/Bank -> Offer
	mutex    sync.RWMutex
}

var runtimeDB *InMemoryDB

func InitDB() {
	runtimeDB = &InMemoryDB{
		Products: make(map[string]models.Product),
		Vouchers: make(map[string]models.Voucher),
		Offers:   make(map[string]models.Offer),
	}
}

func GetDB() *InMemoryDB {
	if runtimeDB == nil {
		InitDB()
	}
	return runtimeDB
}

func (db *InMemoryDB) AddProduct(p models.Product) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Products[p.ID] = p
}

func (db *InMemoryDB) GetProduct(id string) (models.Product, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	p, ok := db.Products[id]
	if !ok {
		return models.Product{}, errors.New("product not found")
	}
	return p, nil
}
