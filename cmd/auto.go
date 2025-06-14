/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// autoCmd represents the auto command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "You can execute git flow aytomatically",
	Long:  `This is a tool that allows you to perform git add, commit, and push in one go.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitAddCmd := exec.Command("git", "add", ".")
		gitAddCmd.Stdout = os.Stdout
		gitAddCmd.Stderr = os.Stderr
		err1 := gitAddCmd.Run()
		if err1 != nil {
			fmt.Println("❌ Error adding all files")
		} else {
			fmt.Println("⭕️ files added successfully")
		}

		gitCommitCmd := exec.Command("git", "commit", "-m", "first commit") //メッセージのところはのちのちLLM生成させたい
		gitCommitCmd.Stdout = os.Stdout
		gitCommitCmd.Stderr = os.Stderr
		err2 := gitCommitCmd.Run()
		if err2 != nil {
			fmt.Println("❌ Error commit files")
		} else {
			fmt.Println("⭕️ files commited successfully")
		}

		gitPushCmd := exec.Command("git", "push", "origin", "HEAD")
		gitPushCmd.Stdout = os.Stdout
		gitPushCmd.Stderr = os.Stderr
		err3 := gitPushCmd.Run()
		if err3 != nil {
			fmt.Println("❌ Error push files")
		} else {
			fmt.Println("⭕️ files pushed successfully")
		}

	},
}

func init() {
	rootCmd.AddCommand(autoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
