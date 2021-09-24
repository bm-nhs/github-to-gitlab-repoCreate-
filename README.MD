# github-to-gitlab-repoCreate
#[How to build and install Golang](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)
## Setting up .env
 Please see .envExample as a guide to create .env
 ```
 export githubPAT = ghp-github access token
 githubTarget = github target org
 gitlabTarget = gitlab target org
 gitlabNamespace = my-awesome-test1  ///TODOFIX
 gitlabToken = gitlab token
 ```


## Query GitHub API for all repositories within an organization:
```bigquery
curl "https://api.github.com/users/$GHUSER/repos?access_token=$GITHUB_API_TOKEN" | grep -w clone_url
```

## Create repository in GitLab with GitLab API V4
```bigquery
curl -H "Content-Type:application/json" https://gitlab.com/api/v4/projects?private_token=TOKEN -d "{ \"name\": \"newRepo\", \"namespace_id\": \"ID\" }"            
```

TOKEN = gitLab personal access token
###### (User Settings > Access Tokens)
ID    = gitLab group namespace_id
###### request all groups associated with account to get namespace id namespace.id
```
curl --header "PRIVATE-TOKEN: <token>" "https://gitlab.com/api/v4/namespaces"    
```
