package main

import (
	"context"
	"github-to-gitlab-repoCreate/gitlab"
	"github.com/google/go-github/v38/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		println(err)
		return
	}
	githubtoken := string(os.Getenv("githubPAT"))
	githubTarget := string(os.Getenv("githubTarget"))
	gitlabNamespaceID := string(os.Getenv("gitlabNamespaceID"))
	gitlabToken := string(os.Getenv("gitlabToken"))
	ctx := context.Background()
	StaticTokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubtoken},
	)

	newClient := oauth2.NewClient(ctx, StaticTokenSource)
	client := github.NewClient(newClient)
	repos, _, err := client.Repositories.ListByOrg(ctx, githubTarget, nil)
	if err != nil {
		println(err)
		return
	}
	// For each GitHub Repo Within the target Org
	// Create a GitLab Repo
	for i := 0; i < len(repos); i++ {
		repo := *repos[i].Name
		payload := gitlab.CreateRepoPayload{
			NamespaceID: gitlabNamespaceID,
			Name:        repo,
		}

		err := gitlab.CreateRepo(payload, gitlabToken)
		if err != nil {
			println(err)
		}
	}
}
