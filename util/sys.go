package util

import (
	"fmt"
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
