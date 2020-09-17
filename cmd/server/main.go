package main

import (
	"fmt"
	"os"

	petstoreServer "github.com/f-go/petstore/pkg/server"
)

func main() {
	if err := petstoreServer.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

