// cl_echo prints it's command line arguments.
package main

import (
	"fmt" // string formatting
	"os" // Includes standard OS I/O ops
	"strconv" // Convert strings and numbers
)

func main() {
	var s, sep string
	// Option 1 for for loop - traditional, no parenthesis EVER
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	// reset s
	s = ""

	// Another for loop option
	// 'range' loop - Gives you both the index and value
	// '_' is blank identifier, takes place of index since index
	// is not used but it is required syntactically
	// If we set Args to Args[0:] we would also get the name of the invoker
	for _, arg := range os.Args[1:] {
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)

	s = ""

	// Modification for printing index and value
	for index, arg := range os.Args[1:] {
		s += "Index: " + strconv.Itoa(index) + ", Arg: " + arg + "\n"
	}

	fmt.Println(s)
}
