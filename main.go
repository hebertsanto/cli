package main

import (
	"cli/commands"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "github-cli",
		Short: "GitHub CLI to create and delete repositories",
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(commands.CreateRepoCmd)
	rootCmd.AddCommand(commands.DeleteRepoCmd)
}
