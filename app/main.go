package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if command != "" {
		str := strings.Split(command, "\n")
		fmt.Printf("%s: command not found\n", str[0])
	}
}
