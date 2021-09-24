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
	godotenv.Load(".env")
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
	repos, _, _ := client.Repositories.ListByOrg(ctx, githubTarget, nil)

	// For each GitHub Repo Within the target Org
	// Create a GitLab Repo
	for i := 0; i < len(repos); i++ {
		repo := *repos[i].Name
		payload := gitlab.Payload{
			NamespaceID: gitlabNamespaceID,
			Name:        repo,
		}

		err := gitlab.CreateRepo(payload, gitlabToken)
		if err != nil {
			println(err)
		}
	}
}
