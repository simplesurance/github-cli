package command

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	buildVersion = "unknown-version"
	buildCommit  = "unknown-commit"
)

var rootCmd = &cobra.Command{
	Use:   "github-cli",
	Short: "A commandline client for the Github API",
}

type generalFlags struct {
	verbose bool
	csv     bool

	timeout time.Duration

	repository string
	oAuthToken string

	username string // TODO: is username needed?? Seems not!
}

var rootFlags generalFlags

func init() {
	rootCmd.Version = fmt.Sprintf("%s (%s)", buildVersion, buildCommit)

	rootCmd.PersistentFlags().BoolVarP(&rootFlags.verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&rootFlags.csv, "csv", "", false, "Print output in RFC4180 CSV format")

	rootCmd.PersistentFlags().DurationVarP(&rootFlags.timeout, "timeout", "t", 120*time.Second, "timeout for the operation")

	rootCmd.PersistentFlags().StringVarP(&rootFlags.repository, "repository", "r", "", "Github repository, format <OWNER/REPOSITORY>")
	rootCmd.PersistentFlags().StringVarP(&rootFlags.oAuthToken, "token", "p", "", "Github OAuth token for authentication")
	rootCmd.PersistentFlags().StringVarP(&rootFlags.username, "username", "u", "", "Github username for authentication")

}

type generalCfg struct {
	ctx             context.Context
	ctxCancelFunc   context.CancelFunc
	repositoryOwner string
	repository      string
}

var rootCfg generalCfg

func populateCfg() {
	if rootFlags.repository == "" {
		printErrln("--repository parameter must be passed")
		os.Exit(1)
	}

	spl := strings.Split(rootFlags.repository, "/")
	if len(spl) != 2 {
		printErrln("invalid --repository parameter passed")
		os.Exit(1)
	}

	rootCfg.repositoryOwner = spl[0]
	rootCfg.repository = spl[1]

	if rootFlags.timeout == 0 {
		rootCfg.ctx = context.Background()
	} else {
		rootCfg.ctx, rootCfg.ctxCancelFunc = context.WithTimeout(context.Background(), rootFlags.timeout)
	}
}

func Execute() {
	cobra.OnInitialize(populateCfg)
	if err := rootCmd.Execute(); err != nil {
		printErrln(err)
		os.Exit(1)
	}
}
