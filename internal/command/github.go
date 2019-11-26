package command

import (
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

func githubClient() *github.Client {
	tokenSrc := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: rootFlags.oAuthToken})
	tokenClt := oauth2.NewClient(rootCfg.ctx, tokenSrc)

	return github.NewClient(tokenClt)
}
