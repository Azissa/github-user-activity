package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os" // This package is used to interact with the operating system, including reading terminal input.
)

// GitHubEvent is a data structure for storing responses from the GitHub API.
type GitHubEvent struct {
	Type string `json:"type"` // For Ex: PushEvent, WatchEvent
	Repo struct {
		Name string `type:"name"` // Repository namw
	} `json:"repo"`
}

func main() {
	// os.Args is a built-in Go variable that contains a list of inputs from the terminal.
	// Let's check: is the number of inputs less than 2? if so then throw an error
	// (Index 0 = program name, Index 1 = username)
	if len(os.Args) < 2 {
		fmt.Println("Error: You must enter your username!")
		fmt.Println("Example: go run main.go 'username'")
	}

	// if the input is correct, than take the second input (index 1)
	username := os.Args[1]

	// Print the result for testing
	fmt.Printf("Searching github activity for user: %s\n", username)

	// 1. Arrange the URL
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// 2. Fetch API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch data: %v\n", err)
		return
	}
	defer resp.Body.Close() // Make sure to close the connection after you fetch it

	// 3. Check if the user is there or not
	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("Error: User not found!")
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: API issues (Status: %d)\n", resp.StatusCode)
		return
	}

	// 4. a place to hold the result (SLice/Github Event)
	var events []GitHubEvent

	// 5. Convert plain JSON into Go Struct
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		fmt.Printf("Failed to read JSON: %v\n", err)
		return
	}

	// 6. Check for if the activity is empty
	if len(events) == 0 {
		fmt.Printf("There are no recent activities for this user.")
		return
	}

	fmt.Println("Recent activities: ")
	printActivity(events)

}

func printActivity(events []GitHubEvent) {
	for _, event := range events {
		action := ""

		// Switch Case to differentiate messages based on event type
		switch event.Type {
		case "PushEvent":
			action = "Pushed commits to"
		case "WatchEvent":
			action = "Starred"
		case "CreateEvent":
			action = "Created a new resource in"
		case "IssuesEvent":
			action = "Opened/Closed an issue in"
		default:
			action = event.Type // Display the original type if we haven't handled it yet.
		}

		fmt.Printf("- %s %s\n", action, event.Repo.Name)
	}
}
