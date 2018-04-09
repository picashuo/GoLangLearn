//指標 (感覺類似C# out)過程
package main

import (
	"fmt"
)

func zero(xPtr *int, tests *int) {
	*xPtr = 0
	*tests = 100
}
func main() {
	x := 5
	y := 5
	zero(&x, &y)
	fmt.Println(x, y)
}
