package main

// Don’t force clients (or implementers) to depend on methods they don’t need.
type Worker interface {
	Work()
	Eat()
}

func main() {

}
