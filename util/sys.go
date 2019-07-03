package util

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mitchellh/go-homedir"
)

func FindHome() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}

func ReadLines(streamOrFile io.Reader) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
