package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error")
	}

	text = strings.Join(strings.Fields(text), "")
	if text != "" {
		fmt.Printf("Você digitou %s", text)
	}
}
