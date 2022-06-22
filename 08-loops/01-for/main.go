package main

import "fmt"

func Alphabet(length int) []string {
	chars := make([]string, length)

	for i := 0; i < length; i++ {
		chars[i] = characterByIndex(i)
	}

	return chars
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
