package main

import (
	"context"
	"github.com/avani-bit/e-commerce-backend/internal/database"
	"github.com/avani-bit/e-commerce-backend/testdata"
)

func main() {
	ctx := context.Background()
	database.InitDB()
	testdata.LoadFakeData()

	// Initialize service here
	_ = ctx
}
