/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go_git_cli/git"

	"github.com/spf13/cobra"
)

var message string

// autoCmd represents the auto command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "You can execute git flow aytomatically",
	Long:  `This is a tool that allows you to perform git add, commit, and push in one go.`,
	Run: func(cmd *cobra.Command, args []string) {
		git.Add()
		git.Commit(message)
		git.Push()
	},
}

func init() {
	rootCmd.AddCommand(autoCmd)
	autoCmd.Flags().StringVarP(&message, "--message", "m", "commit", "commit")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
