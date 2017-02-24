package main

import "github.com/gordonmleigh/redpill"
import "fmt"
import "os"

func main() {
	c, err := redpill.GetContainerID()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(c)
}
