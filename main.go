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

type cmdFunc map[string]func(args []string)

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := cmdFunc{
		"cd": func(args []string) {
			err := os.Chdir(strings.Join(args, " "))
			if err != nil {
				fmt.Println(err)
			}
		},
		"pwd": func(args []string) {
			cwd, _ := os.Getwd()
			fmt.Println("cwd:", cwd)
		},
		"exit": func(args []string) {
			os.Exit(1)
		},
	}

	for {
		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)
		text_tokens := strings.Fields(text)

		if text == "" {
			continue
		}

		result := CommandStruct{command: text_tokens[0], args: text_tokens[1:]}

		for index := range result.args {
			fmt.Printf("%s \n", result.args[index])
		}

		if fn, ok := commands[result.command]; ok {
			fn(result.args)
			continue
		}

		if result.command != "cd" && result.command != "pwd" && result.command != "exit" {
			cmd := exec.Command(result.command, result.args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			cmd.Run()
		}

	}

}
