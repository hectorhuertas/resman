package cmd

import (
	"fmt"

	"github.com/hectorhuertas/resman/getgit"
	"github.com/spf13/cobra"
)

func init() {
	ResmanCmd.AddCommand(getgitCmd)
}

var getgitCmd = &cobra.Command{
	Use:   "getgit",
	Short: "gets a list of git repos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getgit.Exec())
	},
}
