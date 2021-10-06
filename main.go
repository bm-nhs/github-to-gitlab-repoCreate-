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
	// Init env vars
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

	// Initialize options for pagination
	var listOptions = github.ListOptions{
		Page: 1,
		PerPage: 100,
	}
	repositoryListByOrgOptions := github.RepositoryListByOrgOptions{
		Type: "all",
		ListOptions: listOptions,
	}
	// Make an initial request to get number of repositories
	repositories, response, err := client.Repositories.ListByOrg(ctx, githubTarget, &repositoryListByOrgOptions)
	if err != nil {
		println(err)
		return
	}
	//Paginate through repositories.
	for i := 1; i < response.LastPage; i++ {
		listOptions.Page = i
		repositoryListByOrgOptions.ListOptions = listOptions
		pagination, _, err := client.Repositories.ListByOrg(ctx, githubTarget, &repositoryListByOrgOptions)
		if err != nil {
			println(err)
			return
		}
		repositories = append(repositories,pagination...)
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
