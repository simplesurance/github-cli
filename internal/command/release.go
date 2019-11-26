package command

import "github.com/spf13/cobra"

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Interact with Releases",
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}
