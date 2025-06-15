/*
Copyright © 2025 IwatsukaYura <iwatsukayura@gmail.com>
*/
package cmd

import (
	"fmt"
	"go_git_cli/color"
	"go_git_cli/git"
	"go_git_cli/openai"
	"go_git_cli/utils"

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
		fmt.Println(color.Cyan, "🚀 Executing Git Auto Commit Flow", color.Reset)

		git.Add()
		if useAI {
			diff, err := git.GetDiff()

			if err != nil || diff == "" {
				fmt.Println("===================================")
				fmt.Println(color.Red, "⚠️ Error getting git diff:", err, color.Reset)
				fmt.Println("===================================")
				return
			}

			msg, err := openai.GenerateCommitMessageWithOllama(diff)
			if err != nil {
				fmt.Println("===================================")
				fmt.Println(color.Red, "⚠️ Error generating commit message:", err, color.Reset)
				fmt.Println("===================================")
				return
			}
			commitMsg = msg
			fmt.Println("===================================")
			fmt.Println(color.Cyan, "💬 AI-generated commit message:", color.Reset)
			fmt.Println("===================================")
			fmt.Println(commitMsg)
		} else {
			fmt.Println("===================================")
			fmt.Println(color.Yellow, "⚠️ You don't use AI to create commit message", color.Reset)
			fmt.Println(color.Yellow, "⚠️ If you'd like to use AI to create commit message, attach option --ai", color.Reset)
			fmt.Println("===================================")
			commitMsg = message
		}
		if utils.ConfirmCommitMessage(commitMsg) {
			git.Commit(commitMsg)
		} else {
			return
		}

		git.Push()
	},
}

func init() {

	rootCmd.AddCommand(autoCmd)
	autoCmd.Flags().StringVarP(&message, "message", "m", "commit", "commit message (default: \"commit\")")
	autoCmd.Flags().BoolVar(&useAI, "ai", false, "Use AI to generate commit message")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
