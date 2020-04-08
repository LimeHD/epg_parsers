package main

import (
	"./parsers"
	"fmt"
)

func main() {
	digea := parsers.Digea{}
	digea.Parse()

	fmt.Println(digea.Marshal())
}
