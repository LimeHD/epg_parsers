package utils

import "fmt"

func PrintMeta(name string, version string, output string) {
	fmt.Println(fmt.Sprintf("Run parser: %s", name))
	fmt.Println(fmt.Sprintf("Parser version: %s", version))
	fmt.Println(fmt.Sprintf("Output to: %s", output))
}
