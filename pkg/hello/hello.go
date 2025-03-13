package hello

import (
	"fmt"
)

func Hello(i int) {
	// Print the message
	switch i {
	case 1:
		fmt.Println("Hello, World!")
	case 2:
		fmt.Println("Hello, Universe!")
	default:
		fmt.Println("Hello, There!")
	}
}
