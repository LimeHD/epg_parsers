package utils

import (
	"io"
	"log"
	"net/http"
)

func GetHtmlDocumentReader(url string) (int, io.ReadCloser) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	return res.StatusCode, res.Body
}
