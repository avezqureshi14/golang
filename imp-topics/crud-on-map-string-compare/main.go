package main

import (
	"fmt"
	"sync"
)

//
// =======================
// 1. GLOBAL CONFIG MAP
// =======================
//

var Config = map[string]string{
	"env":  "dev",
	"port": "8080",
}

//
// =======================
// 2. USER REGISTRY (CRUD)
// =======================
//

type User struct {
	Name string
	Age  int
}

var (
	users = make(map[string]User)
	mu    sync.RWMutex
)

// CREATE
func CreateUser(id string, user User) {
	mu.Lock()
	defer mu.Unlock()
	users[id] = user
}

// READ
func GetUser(id string) (User, bool) {
	mu.RLock()
	defer mu.RUnlock()
	u, ok := users[id]
	return u, ok
}

// UPDATE
func UpdateUser(id string, user User) {
	mu.Lock()
	defer mu.Unlock()
	users[id] = user
}

// DELETE
func DeleteUser(id string) {
	mu.Lock()
	defer mu.Unlock()
	delete(users, id)
}

//
// =======================
// 3. STRING TYPE DEMO
// =======================
//

func stringDemo() {
	s1 := "hello\nworld" // double quotes → escape sequences work
	s2 := `hello\nworld` // backticks → raw string, no escape processing
	var r rune = 'a'     // single quotes → rune (Unicode code point)

	fmt.Println("double quotes:", s1)
	fmt.Println("backticks:", s2)
	fmt.Println("rune:", r)
}

func main() {

	// 1. Config map
	fmt.Println("Config:", Config)

	// 2. CRUD operations
	CreateUser("1", User{Name: "Avez", Age: 23})
	CreateUser("2", User{Name: "John", Age: 30})

	fmt.Println("Get User 1:", func() interface{} {
		u, _ := GetUser("1")
		return u
	}())

	UpdateUser("1", User{Name: "Avez Updated", Age: 24})

	fmt.Println("Updated User 1:", func() interface{} {
		u, _ := GetUser("1")
		return u
	}())

	DeleteUser("2")
	fmt.Println("After Delete User 2:", users)

	// 3. String demo
	stringDemo()

}
