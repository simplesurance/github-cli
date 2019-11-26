package command

import (
	"os"

	"github.com/google/go-github/v28/github"
	"github.com/spf13/cobra"
)

var releaseCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Release",
	Run:   releaseCreate,
}

var releaseCreateFlags struct {
	tagName     string
	targetRef   string
	name        string
	description string
	prerelease  bool
	draft       bool
}

func init() {
	releaseCmd.AddCommand(releaseCreateCmd)
	releaseCreateCmd.Flags().StringVarP(&releaseCreateFlags.tagName, "tag", "", "", "Git tag (required)")
	releaseCreateCmd.Flags().StringVarP(&releaseCreateFlags.targetRef, "target", "", "", "Target branch or commit ID where the tag is created. If the tag already exist, it's ignored")
	releaseCreateCmd.Flags().StringVarP(&releaseCreateFlags.name, "name", "", "", "The name of the release.")
	releaseCreateCmd.Flags().StringVarP(&releaseCreateFlags.description, "description", "", "", "Description of the release")
	releaseCreateCmd.Flags().BoolVarP(&releaseCreateFlags.prerelease, "prerelease", "", false, "Identify the release as a prerelease")
	releaseCreateCmd.Flags().BoolVarP(&releaseCreateFlags.draft, "draft", "", false, "Create a draft (unpublished) release")
}

func releaseCreate(cmd *cobra.Command, args []string) {
	clt := githubClient()

	release, _, err := clt.Repositories.CreateRelease(rootCfg.ctx, rootCfg.repositoryOwner, rootCfg.repository, &github.RepositoryRelease{
		TagName:         &releaseCreateFlags.tagName,
		TargetCommitish: &releaseCreateFlags.targetRef,
		Name:            &releaseCreateFlags.name,
		Body:            &releaseCreateFlags.description,
		Draft:           &releaseCreateFlags.draft,
		Prerelease:      &releaseCreateFlags.prerelease,
	})
	if err != nil {
		printErrln(err)
		os.Exit(1)
	}

	output := [][]interface{}{
		[]interface{}{"URL", release.GetHTMLURL()},
	}

	mustWriteRows(output)
}
