package main

import "fmt"

// Interface
type NotifierRepo interface {
	SendEmail(email string) string
	SendSMS(mobile string) string
}

// Concrete implementation
type Notifier struct {
	count int
}

// Compile-time check
var _ NotifierRepo = (*Notifier)(nil)

// Method receivers
func (n *Notifier) SendEmail(email string) string {
	return fmt.Sprintf("Email sent to %s", email)
}

func (n *Notifier) SendSMS(mobile string) string {
	return fmt.Sprintf("SMS sent to mobile number %s", mobile)
}

// Structs
type User struct {
	name   string
	email  string
	mobile string
}

// Service using composition (interface, not concrete type)
type UserService struct {
	user     User
	notifier NotifierRepo
}

// Proper constructor (dependency injection)
func NewUser(user User, notifier NotifierRepo) *UserService {
	return &UserService{
		user:     user,
		notifier: notifier,
	}
}

// Business logic
func (u *UserService) OnBoard() {
	fmt.Println(u.notifier.SendEmail(u.user.email))
	fmt.Println(u.notifier.SendSMS(u.user.mobile))
}

type Count struct {
	count int
}

func (c Count) IncreValueR() {
	c.count++
	fmt.Println("The value of count is VR ", c.count)
}
func (c *Count) IncrePinterR() {
	c.count++
	fmt.Println("The value of count is PR ", c.count)
}

func main() {
	notifier := &Notifier{}

	user := User{
		name:   "avez",
		email:  "avez@webknot.in",
		mobile: "9890562214",
	}

	service := NewUser(user, notifier)
	service.OnBoard()

	// cnt := Count{} : this also works even if not specifically using pointer because go silently converts behind the scenes
	(&Count{}).IncrePinterR()
	Count{}.IncreValueR()
	(&Count{}).IncrePinterR()
	(&Count{}).IncrePinterR()
}
