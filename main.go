package main

import (
	"fmt"

	do_format "github.com/mcclayac/package_example/do-format"
	"github.com/mcclayac/package_example/math"
)

func main() {
	fmt.Println("Run main in package: ")
	fmt.Println("module github.com/mcclayac/package_example")

	num := math.Double(2)
	output := do_format.Number(num)
	fmt.Println(output)

}
