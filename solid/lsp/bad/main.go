package main

import "errors"

// Subtypes should be replaceable without breaking behaviour , now this below code is LSP violation because u told all PaymentMethods support Refund but some aren't

type PaymentMethod interface {
	Pay(amount int) error
	Refund(amount int) error
}

type UPI struct{}

func (u UPI) Pay(amount int) error {
	return nil
}

func (u UPI) Refund(amount int) error {
	return errors.New("Refund not supported")
}

// now if someone passes UPI it will be runtime failure
func ProcessRefund(p PaymentMethod){
	p.Refund(100) 
}

func main() {

}
