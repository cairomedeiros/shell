package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CommandStruct struct {
	command string
	args    []string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error")
	}

	text_tokens := strings.Split(text, " ")

	result := CommandStruct{command: text_tokens[0], args: text_tokens[1:]}

	fmt.Println(result.command)

	for index := range result.args {
		fmt.Printf("%s \n", result.args[index])
	}

}
