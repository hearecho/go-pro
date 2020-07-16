package cmd

import "github.com/spf13/cobra"

/**
帮助文档
 */
func ShowHelp(cmd *cobra.Command)  {
	cmd.Help()
}
