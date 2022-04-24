package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

func containNumsInChar(chars string) bool {
	for _, text := range chars {
		if '0' <= text && text <= '9' {
			return false
		}
	}
	return true
}
