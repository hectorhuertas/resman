package cmd

import (
	"fmt"

	"github.com/hectorhuertas/resman/git"
	"github.com/spf13/cobra"
)

func init() {
	ResmanCmd.AddCommand(gitCmd)
	gitCmd.Flags().StringP("root", "r", "/Users/hh/xdev", "root path to begin repository lookup")
	gitCmd.Flags().StringP("exclude", "e", "/Users/hh/xdev/go", "exclude a path from the repository lookup")
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "list local repos",
	// Restricted arguments just for fun
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"git", "two"},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(strings.Join(gitLocal.GetRepos("/Users/hh/xdev"), "\n"))
		fmt.Println(cmd.Flag("root").Value)
		fmt.Println(cmd.Flag("exclude").Value)
		fmt.Print(git.MockGet())
		git.Get()
	},
}
