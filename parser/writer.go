package base

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func (d *Day) ToTSV() []string {
	csv := []string{}

	// надо как-то разрулить этот O(N^2) на более высоком уровне дойдет и до куба...
	for _, c := range d.Common.Channels {
		for _, p := range c.Programs {
			csv = append(csv, fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s",
				strings.TrimSpace(p.Timestart),
				strings.TrimSpace(p.Timestop),
				strings.TrimSpace(c.Name),
				strings.TrimSpace(p.Title),
				strings.TrimSpace(c.Icon),
				p.Description,
			))
		}
	}

	return csv
}

func WriteTSV(output string, days map[string]*Day) {
	file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)

	WriteHeader(writer)
	for _, day := range days {
		for _, data := range day.ToTSV() {
			_, _ = writer.WriteString(data + "\n")
		}
	}

	writer.Flush()
}

func WriteHeader(w *bufio.Writer) {
	_, _ = w.WriteString("datetime_start\tdatetime_finish\tchannel\ttitle\tchannel_logo_url\tdescription\n")
}
