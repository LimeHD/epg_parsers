package utils

import (
	"bufio"
	"log"
	"os"
)

func WriteCSV(output string, data []string) {
	file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)

	WriteHeader(writer)
	for _, data := range data {
		_, _ = writer.WriteString(data + "\n")
	}

	writer.Flush()
}

func WriteHeader(w *bufio.Writer) {
	_, _ = w.WriteString("datetime_start;datetime_finish;channel;broardcast;channel_logo_url;description\n")
}
