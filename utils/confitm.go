package utils

import (
	"bufio"
	"fmt"
	"go_git_cli/color"
	"os"
	"strings"
)

// ConfirmCommitMessage is
func ConfirmCommitMessage(commitMessage string) bool {
	fmt.Println("💬 Suggested commit message:")
	fmt.Printf("%s\n\n", commitMessage)

	
	for {
		fmt.Print("commit & push are you ready? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		isReady, _ := reader.ReadString('\n')
		isReady = strings.TrimSpace(strings.ToLower(isReady))
		
		switch isReady {
		case "n":
			fmt.Println("===================================")
			fmt.Println("===================================")
			fmt.Println(color.Red, "🚫 Commit canceled.", color.Reset)
			return false
		case "y":
			fmt.Println("===================================")
			fmt.Println(color.Green, "🚀thank you for accepting commit message!", color.Reset)
			return true
		default:
			fmt.Println(color.Red, "🚫 Please type 「y」 or 「n」", color.Reset)
		}

	}

}
