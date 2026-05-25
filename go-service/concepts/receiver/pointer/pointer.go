// In go we don't have classes but we can attaches function to the types (structs) now this function are now known as methods and the type they are attached to is called the Receiver
// There are two type of receivers value and pointer

package pointer

type User struct{
	name string
}


// now when we use pointer receiver , the method points to the actual memory address of the struct changes made here persist outside the method
func (u *User) getUserName() string{
	return u.name
}


func Pointer() {

}