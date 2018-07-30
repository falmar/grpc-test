package main

import (
	"bitbucket.org/falmar/grpc-test/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	con, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewTODOClient(con)

	command, err := parseCommand(os.Args[1:])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch command.(type) {
	case ListFlags:
		err = List(client, command.(ListFlags))
	}

	if err != nil {
		log.Fatal(err)
	}
}
