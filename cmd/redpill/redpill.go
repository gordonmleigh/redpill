package main

import (
	"fmt"
	"os"

	"github.com/gordonmleigh/redpill"
)

func main() {
	c, err := redpill.GetContainerID()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(c)
}
