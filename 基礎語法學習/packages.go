//引用pacgeage初學
package main

import (
	"fmt"

	"./src/HelloWorld"
)

func main() {
	var test1 string
	test1 = HelloWorld.HelloWorlds()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", HelloWorld.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
}
