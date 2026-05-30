package main

// Now good thing will be not forcing capabilties into base interfaces


type PaymentMethod interface {
	Pay(amount int) error 
}

type Refundable interface {
	Refund(amount int) error 
}

func main(){

}