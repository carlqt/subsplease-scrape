package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"golang.org/x/net/html"
)

const host = "https://subsplease.org/shows"

func Run() {
	// Code
	// Arguments from CLI -> type is a URL
	path := "7th-time-loop"
	fullUrl := fmt.Sprintf("%s/%s", host, path)

	resp, err := http.Get(fullUrl)
	if err != nil {
		slog.Error(err.Error())
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		slog.Error("Cannot parse HTML")
		return
	}

	sid := getSID(doc)
	getEpisodes(sid)
}

func main() {
	// config log with line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Start")
	Run()
	fmt.Println("Completed")
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
