package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	cyan  = lipgloss.Color("#00ADD8")
	gold  = lipgloss.Color("#FFD700")
	gray  = lipgloss.Color("#5F5F5F")
	green = lipgloss.Color("#00FF00")

	styleTitle   = lipgloss.NewStyle().Bold(true).Foreground(cyan).Padding(1).Border(lipgloss.RoundedBorder())
	stylePrompt  = lipgloss.NewStyle().Foreground(gray).Bold(true)
	styleAI      = lipgloss.NewStyle().Foreground(gold).Italic(true)
	styleSuccess = lipgloss.NewStyle().Foreground(green)
)

var selectedModel string

// System Prompt in English, but instructed to reply in user's language
const systemPrompt = `You are Hexa Code V1, a professional software engineering assistant. 
Instructions:
1. Provide concise, expert-level technical advice.
2. ALWAYS detect the user's input language and respond in that same language.
3. If the user greets you in Turkish, respond in Turkish. If English, respond in English.`

func main() {
	fmt.Println(styleTitle.Render(" HEXA CODE V1 | ENTERPRISE LOCAL AGENT "))

	models := getOllamaModels()
	if len(models) == 0 {
		fmt.Println("Error: No models found. Please ensure Ollama is running and models are pulled.")
		return
	}

	fmt.Println(styleSuccess.Render("\n[Available Local Models]"))
	for i, m := range models {
		fmt.Printf("%d) %s\n", i+1, m)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(stylePrompt.Render("\nSelect model index: "))
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var idx int
	fmt.Sscanf(choice, "%d", &idx)
	if idx > 0 && idx <= len(models) {
		selectedModel = models[idx-1]
	} else {
		selectedModel = models[0]
	}

	fmt.Printf(styleSuccess.Render("\n✔ Active Session: %s\n\n"), selectedModel)

	for {
		fmt.Print(stylePrompt.Render(" hexa > "))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" || input == "quit" {
			fmt.Println("Terminating Hexa session...")
			break
		}
		if input == "" {
			continue
		}

		if !handleLocalCommands(input) {
			askOllama(input)
		}
	}
}

func getOllamaModels() []string {
	resp, err := http.Get("http://localhost:11434/api/tags")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var data struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}
	json.NewDecoder(resp.Body).Decode(&data)

	var list []string
	for _, m := range data.Models {
		list = append(list, m.Name)
	}
	return list
}

func handleLocalCommands(input string) bool {
	switch input {
	case "ls":
		files, _ := os.ReadDir(".")
		fmt.Println(styleSuccess.Render("\n[Directory Listing]"))
		for _, f := range files {
			icon := "📄"
			if f.IsDir() {
				icon = "📁"
			}
			fmt.Printf("%s %s\n", icon, f.Name())
		}
		fmt.Println()
		return true
	case "help":
		fmt.Println("\nCommands: ls, help, exit, or ask anything.\n")
		return true
	}
	return false
}

func askOllama(userInput string) {
	url := "http://localhost:11434/api/generate"
	fullPrompt := fmt.Sprintf("%s\n\nUser: %s", systemPrompt, userInput)

	payload, _ := json.Marshal(map[string]interface{}{
		"model":  selectedModel,
		"prompt": fullPrompt,
		"stream": false,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error: Communication with Ollama failed.")
		return
	}
	defer resp.Body.Close()

	var data struct {
		Response string `json:"response"`
	}
	json.NewDecoder(resp.Body).Decode(&data)

	fmt.Printf("\n%s %s\n\n", styleAI.Render("Hexa:"), data.Response)
}
