package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const host = "https://subsplease.org/shows"

func main() {
	// config slog with line number
	slog := log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	// Code
	// Arguments from CLI -> type is a URL
	path := "7th-time-loop"
	fullUrl := fmt.Sprintf("%s/%s", host, path)

	resp, err := http.Get(fullUrl)
	if err != nil {
		slog.Println(err)
	}

	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)
	fmt.Println("HTML:\n\n", string(bytes))

	// Parse and get table element with id "show-release-table"
	// Get the value of the attribute "sid"

	// Then make another request to https://subsplease.org/api/?f=show&tz=Australia/Sydney&sid=%s
}
