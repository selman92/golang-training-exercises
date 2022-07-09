package main

import "fmt"

func WordGenerator(words []string) func() string {
	idx := 0

	return func() string {
		w := words[idx]

		idx++

		if idx == len(words) {
			idx = 0
		}

		return w
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
