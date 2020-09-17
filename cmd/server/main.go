package main

import (
	"fmt"
	"os"

	gha "github.com/f-go/ratehub/pkg/server"
)

func main() {
	if err := gha.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

