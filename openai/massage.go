package openai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"go_git_cli/model"
)

func GenerateCommitMessageWithOllama(diff string) (string, error) {
	reqBody := model.OllamaRequest{
		Model:  "llama3.2", // or llama3, gemma
		Prompt: fmt.Sprintf("Write a concise and conventional commit message for this diff:\n\n%s", diff),
	}

	jsonData, _ := json.Marshal(reqBody)
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("❌ failed to call Ollama: %w", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var fullMessage string

	for scanner.Scan() {
		var part model.OllamaResponse
		line := scanner.Bytes()

		if err := json.Unmarshal(line, &part); err != nil {
			return "", fmt.Errorf("❌ failed to parse Ollama response: %w\nraw: %s", err, line)
		}
		fullMessage += part.Response
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("❌ failed to scan Ollama response: %w", err)
	}

	return fullMessage, nil
}
