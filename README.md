# e-commerce-backend

This project is a modular backend service written in Go that simulates core discounting functionality for a fashion-focused e-commerce platform. The system is built to be extensible, maintainable, and testable from the ground up — with zero external dependencies in the early phase.

---

## In-Memory Data Infrastructure

The application uses a custom-built, runtime in-memory database defined under `internal/database/`. This provides a lightweight, dependency-free environment ideal for prototyping and testing business logic.

---

### `internal/database/db.go`

#### Purpose

Implements a global in-memory singleton database used for storing runtime data such as:

- Product Catalog
- Discount Vouchers
- Promotional Offers

#### Design

- **Singleton Pattern** — accessible via `GetDB()` across the app
- **Thread-safe usage** in single-threaded test/dev environments
- **Resettable** via `InitDB()` to isolate test environments
- No external dependencies (pure Go maps)

#### Data Models Stored

| Type      | Description                         | Key         |
|-----------|-------------------------------------|-------------|
| Product   | All available inventory items       | `product_id`|
| Voucher   | Coupon codes with rules             | `code`      |
| Offer     | Discounts by brand/category/bank    | `name`      |

#### Example Usage

```go
db := database.GetDB()

// Add a product
db.AddProduct(models.Product{ID: "prod_001", Brand: "PUMA", Category: "T-shirts"})

// Get a product
p := db.GetProductByID("prod_001")
```

## `testdata/fake_data.go`

### Purpose

Populates the in-memory DB with seed data representing realistic e-commerce scenarios — primarily used for bootstrapping development environments and validating discounting logic.

---

### Seeded Entities

#### Products

- **PUMA T-shirt** — Base Price: ₹1000  
- **NIKE Shoes** - Base Price: ₹5000
- **FASTRACK Sunglasses** - Base Price: ₹800

---

#### Vouchers

- **SUPER69** — 69% off  
  - *Applies only to* `budget` / `regular` brand tiers  
  - *Excludes* `Nike` & `Shoes` category

---

#### Offers

- **Min 40% off on PUMA** *(Brand-level offer)*  
- **Extra 10% off on T-shirts** *(Category-level offer)*  
- **ICICI Bank 10% instant discount** *(Bank offer)*
