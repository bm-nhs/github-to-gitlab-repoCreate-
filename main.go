package main

import (
	"context"
	"github-to-gitlab-repoCreate/gitlab"
	"github.com/google/go-github/github"
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
	githubToken := string(os.Getenv("githubPAT"))
	githubTarget := string(os.Getenv("githubTarget"))
	gitlabNamespaceID := string(os.Getenv("gitlabNamespaceID"))
	gitlabToken := string(os.Getenv("gitlabToken"))
	ctx := context.Background()
	StaticTokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)

	newClient := oauth2.NewClient(ctx, StaticTokenSource)
	client := github.NewClient(newClient)
	//TODO fix pagination manual process right now...
	var listOptions = github.ListOptions{
		Page: 2,
		PerPage: 100,
	}
	repositoryListByOrgOptions := github.RepositoryListByOrgOptions{
		Type: "all",
		ListOptions: listOptions,
	}
	repositories, _, err := client.Repositories.ListByOrg(ctx, githubTarget, &repositoryListByOrgOptions)
	if err != nil {
		println(err)
		return
	}
	// For each GitHub Repo Within the target Org
	// Create a GitLab Repo
	for i := 0; i < len(repositories); i++ {
		repo := *repositories[i].Name
		payload := gitlab.CreateRepoPayload{
			NamespaceID: gitlabNamespaceID,
			Name:        repo,
		}
		println(repo)
		err := gitlab.CreateRepo(payload, gitlabToken)
		if err != nil {
			println(err)
		}
	}
}
