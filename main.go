package main

import (
	"fmt"
	"log"

	"github.com/n704/go_parking/lib"
)

func main() {
	command, err := lib.GetCommandObject()
	if err != nil {
		log.Fatal(err)
	}
	outputFunc := command.FetchCommand()
	for output, ok := outputFunc(); ok; output, ok = outputFunc() {
		fmt.Println(output)
	}
}
