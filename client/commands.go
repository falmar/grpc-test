package main

import (
	"flag"
	"fmt"
)

type ListFlags struct {
	query     string
	limit     int64
	offset    int64
	completed bool
}

type CreateFlags struct {
	name      string
	completed bool
}

type UpdateFlags struct {
	id        string
	completed bool
}

type DeleteFlags struct {
	id        string
	completed bool
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
		set.BoolVar(&lf.completed, "c", false, "retrieve completed TODOs")

		err := set.Parse(args[1:])

		return lf, err

	case "create":
		f := CreateFlags{}

		set := flag.NewFlagSet("create", flag.ExitOnError)

		set.StringVar(&f.name, "name", "", "TODO name")
		set.BoolVar(&f.completed, "c", false, "set TODO as completed")

		err := set.Parse(args[1:])

		if len(f.name) == 0 {
			set.PrintDefaults()

			return nil, fmt.Errorf("body cannot be empty")
		}

		return f, err

	case "mark-completed":
		f := UpdateFlags{}

		set := flag.NewFlagSet("update", flag.ExitOnError)

		set.StringVar(&f.id, "id", "", "TODO id")
		set.BoolVar(&f.completed, "c", false, "set TODO as completed")

		err := set.Parse(args[1:])

		if len(f.id) == 0 {
			set.PrintDefaults()

			return nil, fmt.Errorf("id cannot be empty")
		}

		return f, err

	case "delete":
		f := DeleteFlags{}

		set := flag.NewFlagSet("delete", flag.ExitOnError)

		set.StringVar(&f.id, "id", "", "TODO id")

		err := set.Parse(args[1:])

		if len(f.id) == 0 {
			set.PrintDefaults()

			return nil, fmt.Errorf("id cannot be empty")
		}

		return f, err
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
