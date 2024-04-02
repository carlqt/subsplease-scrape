package subsplease

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"sync"
	"time"
)

type Episode map[string]EpisodeDetail

type Page struct {
	Batch   interface{}
	Episode Episode `json:"episode"`
}

type EpisodeDetail struct {
	Show      string     `json:"show"`
	Episode   string     `json:"episode"`
	Downloads []Download `json:"downloads"`
}

type Download struct {
	Res     string `json:"res"`
	Torrent string `json:"torrent"`
	Magnet  string `json:"magnet"`
	Xdccs   string `json:"xdccs"`
}

// "https://subsplease.org/api/?f=show&tz=Australia/Sydney&sid=701"

func getEpisodes(sid string) {
	var page Page

	episodesUrl := fmt.Sprintf("https://subsplease.org/api/?f=show&tz=Australia/Sydney&sid=%s", sid)

	// Get request to the API
	resp, err := http.Get(episodesUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytesVal, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(bytesVal, &page)

	downloadAllEpisodes(page.Episode)
}

// downloadAllEpisodes extracts the 720p links and downloads it
func downloadAllEpisodes(episode Episode) {
	var wg sync.WaitGroup

	for _, v := range episode {
		for _, d := range v.Downloads {
			if d.Res == "720" {
				wg.Add(1)

				// Throttle execution by 1 second because the API doesn't return Content-Disposition if it's too fast
				time.Sleep(1 * time.Second)
				go func() {
					defer wg.Done()
					downloadTorrent(d.Torrent)
				}()
				break
			}
		}
	}

	wg.Wait()
}

func downloadTorrent(url string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	val := resp.Header.Get("Content-Disposition")
	filename := getFilename(val)

	filepath := fmt.Sprintf("%s/Downloads/%s", homeDir, filename)
	out, _ := os.Create(filepath)
	defer out.Close()

	io.Copy(out, resp.Body)
}

func getFilename(contentDisposition string) string {
	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		log.Println("Invalid media type:", contentDisposition)
	}

	return params["filename"]
}
