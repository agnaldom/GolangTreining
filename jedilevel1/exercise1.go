/* 1 - Using the sort declaration operator, ASSIGN these VALUES to VARIAB
IDENTIFIERS 'x' and 'y' an 'z'
	a. 42
	b. "James Bond"
	c. true
2. Now print the values stored in those variables using 
	a. a single print statement
	b. multiple print statement
*/

packege main

import (
	"fmt"	
)

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}