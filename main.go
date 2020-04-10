package main

import (
	"./parsers"
	"fmt"
)

func main() {
	digea := parsers.Ept{}
	digea.Parse()

	fmt.Println(digea.Marshal())
}
