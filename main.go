// This is the main function of the package
// github.com/mcclayac/package_example
//
// create by: mcclayac
package main

import (
	"fmt"

	do_format "github.com/mcclayac/package_example/do-format"
	"github.com/mcclayac/package_example/math"
)

//http://localhost:8080
//go install golang.org/x/pkgsite/cmd/pkgsite@latest
//pkgsite

func main() {
	fmt.Println("Run main in package: ")
	fmt.Println("module github.com/mcclayac/package_example")

	num := math.Double(2)
	output := do_format.Number(num)
	fmt.Println(output)

}
