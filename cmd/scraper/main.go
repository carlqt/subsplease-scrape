package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/carlqt/anime-downloader/commands"
)

func main() {
	// config log with line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	downloadCommand := commands.NewDownloadCommand()
	organizeCommand := commands.NewOrganizeCommand()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: subscrape <command> [arguments]\n\n")
		fmt.Fprintf(os.Stderr, "The commands are:\n")
		fmt.Fprintf(os.Stderr, "  download       Downloads episodes in subs-please\n")
		fmt.Fprintf(os.Stderr, "  organize       Organizes the downloaded episodes into a folder\n\n")
		fmt.Fprintf(os.Stderr, "Use \"subscrape <command> -h\" for more information about a command.\n")
	}

	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "download":
		if len(os.Args) < 3 {
			downloadCommand.Usage()
			return
		}

		downloadCommand.Run()
	case "organize":
		if len(os.Args) < 3 {
			organizeCommand.Usage()
			return
		}
		organizeCommand.Run()

	default:
		fmt.Println("Nothing")
	}
}
