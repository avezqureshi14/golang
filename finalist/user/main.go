package main

import "fmt"

type NotifierRepo interface {
	SendEmail(name string, email string) string
	SendSMS(name string, mobile string) string
}

type Notifier struct{}

var _ NotifierRepo = (*Notifier)(nil)

func (n *Notifier) SendEmail(name string, email string) string {
	return fmt.Sprintf("Email sent to %s on email %s ", name, email)
}

func (n *Notifier) SendSMS(name string, mobile string) string {
	return fmt.Sprintf("SMS sent to %s on mobile number %s", name, mobile)
}

type User struct {
	name   string
	email  string
	mobile string
}

type UserService struct {
	user     User
	notifier Notifier
}

func NewUser(user User, notifier Notifier) *UserService {
	return &UserService{
		user:     user,
		notifier: notifier,
	}
}

func (u *UserService) CreateUser() {
	res := u.notifier.SendEmail(u.user.name, u.user.email)
	fmt.Println(res)
}

type Counter struct {
	count int
}

func (c *Counter) IncrePtr() {
	c.count++
	fmt.Println("Coming frm pointer ", c.count)
}

func (c Counter) IncreVal() {
	c.count++
	fmt.Println("Coming frm value", c.count)
}

func main() {
	notifier := Notifier{}
	user := User{
		name:   "avez",
		email:  "avez@webknot.in",
		mobile: "9890562214",
	}
	u := NewUser(user, notifier)
	u.CreateUser()

	c := Counter{}
	c.IncrePtr()
	c.IncreVal()
	c.IncreVal()
	c.IncrePtr()
	c.IncrePtr()
}
