package cmd

import "fmt"
import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "resman",
	Short: "Resman is a resource manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wabalubadubdub")
	},
}
