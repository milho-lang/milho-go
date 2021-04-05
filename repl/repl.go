package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/danfragoso/milho"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Milho REPL 🌽")
	fmt.Printf("© Danilo Fragoso <danilo.fragoso@gmail.com> - 2021\n")

	prompt()

	for scanner.Scan() {
		cmd := scanner.Text()
		if strings.TrimSpace(cmd) != "" {
			results := milho.Run(cmd)

			for _, result := range strings.Split(results, "\n") {
				r := strings.TrimSpace(result)
				if r != "" {
					fmt.Print("🍿 " + r + "\n")
				}
			}

			prompt()
		}
	}

	if scanner.Err() != nil {
		fmt.Printf("\n\nIO Err: %s", scanner.Err())
		os.Exit(1)
	}
}

func prompt() {
	fmt.Printf("🌽 > ")
}
