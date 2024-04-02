package subsplease

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/carlqt/anime-downloader/domquery"
	"golang.org/x/net/html"
)

type Subsplease struct {
	ShowAddress     string
	OutputDirectory string
}

func NewSubsplease(address *url.URL, outputDir string) (Subsplease, error) {
	var subsplease Subsplease

	if address.Host != "subsplease.org" {
		return subsplease, errors.New("host needs to be subsplease.org")
	}

	subsplease = Subsplease{
		ShowAddress:     address.String(),
		OutputDirectory: outputDir,
	}

	return subsplease, nil
}

func (s Subsplease) Run() {
	// Scrape page to get the sid
	htmlDoc, _ := getHtmlDocument(s.ShowAddress)

	sid := getSID(htmlDoc)

	// Call the API
	getEpisodes(sid)
}

func getHtmlDocument(address string) (*html.Node, error) {
	// Parse page
	resp, err := http.Get(address)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("request page unsuccessful")
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		slog.Error("Cannot parse HTML")
		return nil, errors.New("cannot parse html")
	}

	return doc, nil
}

func getSID(n *html.Node) string {
	var sid string
	node := domquery.GetElementById(n, "show-release-table")

	sid, ok := domquery.GetAttribute(node, "sid")

	if !ok {
		log.Println("Cannot find sid")
	}

	return sid
}
