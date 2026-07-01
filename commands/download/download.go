package download

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

type DownloadCommand struct {
	// Description is used for global display
	Description string
	// Usage is for displaying more information about the command
	Name    string
	FlagSet *flag.FlagSet
}

type downloadCommandArgs struct {
	Url *url.URL
}

func newDownloadCommandArgs(flagSetArgs []string) (downloadCommandArgs, error) {
	urlArg := flagSetArgs[0]

	parsedUrl, err := url.ParseRequestURI(urlArg)
	if err != nil {
		return downloadCommandArgs{}, fmt.Errorf("%s is not a valid url: %v", urlArg, err)
	}

	return downloadCommandArgs{Url: parsedUrl}, nil
}

func NewDownloadCommand() DownloadCommand {
	name := "download"
	downloadCommand := flag.NewFlagSet(name, flag.ExitOnError)
	downloadCommand.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: subscrape download <url>\n\n")
		fmt.Fprintf(os.Stderr, "The url is the url of the subs-please episode or show page.\n\n")
		fmt.Fprintf(os.Stderr, "Example: subscrape download https://subs-please.org/shows/<show-name>/\n")
	}

	return DownloadCommand{
		Description: "Download episodes from subs-please.org. You can provide a URL to a specific episode or a show page.",
		Name:        name,
		FlagSet:     downloadCommand,
	}
}

func (c DownloadCommand) Usage() {
	c.FlagSet.Usage()
}

func (c DownloadCommand) Run() {
	c.FlagSet.Parse(os.Args[2:])

	args, err := newDownloadCommandArgs(c.FlagSet.Args())
	if err != nil {
		log.Println(err)
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return
	}

	outputDir := filepath.Join(homeDir, "Downloads")

	subsplease, err := newSubsplease(args.Url, outputDir)
	if err != nil {
		log.Println(err)
		return
	}

	subsplease.Run()
}
