package cmd

import (
	"fmt"

	"github.com/hectorhuertas/resman/git"
	"github.com/spf13/cobra"
)

func init() {
	ResmanCmd.AddCommand(gitCmd)
	gitCmd.Flags().StringP("root", "r", "/Users/hh/xdev", "root path to begin repository lookup")
	gitCmd.Flags().StringArrayP("exclude", "e", []string{}, "exclude a path from the repository lookup")
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "list local repos",
	// Restricted arguments just for fun
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"git", "two"},
	Run: func(cmd *cobra.Command, args []string) {
		root, _ := cmd.Flags().GetString("root")
		exclude, _ := cmd.Flags().GetStringArray("exclude")
		out := git.Scan(root, exclude)
		fmt.Println(out)
	},
}
