package command

import (
	"os"
	"simplesurance/github-cli/internal/command/flag"
)

func mustReadFilePathFlagContentString(flag *flag.FilePathFlag) string {
	content, err := flag.FileContentString()
	if err != nil {
		printErrln(err)
		os.Exit(1)
	}

	return content
}
