package command

import (
	"os"
	"simplesurance/github-cli/internal/command/flag"
	"strconv"

	"github.com/google/go-github/v28/github"
	"github.com/spf13/cobra"
)

var commitCreateCommentCmd = &cobra.Command{
	Use:   "create-comment",
	Short: "Create a Commit Comment",
	Run:   commitCreateComment,
}

var commitCreateCommentFlags struct {
	commit        string
	commentFile   flag.FilePathFlag
	relFilePath   string
	diffLineIndex string
}

func init() {
	commitCmd.AddCommand(commitCreateCommentCmd)

	commitCreateCommentCmd.Flags().VarP(&commitCreateCommentFlags.commentFile, "comment-file", "", "Path to a file containing the comment (required). Pass - to read from STDIN. (required)")
	commitCreateCommentCmd.Flags().StringVarP(&commitCreateCommentFlags.relFilePath, "file", "", "", "Relative path of the file to comment on")
	commitCreateCommentCmd.Flags().StringVarP(&commitCreateCommentFlags.diffLineIndex, "diff-line-index", "", "", "Line index in the diff to comment on")

}

func commitCreateComment(cmd *cobra.Command, args []string) {
	var diffLineIdx int

	if commitCreateCommentFlags.commentFile.Path() == "" {
		printErrln("--comment-file parameter is required")
		os.Exit(1)
	}

	content := mustReadFilePathFlagContentString(&commitCreateCommentFlags.commentFile)

	if commitCreateCommentFlags.diffLineIndex != "" {
		var err error

		diffLineIdx, err = strconv.Atoi(commitCreateCommentFlags.diffLineIndex)
		if err != nil {
			printErrf("invalid argument passed to --line-number: '%s' is not a number\n", commitCreateCommentFlags.diffLineIndex)
			os.Exit(1)
		}
	}

	clt := githubClient()

	comment, _, err := clt.Repositories.CreateComment(
		rootCfg.ctx,
		rootCfg.repositoryOwner,
		rootCfg.repository,
		commitFlags.commit,
		&github.RepositoryComment{
			Body:     &content,
			Path:     &commitCreateCommentFlags.relFilePath,
			Position: &diffLineIdx,
		})
	if err != nil {
		printErrln(err)
		os.Exit(1)
	}

	output := [][]interface{}{
		[]interface{}{"URL", comment.GetHTMLURL()},
	}

	mustWriteRows(output)
}
