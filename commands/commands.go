package commands

import (
	"github.com/carlqt/anime-downloader/commands/download"
	"github.com/carlqt/anime-downloader/commands/organize"
)

type DownloadCommand = download.DownloadCommand
type OrganizeCommand = organize.OrganizeCommand

func NewDownloadCommand() DownloadCommand {
	return download.NewDownloadCommand()
}

func NewOrganizeCommand() OrganizeCommand {
	return organize.NewOrganizeCommand()
}
