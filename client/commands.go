package main

import (
	"flag"
	"fmt"
)

type ListFlags struct {
	query  string
	limit  int64
	offset int64
}

func parseCommand(args []string) (interface{}, error) {
	if len(args) < 1 {
		printBasicUsage()

		return nil, flag.ErrHelp
	}

	switch args[0] {
	case "list":
		lf := ListFlags{}
		set := flag.NewFlagSet("list", flag.ExitOnError)

		set.Int64Var(&lf.limit, "n", 15, "max amount of items")
		set.Int64Var(&lf.offset, "o", 0, "offset of items")
		set.StringVar(&lf.query, "q", "", "search query string")

		err := set.Parse(args[1:])

		return lf, err
	default:
		printBasicUsage()

		return nil, flag.ErrHelp
	}
}

func printBasicUsage() {
	fmt.Printf("grpc-test CLI\n\n")
	fmt.Printf("usage: command [arguments]\n\n")
	fmt.Println("Commands:")
	fmt.Printf("\tlist\tshow existing todos\n")
}
