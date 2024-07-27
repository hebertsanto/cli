package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func CloneTemplateRepo(repoName, projectType, username, templateRepoName string) {
	var templateRepo string

	switch projectType {
	case "node":
		templateRepo = fmt.Sprintf("https://github.com/%s/%s", username, templateRepoName)
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

	if isEmpty, err := isDirEmpty(repoPath); err != nil {
		log.Fatalf("Failed to check directory: %v", err)
	} else if isEmpty {
		fmt.Println("No files found in the repository. Skipping initialization.")
		return
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

	cmd = exec.Command("git", "remote", "add", "origin", fmt.Sprintf("https://github.com/%s/%s.git", username, repoName))
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

func isDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return false, nil
}
