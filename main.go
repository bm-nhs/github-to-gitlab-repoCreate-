package main

import (
	"context"
	"os"
	"github-to-gitlab-repoCreate/gitlab"
	"github.com/google/go-github/v38/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	godotenv.Load(".env")
	gitHubtoken := string(os.Getenv("githubPAT"))
	githubTarget := string(os.Getenv("githubTarget"))
	gitlabTarget := string(os.Getenv("gitlabTarget"))
	gitlabToken := string(os.Getenv("gitlabToken"))
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitHubtoken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repos, _, _ := client.Repositories.ListByOrg(ctx, githubTarget, nil)


	// For each GitHub Repo Within the target Org
	// Create a GitLab Repo
	for i := 0; i < len(repos); i++ {
		repo := *repos[i].Name
		println(repo)
		payload := gitlab.Payload{
			NamespaceID: gitlabTarget,
			Name: repo,
		}

		err := gitlab.Send(payload, gitlabToken)
		if err != nil {
			println(err)
		}

	}
}
