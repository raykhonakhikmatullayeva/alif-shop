package product

import "fmt"

func ExampleMonthlyPayment() {
	p := Product{
		ID:     7921,
		Name:   "Xiaomi Redmi Note 14",
		Price:  2499000,
		Months: 12,
	}
	fmt.Println(MonthlyPayment(p))
	// Output: 208250
}
