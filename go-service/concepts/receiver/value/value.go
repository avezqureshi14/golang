// In go we don't have classes but we can attaches function to the types (structs) now this function are now known as methods and the type they are attached to is called the Receiver
// There are two type of receivers value and pointer

package pointer

type User struct{
	name string
}


// now when we use value receiver , Go makes a full copy of the struct Any changes you make inside a method stay inside that method 
func (u User) getUserName() string{
	return u.name
}


func Value() {

}