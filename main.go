package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gmkohler/golox/scan"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		_, _ = fmt.Fprint(os.Stdout, "Usage: scanner [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		err := runFile(os.Args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := runPrompt()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func runFile(fp string) error {
	var err error

	b, err := os.ReadFile(fp)
	if err != nil {
		return fmt.Errorf("runFile: %w", err)
	}

	err = run(string(b))
	if err != nil {
		return fmt.Errorf("runFile: %w", err)
	}
	return nil
}

func runPrompt() error {
	stdin := bufio.NewReader(os.Stdin)
	for {
		_, err := fmt.Print("> ")
		if err != nil {
			return fmt.Errorf("runPrompt: %w", err)
		}
		text, _, err := stdin.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("runPrompt: %w", err)
		}
		err = run(string(text))

		if err != nil {
			return fmt.Errorf("runPrompt: %w", err)
		}
	}
	return nil
}

func run(source string) error {
	s := scan.NewSourceScanner(source)
	tokens, err := s.ScanTokens()
	if err != nil {
		return fmt.Errorf("run: %w", err)
	}
	for _, t := range tokens {
		fmt.Printf("%v\n", t)
	}
	return nil
}
