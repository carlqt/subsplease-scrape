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
		fmt.Fprintf(os.Stderr, "Usage: subscrape [command] [arguments]\n\n")
		fmt.Fprintf(os.Stderr, "The commands are:\n")
		fmt.Fprintf(os.Stderr, "  %s\t%s\n", downloadCommand.Name, downloadCommand.Description)
		fmt.Fprintf(os.Stderr, "  %s\t%s\n", organizeCommand.Name, organizeCommand.Description)
		fmt.Fprintf(os.Stderr, "\nUse \"subscrape [command] -h\" for more information about a command.\n")
	}

	if len(os.Args) == 1 || len(os.Args) == 2 && os.Args[1] == "-h" {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case downloadCommand.Name:
		if len(os.Args) < 3 {
			downloadCommand.Usage()
			return
		}

		err := downloadCommand.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}

		downloadCommand.Run()
	case organizeCommand.Name:
		if len(os.Args) < 3 {
			organizeCommand.Usage()
			return
		}

		err := organizeCommand.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		organizeCommand.Run()

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
