// In Go, interface satisfaction depends on method receivers: pointer receiver methods are only available to pointer types, while value receiver methods are available to both value and pointer types.

package polymorphism

import "fmt"

// Step 1: Define interface
type Payment interface {
	Pay(amount float64)
}

// Step 2: Implement for CreditCard
type CreditCard struct {
	Name string
}

func (c *CreditCard) Pay(amount float64) {
	fmt.Println("Paid", amount, "using Credit Card by", c.Name)
}

// Step 3: Implement for UPI
type UPI struct {
	ID string
}

func (u UPI) Pay(amount float64) {
	fmt.Println("Paid", amount, "using UPI with ID", u.ID)
}

// Step 4: Polymorphic function
func MakePayment(p Payment, amount float64) {
	p.Pay(amount)
}

func Polymorphism() {
	// cc := CreditCard{Name: "Avez"}
	pcc := &CreditCard{Name: "Avez"}
	upi := UPI{ID: "avez@upi"}

	// MakePayment(cc, 1000) this gives an error
	MakePayment(pcc, 1000)
	MakePayment(upi, 500)
}