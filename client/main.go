package main

import (
	"fmt"
	"github.com/falmar/grpc-test/pb"
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

	case CreateFlags:
		err = Create(client, command.(CreateFlags))

	case UpdateFlags:
		err = Update(client, command.(UpdateFlags))

	case DeleteFlags:
		err = Delete(client, command.(DeleteFlags))
	}

	if err != nil {
		log.Fatal(err)
	}
}
