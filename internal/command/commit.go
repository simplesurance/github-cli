package command

import "github.com/spf13/cobra"

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Interact with Commits",
}

var commitFlags struct {
	commit string
}

func init() {
	rootCmd.AddCommand(commitCmd)

	rootCmd.PersistentFlags().StringVarP(&commitFlags.commit, "commit", "", "", "Commit SHA (required)")
}
