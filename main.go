package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CommandStruct struct {
	command string
	args    []string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)
		text_tokens := strings.Fields(text)

		if text == "" {
			continue
		}

		if text_tokens[0] == "exit" {
			break
		}

		result := CommandStruct{command: text_tokens[0], args: text_tokens[1:]}

		for index := range result.args {
			fmt.Printf("%s \n", result.args[index])
		}

		cmd := exec.Command(result.command, result.args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		cmd.Run()
	}

}
