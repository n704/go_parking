package lib

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type OutputFunc func() (string, bool)

type Fetcher interface {
	FetchCommand() OutputFunc
}

type CommandLine struct {
}

func (p CommandLine) FetchCommand() OutputFunc {
	fileName := flag.Args()[0]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return func() (string, bool) {
		for scanner.Scan() {
			command := scanner.Text()
			return command, true
		}
		file.Close()
		return "", false
	}
}

type InterActiveMode struct {
}

func (p InterActiveMode) FetchCommand() OutputFunc {
	var command string
	return func() (string, bool) {
		fmt.Scanf("%s\n", &command)
		return command, true
	}
}

func GetCommandObject() (Fetcher, error) {
	flag.Parse()
	args := flag.Args()
	switch len(args) {
	case 1:
		return CommandLine{}, nil
	case 0:
		return InterActiveMode{}, nil
	default:
		return nil, errors.New("Too many arguments")

	}

}
