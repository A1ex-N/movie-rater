package main

import (
	"bufio"
	"fmt"
  "os"
)

func GetInput(prompt string, variable *string) {
	// Have to use scanner because fmt.Scanln filters out whitespace :/
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	if scanner.Scan() {
		*variable = scanner.Text()
	}
}
