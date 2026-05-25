package main

import "fmt"

func safeFunction() {

	// 1. Defer is registered here.
	// This function will execute when safeFunction is about to exit
	// (either normally or due to panic).
	defer func() {

		// 4. During panic, deferred functions are executed (LIFO order).
		// recover() only works inside a deferred function.
		if r := recover(); r != nil {

			// 5. If a panic occurred, recover() captures its value.
			// It stops further panic propagation (stack unwinding).
			fmt.Println("Recovered from panic:", r)
		}

		// After this defer finishes, control returns normally to caller.
	}()

	// 2. Normal execution starts
	fmt.Println("Before panic")

	// 3. Panic is triggered here.
	// - Execution stops immediately
	// - Stack unwinding begins
	// - Deferred functions start executing
	panic("something went wrong")

	// ❌ This line is never executed because panic interrupts flow
	fmt.Println("After panic")
}

func main() {

	// safeFunction is called
	safeFunction()

	// 6. Since panic was recovered inside safeFunction,
	// the program does NOT crash and continues execution normally.
	fmt.Println("Program continues...")
}