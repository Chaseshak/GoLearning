// cl_echo prints it's command line arguments.
package main

import (
	"fmt" // string formatting
	"os" // Includes standard OS I/O ops
)

func main() {
	var s, sep string
	// Option 1 for for loop - traditional, no parenthesis EVER
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
