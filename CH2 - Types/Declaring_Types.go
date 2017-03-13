package main

import(
	"fmt"
)

func main() {
	// Simple math
	fmt.Println("1 + 1 = ", 1 + 1);

	// Simple string
	var hello string = "Hello, World!";

	// Str length
	fmt.Println("Length of hello var: ", len(hello));

	// Gives byte representation since strings are characters which are bytes
	fmt.Println("Char at index 3 (hello[3]): ", hello[3]);
}
