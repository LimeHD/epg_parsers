package main

import (
	"fmt"
	"github.com/LimeHD/parser/parsers"
)

func main() {
	digea := parsers.Ept{}
	digea.Parse()

	fmt.Println(digea.Marshal())
}
