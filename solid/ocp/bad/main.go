package main

// OCP says code should be open for extension but closed for modification

// now in this case for every new type u will be coming and modifying this function

func GetDiscount(userType string) int {
	switch userType {
	case "regular":
		return 10
	case "premium":
		return 20
	}
	return 0
}

func main() {

}
