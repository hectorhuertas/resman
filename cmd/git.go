package cmd

import (
	"fmt"

	"github.com/hectorhuertas/resman/git"
	"github.com/spf13/cobra"
)

func init() {
	ResmanCmd.AddCommand(gitCmd)
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "list local repos",
	// Restricted arguments just for fun
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"git", "two"},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(strings.Join(gitLocal.GetRepos("/Users/hh/xdev"), "\n"))
		fmt.Print(git.MockGet())
		git.Get()

	},
}
