package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var rootCmd = &cobra.Command{
	Use:   "github-cli",
	Short: "GitHub CLI to create and delete repositories",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var createRepoCmd = &cobra.Command{
	Use:   "create-repo [name]",
	Short: "Create a new GitHub repository",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		createRepo(repoName)
	},
}

var deleteRepoCmd = &cobra.Command{
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
	rootCmd.AddCommand(createRepoCmd)
	rootCmd.AddCommand(deleteRepoCmd)
}

func createRepo(repoName string) {
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
