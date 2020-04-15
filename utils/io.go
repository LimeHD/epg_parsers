package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WriteCSV(dir string, name string, data []string) {
	file, err := os.OpenFile(fmt.Sprintf("%s/%s.csv", dir, name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	for _, data := range data {
		_, _ = writer.WriteString(data + "\n")
	}

	writer.Flush()
}
