package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/GrazianoJoa/Glox/scan"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(65)
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	run(string(data))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("glox> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		line = strings.TrimRight(line, "\r\n")
		if len(line) == 0 {
			break
		}
		run(line)
	}
}

func run(source string) {
	scn := scan.NewScanner(source)
	tokenList, err := scn.ScanTokens()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, token := range tokenList {
		fmt.Println(token.String())
	}
}

