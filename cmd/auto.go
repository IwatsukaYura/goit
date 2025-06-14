/*
Copyright © 2025 IwatsukaYura <iwatsukayura@gmail.com>
*/
package cmd

import (
	"fmt"
	"go_git_cli/color"
	"go_git_cli/git"
	"go_git_cli/openai"

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
		fmt.Println("===================================")
		fmt.Println(color.Green, "🚀 Executing Git Auto Commit Flow", color.Reset)
		fmt.Println("===================================")

		git.Add()
		if useAI {
			diff, err := git.GetDiff()

			if err != nil || diff == "" {
				fmt.Println(color.Red, "⚠️ Error getting git diff:", err, color.Reset)
				return
			}

			msg, err := openai.GenerateCommitMessageWithOllama(diff)
			if err != nil {
				fmt.Println(color.Red, "⚠️ Error generating commit message:", err, color.Reset)
				return
			}
			commitMsg = msg
			fmt.Println("===================================")
			fmt.Println("💬 AI-generated commit message:")
			fmt.Println("===================================")
			fmt.Println(commitMsg)
		} else {
			if message == "" {
				fmt.Println(color.Red, "⚠️ Commit message is empty. Use -m to specify a message.", color.Reset)
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
