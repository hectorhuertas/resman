package main

import (
	"fmt"
	"os"

	"github.com/hectorhuertas/resman/cmd"
)

func main() {
	if err := cmd.ResmanCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
