package main

import "alif-shop/internal/product"

func main() {
	p := product.Product{
		ID:     7291,
		Name:   "Xiaomi Redmi Note 14",
		Price:  2499000,
		Months: 12,
	}
	product.PrintProduct(p)
}
