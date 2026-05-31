package main

type DiscountStrategy interface {
	GetDiscount() int
}

type Regular struct{}

func (r Regular) GetDiscount() int { return 10 }

type Premium struct{}

func (r Premium) GetDiscount() int { return 20 }

// now further using factory pattern at runtime u can get the required one

func main() {

}
