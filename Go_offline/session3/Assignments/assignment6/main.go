// Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent.

package main
import (
	"fmt"
)


func main() {
	var hex string
	fmt.Println("Enter a hexadecimal number: ")
	fmt.Scanln(&hex)
	i := hexToInt(hex)
	fmt.Println(i)
}

func hexToInt(hex string) int {
	var i int
	fmt.Sscanf(hex, "%x", &i)
	return i
}
