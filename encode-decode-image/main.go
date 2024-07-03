package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	// Read the entire file into a byte slice
	file, err := ioutil.ReadFile("./banner.jpg")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(file)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(file)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)

	emptyStr := ""

	decoded, err := base64.StdEncoding.DecodeString(emptyStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=================")
	fmt.Println("decoded empty str", decoded)

	reader := bytes.NewReader(decoded)
	fmt.Println(reader)
}
