package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func Add() {
	gitAddCmd := exec.Command("git", "add", ".")
	gitAddCmd.Stdout = os.Stdout
	gitAddCmd.Stderr = os.Stderr
	err := gitAddCmd.Run()
	if err != nil {
		fmt.Println("❌ Error adding all files")
	} else {
		fmt.Println("⭕️ files added successfully")
	}
}

func Commit(message string) {
	gitCommitCmd := exec.Command("git", "commit", "-m", message)
	gitCommitCmd.Stdout = os.Stdout
	gitCommitCmd.Stderr = os.Stderr
	err := gitCommitCmd.Run()
	if err != nil {
		fmt.Println("❌ Error commit files")
	} else {
		fmt.Println("⭕️ files commited successfully")
	}
}

func Push() {
	gitPushCmd := exec.Command("git", "push", "origin", "HEAD")
	gitPushCmd.Stdout = os.Stdout
	gitPushCmd.Stderr = os.Stderr
	err := gitPushCmd.Run()
	if err != nil {
		fmt.Println("❌ Error push files")
	} else {
		fmt.Println("⭕️ files pushed successfully")
	}
}

func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}
