package generics

import "fmt"

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func sumNumber[T Number](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func Generics() {
	numbers := []float32{1.0, 0.2, 3.0, 4.0, 5.0}
	fmt.Println("Sum of numbers ", sumNumber(numbers))
}
