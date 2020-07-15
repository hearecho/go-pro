package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd  = &cobra.Command{
	Use:"version",
	Short:"Print the version number of Go-Tool",
	Long:`Test Verison study`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go-Tool version v1.0")
	},
}