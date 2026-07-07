package internal

import "flag"

type BaseCommand struct {
	Description string
	Name        string
	FlagSet     *flag.FlagSet
}
