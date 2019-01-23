package cmd

import (
	"github.com/hectorhuertas/resman/hello"
	"github.com/spf13/cobra"
)

func init() {
	ResmanCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Greets you",
	Run: func(cmd *cobra.Command, args []string) {
		hello.Hola()
	},
}
