package product

import "fmt"

type ProductID int
type Price int
type Months int
type Product struct {
	ID     ProductID
	Name   string
	Price  Price
	Months Months
}

func MonthlyPayment(p Product) int {
	if p.Months == 0 || p.Months > 12 {
		return 0
	}
	return int(p.Price) / int(p.Months)
}
func PrintProduct(p Product) { // Alif Shop product card
	fmt.Println("====== Alif Shop ==========")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Товар: %s\n", p.Name)
	fmt.Printf("Цена: %d сум\n", p.Price)
	fmt.Printf("Рассрочка: %d месяцев\n", p.Months)
	fmt.Printf("В месяц: %d сум\n", MonthlyPayment(p))
	fmt.Println("===========================")
}
