package cmd

import (
	"fmt"
	"os"
)

func Exit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
