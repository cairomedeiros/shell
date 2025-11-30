package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

type CommandStruct struct {
	command string
	args    []string
}

type cmdFunc map[string]func(args []string)

func setupSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for sig := range c {
			switch sig {
			case os.Interrupt:
				fmt.Println()
			case syscall.SIGTERM:
				fmt.Println("\n(SIGTERM)")
			}
		}
	}()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := cmdFunc{
		"cd": func(args []string) {
			if len(args) == 0 {
				fmt.Println("cd: missing args")
				return
			}
			if err := os.Chdir(strings.Join(args, " ")); err != nil {
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
	setupSignals()
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

		cmd := exec.Command(result.command, result.args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		cmd.Run()

		if err := cmd.Run(); err != nil {
			fmt.Printf("%s: command not found\n", result.command)
		}

	}

}
