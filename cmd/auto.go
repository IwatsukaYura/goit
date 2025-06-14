/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go_git_cli/git"
	"go_git_cli/openai"
	"strings"

	"github.com/spf13/cobra"
)

var message string
var useAI bool

// autoCmd represents the auto command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "You can execute git flow automatically",
	Long:  `This is a tool that allows you to perform git add, commit, and push in one go.`,

	Run: func(cmd *cobra.Command, args []string) {
		var commitMsg string
		git.Add()
		if useAI {
			diff, err := git.GetDiff()
			if err != nil {
				fmt.Println("❌ Error getting git diff:", err)
				return
			}

			if strings.TrimSpace(diff) == "" {
				fmt.Println("⚠️  No staged changes detected. Did you forget to `git add`?")
				return
			}

			msg, err := openai.GenerateCommitMessageWithOllama(diff)
			if err != nil {
				fmt.Println("❌ Error generating commit message:", err)
				return
			}
			commitMsg = msg
			fmt.Println("💬 AI-generated commit message:")
			fmt.Println(commitMsg)
		} else {
			if message == "" {
				fmt.Println("❌ Commit message is empty. Use -m to specify a message.")
				return
			}
			commitMsg = message
		}
		git.Commit(commitMsg)
		git.Push()
	},
}

func init() {

	rootCmd.AddCommand(autoCmd)
	autoCmd.Flags().StringVarP(&message, "message", "m", "commit", "commit message (default: \"commit\")")
	autoCmd.Flags().BoolVar(&useAI, "ai", true, "Use AI to generate commit message")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
