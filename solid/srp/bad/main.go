package main

// A struct type should have only one reason to change

type UserService struct{}

// so over here everything relation to business logic + email + db all are mixed any changes everyone get affected

func (u UserService) CreateUser() {}
func (u UserService) SendEmail()  {}
func (u UserService) SaveToDB()   {}

func main() {

}
