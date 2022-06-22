package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(values ...int) int {
	sum := 0
	for _, v := range values {
		sum = sum + v
	}

	return sum
}
