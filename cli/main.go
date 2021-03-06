package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/danfragoso/milho"
)

func main() {
	if len(os.Args) > 2 && strings.HasPrefix(os.Args[1], "-c") {
		file, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		target := strings.Split(os.Args[1], "-c")[1]
		err = compileMilho(string(file), target)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	} else if len(os.Args) >= 2 && os.Args[1] != "" {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		runFile(string(file))
	} else {
		initREPL()
	}
}

func runFile(file string) {
	_, e := milho.RunRaw(file)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
