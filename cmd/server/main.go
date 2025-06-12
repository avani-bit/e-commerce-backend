package main

import (
	"context"
	"github.com/avani-bit/e-commerce-backend/internal/database"
	testdata "github.com/avani-bit/e-commerce-backend/testData"
)

func main() {
	ctx := context.Background()
	database.InitDB()
	testdata.LoadFakeData()

	// Initialize service here
	_ = ctx
}
