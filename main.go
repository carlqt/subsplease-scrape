package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
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

	// bytes, _ := io.ReadAll(resp.Body)
	// fmt.Println("HTML:\n\n", string(bytes))

	// sid := func()

	// Parse and get table element with id "show-release-table"
	// Get the value of the attribute "sid"
	doc, err := html.Parse(resp.Body)
	if err != nil {
		slog.Println("Cannot parse HTML")
		return
	}

	sid := getSID(doc)

	fmt.Println("sid: ", sid)

	// Then make another request to https://subsplease.org/api/?f=show&tz=Australia/Sydney&sid=%s
}

func getSID(n *html.Node) string {
	var sid string
	node := getElementById(n, "show-release-table")

	sid, ok := GetAttribute(node, "sid")
	if !ok {
		log.Println("Cannot find sid")
	}

	return sid
}
