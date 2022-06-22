package main

import "fmt"

var bufferCount = 0

func AllocateBuffer() *string {
	if bufferCount == 3 {
		return nil
	}

	buffer := new(string)
	bufferCount++
	return buffer
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
