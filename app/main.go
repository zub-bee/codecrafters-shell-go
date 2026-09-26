package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)

		}

		command = strings.TrimSpace(command[:len(command)-1])

		if strings.TrimSpace(command) == "exit" {
			os.Exit(0)
		}

		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		}

		if command != "" {
			fmt.Printf("%s: command not found\n", command)
		}

	}
}
