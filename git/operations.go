package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Add() {
	gitAddCmd := exec.Command("git", "add", ".")
	gitAddCmd.Stdout = os.Stdout
	gitAddCmd.Stderr = os.Stderr
	err := gitAddCmd.Run()
	if err != nil {
		fmt.Println("⚠️ Error adding all files")
	} else {
		fmt.Println("===================================")
		fmt.Println("🚀 files added successfully")
		fmt.Println("===================================")

	}
}

func Commit(message string) {
	gitCommitCmd := exec.Command("git", "commit", "-m", message)
	gitCommitCmd.Stdout = os.Stdout
	gitCommitCmd.Stderr = os.Stderr
	err := gitCommitCmd.Run()
	if err != nil {
		fmt.Println("⚠️ Error commit files")
	} else {
		fmt.Println("===================================")
		fmt.Println("🚀 files commited successfully")
		fmt.Println("===================================")

	}
}

func Push() {
	gitPushCmd := exec.Command("git", "push", "origin", "HEAD")
	gitPushCmd.Stdout = os.Stdout
	gitPushCmd.Stderr = os.Stderr
	err := gitPushCmd.Run()
	if err != nil {
		fmt.Println("⚠️ Error push files")
	} else {
		fmt.Println("===================================")
		fmt.Println("🚀 files pushed successfully")
		fmt.Println("===================================")
	}
}

func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if strings.TrimSpace(out.String()) == "" {
		fmt.Println("⚠️  No staged changes detected. Did you forget to `git add`?")
		return "", fmt.Errorf("no staged changes found")
	}
	return out.String(), err
}
