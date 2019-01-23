package main

import (
	"fmt"
	"os"

	"github.com/hectorhuertas/resman/cmd"
)

func main() {
	fmt.Println("Start")
	cmd.HelloCmd()
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
