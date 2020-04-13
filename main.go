package main

import (
	"fmt"
	"github.com/LimeHD/parser/parsers"
)

func main() {
	digea := parsers.Digea{}
	digea.Parse()

	for _, v := range digea.ToCSV() {
		fmt.Println(v)
	}
}
