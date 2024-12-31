package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/carlqt/anime-downloader/organizer"
	"github.com/carlqt/anime-downloader/subsplease"
)

func Run(address *url.URL) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}

	outputDir := filepath.Join(homeDir, "Downloads")

	subsplease, err := subsplease.NewSubsplease(address, outputDir)
	if err != nil {
		panic(err)
	}

	subsplease.Run()
}

func main() {
	// config log with line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Handle command line args
	urlVal := flag.String("u", "", "The url of subsplease to scrape")
	episodeName := flag.String("o", "", "Name of the episode")
	flag.Parse()

	if *urlVal != "" {
		parsedUrl, err := url.ParseRequestURI(*urlVal)
		if err != nil {
			fmt.Printf("%s is not a valid url\n", *urlVal)
			return
		}

		fmt.Println("Start")
		Run(parsedUrl)
		fmt.Println("Completed")
	} else if *episodeName != "" {
		sourceDir := ""

		organizer.Run(*episodeName, sourceDir)
	}
}

// Initializers
// Visit page to get sid
//	- scrape
