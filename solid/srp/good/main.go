package main

type UserService struct {
	repo  UserRepository
	email EmailService
}

func (u UserService) CreateUser() {
	u.repo.Save()
	u.email.Send()
}

func main() {

}