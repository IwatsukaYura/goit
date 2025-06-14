package git

import (
	"fmt"
	"os"
	"os/exec"
)

func Add() {
	gitAddCmd := exec.Command("git", "add", ".")
	gitAddCmd.Stdout = os.Stdout
	gitAddCmd.Stderr = os.Stderr
	err1 := gitAddCmd.Run()
	if err1 != nil {
		fmt.Println("❌ Error adding all files")
	} else {
		fmt.Println("⭕️ files added successfully")
	}
}

func Commit(message string) {
	gitCommitCmd := exec.Command("git", "commit", "-m", message) //メッセージのところはのちのちLLM生成させたい
	gitCommitCmd.Stdout = os.Stdout
	gitCommitCmd.Stderr = os.Stderr
	err2 := gitCommitCmd.Run()
	if err2 != nil {
		fmt.Println("❌ Error commit files")
	} else {
		fmt.Println("⭕️ files commited successfully")
	}
}

func Push() {
	gitPushCmd := exec.Command("git", "push", "origin", "HEAD")
	gitPushCmd.Stdout = os.Stdout
	gitPushCmd.Stderr = os.Stderr
	err3 := gitPushCmd.Run()
	if err3 != nil {
		fmt.Println("❌ Error push files")
	} else {
		fmt.Println("⭕️ files pushed successfully")
	}
}
