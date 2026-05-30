package main

// DIP says high level modules should not depend on low level modules but both should depend on abstractions

// now over here ur service is tightly coupled with mysql db now if tomorrow u want to change to mongodb it will take lots of efforts

// Now in this sceanrio we will use a reposiroty pattern and use than repo in the service have all main low level logic connection to the repo
type UserService struct {
	db MySQL
}

func main() {

}