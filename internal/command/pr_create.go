package command

import (
	"os"
	"simplesurance/github-cli/internal/command/flag"

	"github.com/google/go-github/v28/github"
	"github.com/spf13/cobra"
)

var prCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Pull-Request",
	Run:   prCreate,
}

var prCreateFlags struct {
	title               string
	branch              string
	baseBranch          string
	descriptionFile     flag.FilePathFlag
	maintainerCanModify bool
	draft               bool
}

func init() {
	prCmd.AddCommand(prCreateCmd)

	prCreateCmd.Flags().StringVarP(&prCreateFlags.title, "title", "", "", "Pull-Request title (required)")
	prCreateCmd.Flags().StringVarP(&prCreateFlags.branch, "branch", "", "", "Branch to merge (required)")
	prCreateCmd.Flags().StringVarP(&prCreateFlags.baseBranch, "base-branch", "", "", "Name of the branch to merge into (required)")
	prCreateCmd.Flags().VarP(&prCreateFlags.descriptionFile, "description-file", "", "Path to a file containing the Pull-Request description. Pass - to read from STDIN.")
	prCreateCmd.Flags().BoolVarP(&prCreateFlags.maintainerCanModify, "maintainer-can-modify", "m", true, "Indicates if the maintainer can modify the PR")
	prCreateCmd.Flags().BoolVarP(&prCreateFlags.draft, "draft", "", false, "Create the PR as draft")
}

func prCreate(cmd *cobra.Command, args []string) {
	var body string

	if prCreateFlags.descriptionFile.Path() != "" {
		body = mustReadFilePathFlagContentString(&prCreateFlags.descriptionFile)
	}

	clt := githubClient()
	pr, _, err := clt.PullRequests.Create(rootCfg.ctx, rootCfg.repositoryOwner, rootCfg.repository, &github.NewPullRequest{
		Title:               &prCreateFlags.title,
		Head:                &prCreateFlags.branch,
		Base:                &prCreateFlags.baseBranch,
		Body:                &body,
		MaintainerCanModify: &prCreateFlags.maintainerCanModify,
		Draft:               &prCreateFlags.draft,
	})
	if err != nil {
		printErrln(err)
		os.Exit(1)
	}

	output := [][]interface{}{
		[]interface{}{"Number", pr.GetNumber()},
		[]interface{}{"URL", pr.GetHTMLURL()},
	}

	mustWriteRows(output)
}
