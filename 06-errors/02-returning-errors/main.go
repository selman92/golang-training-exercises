package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(op1, op2 float64) (float64, error) {
	if op2 == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return op1 / op2, nil
}
