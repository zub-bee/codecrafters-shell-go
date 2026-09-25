package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for true {
		fmt.Print("$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if command != "" {
			fmt.Printf("%s: command not found\n", command[:len(command)-1])
		}

	}
}
