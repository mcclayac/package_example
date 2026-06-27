// This is the package for formatting numbers
// called fo-format
//
// if has the method:
//
//	Number(num int) string
package do_format

import "fmt"

// Number create a formatted string for a number
// it takes a num int and returns a string
func Number(num int) string {
	return fmt.Sprintf("The Number is : %d", num)
}
