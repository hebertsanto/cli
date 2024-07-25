package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	rootCmd = &cobra.Command{
		Use:   "github-cli",
		Short: "GitHub CLI to create and delete repositories",
	}

	projectType string
)

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
		createRepo(repoName, projectType)
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

	createRepoCmd.Flags().StringVarP(&projectType, "type", "t", "", "Type of project (go, node, java)")
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

	cloneTemplateRepo(repoName, projectType)
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

func cloneTemplateRepo(repoName, projectType string) {
	var templateRepo string

	switch projectType {
	case "node":
		templateRepo = "https://github.com/hebertsanto/node-boilerplate"
	default:
		fmt.Println("Unsupported project type:", projectType)
		return
	}

	cliDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	parentDir := filepath.Dir(cliDir)

	repoPath := filepath.Join(parentDir, repoName)

	cmd := exec.Command("git", "clone", templateRepo, repoPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to clone repository: %v", err)
	}

	fmt.Printf("Repository %s cloned successfully\n", templateRepo)

	err = os.Chdir(repoPath)
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	err = os.RemoveAll(".git")
	if err != nil {
		log.Fatalf("Failed to remove .git directory: %v", err)
	}

	cmd = exec.Command("git", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to initialize new Git repository: %v", err)
	}

	cmd = exec.Command("git", "remote", "add", "origin", "https://github.com/hebertsanto/"+repoName+".git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to add remote: %v", err)
	}

	cmd = exec.Command("git", "branch", "-M", "main")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to rename branch to main: %v", err)
	}

	cmd = exec.Command("git", "add", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to add files to repository: %v", err)
	}

	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to commit: %v", err)
	}

	cmd = exec.Command("git", "push", "-u", "origin", "main")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to push to GitHub: %v", err)
	}

	fmt.Printf("Repository %s pushed to GitHub successfully\n", repoName)
}
