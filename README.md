# GitHub User Activity CLI

A simple Command Line Interface (CLI) built with **Go** to fetch and display recent GitHub user activity.
this project idea is from https://roadmap.sh/projects/github-user-activity

## 🚀 About This Project
This project is my first project in learning the Go programming language. The goal is to understand the basics of networking in Go, handle terminal input, and parse JSON data from the GitHub API without using external libraries.

## ✨ Features
- Retrieves the latest activity data from any GitHub user.
- Displays the activity type (Push, Star, Create, etc.) along with the repository name.
- Simple error handling (user not found, connection issues, etc.).
- Built using only the Go **Standard Library**.

## 🛠️ How to Run
Make sure you have [Go](https://go.dev/) installed.

1. Clone this repository:
   ```bash
   git clone [https://github.com/USERNAME_KAMU/github-user-activity.git](https://github.com/USERNAME_KAMU/github-user-activity.git)
2. Enter the project folder:
   ```bash
   cd github-user-activity
3. Run the program:
   ```bash
   go run main.go <github_username>

## 🧠 What I Learned
- Go Fundamentals: Structs, Slices, Control Flows (Switch Case).
- Networking: Using the net/http package to call REST APIs.
- JSON Handling: Using encoding/json to decode data.
- Error Handling: The practice of if err != nil, which is characteristic of Go.
