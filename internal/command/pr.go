package command

import "github.com/spf13/cobra"

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Interact with Pull-Requests",
}

func init() {
	rootCmd.AddCommand(prCmd)
}
