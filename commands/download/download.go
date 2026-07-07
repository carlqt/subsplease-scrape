package download

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"slices"

	"github.com/carlqt/anime-downloader/commands/internal"
)

type DownloadCommand struct {
	// Description is used for global display
	Description string
	// Usage is for displaying more information about the command
	Name     string
	Argument *url.URL
	Flags    *DownloadCommandFlags
	FlagSet  *flag.FlagSet
}

type DownloadCommandFlags struct {
	Resolution int // Should only be 1080, 720 or 480. Default is 720p.
}

func NewDownloadCommand() DownloadCommand {
	commandFlags := &DownloadCommandFlags{}

	name := "download"
	downloadCommand := flag.NewFlagSet(name, flag.ExitOnError)
	downloadCommand.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: subscrape [OPTIONS] download <url>\n\n", internal.Underline("Usage"))
		fmt.Fprintf(os.Stderr, "The url is the url of the subs-please episode or show page.\n\n")
		fmt.Fprintf(os.Stderr, "%s: subscrape download https://subs-please.org/shows/<show-name>/\n\n", internal.Underline("Example"))
		fmt.Fprintf(os.Stderr, "%s\n", internal.Underline("Options:"))
		downloadCommand.PrintDefaults()
	}

	downloadCommand.IntVar(&commandFlags.Resolution, "r", 720, "Resolution to download. Choose between 1080, 720 and 480.")

	return DownloadCommand{
		Description: "Download episodes from subs-please.org. You can provide a URL to a specific episode or a show page.",
		Name:        name,
		Flags:       commandFlags,
		FlagSet:     downloadCommand,
	}
}

func (c DownloadCommand) Usage() {
	c.FlagSet.Usage()
}

// Parse validates and handles the cli arguments and flags
func (c *DownloadCommand) Parse(arguments []string) error {
	err := c.FlagSet.Parse(arguments)
	if err != nil {
		return fmt.Errorf("failed to parse argument for download: %v", err)
	}

	urlArg := c.FlagSet.Arg(0)
	parsedUrl, err := url.ParseRequestURI(urlArg)
	if err != nil {
		return fmt.Errorf("%s is not a valid url: %v", urlArg, err)
	}

	allowedResolutions := []int{1080, 720, 480}
	resolutionFlag := c.Flags.Resolution

	if slices.Contains(allowedResolutions, resolutionFlag) {
		c.Argument = parsedUrl
		return nil
	}

	return fmt.Errorf("Invalid resolution: %d. Choose between 1080, 720 and 480.", resolutionFlag)
}

func (c DownloadCommand) Run() {
	if !c.FlagSet.Parsed() {
		log.Println("FlagSet not parsed. Please call Parse() before Run().")
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return
	}

	outputDir := filepath.Join(homeDir, "Downloads")
	resolution := c.Flags.Resolution

	subsplease, err := newSubsplease(c.Argument, outputDir, resolution)
	if err != nil {
		log.Println(err)
		return
	}

	subsplease.Run()
}
