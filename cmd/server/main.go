package main

import (
	"context"
	"github.com/avani-bit/e-commerce-backend/internal/database"
)

func main() {
	ctx := context.Background()
	database.InitDB()

	// Initialize service here
	_ = ctx
}
