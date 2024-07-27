package commands

import (
	"context"
	"fmt"
	"log"
	"os"

	"cli/utils"
	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	projectType string
)

var CreateRepoCmd = &cobra.Command{
	Use:   "create-repo [name]",
	Short: "Create a new GitHub repository",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		createRepo(repoName, projectType)
	},
}

var DeleteRepoCmd = &cobra.Command{
	Use:   "delete-repo [owner] [name]",
	Short: "Delete a GitHub repository",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		owner := args[0]
		repoName := args[1]
		deleteRepo(owner, repoName)
	},
}

func init() {
	CreateRepoCmd.Flags().StringVarP(&projectType, "type", "t", "", "Type of project (go, node, java)")
}

func createRepo(repoName, projectType string) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(false),
	}

	repo, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	fmt.Printf("Repository %s created successfully\n", *repo.Name)

	utils.CloneTemplateRepo(repoName, projectType, "hebertsanto", "boilerplate-node")
}

func deleteRepo(owner, repoName string) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	_, err := client.Repositories.Delete(ctx, owner, repoName)
	if err != nil {
		log.Fatalf("Failed to delete repository: %v", err)
	}

	fmt.Printf("Repository %s/%s deleted successfully\n", owner, repoName)
}
