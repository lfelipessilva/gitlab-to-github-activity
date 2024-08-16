package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func findLastCommmitDate() time.Time {
	cmd := exec.Command(
		"git",
		"-C",
		"./repo",
		"log",
		"-1",
		"--format=%cd",
	)

	out, err := cmd.Output()

	if err != nil {
		log.Printf("Error finding last commit date: %s\n", err)

		startOfYear3000, err := time.Parse("2006-01-02", "3000-01-01")
		if err != nil {
			log.Printf("Error on start of year 3000: %s\n", err)
		}

		return startOfYear3000
	}

	commitDateStr := string(out)
	lastCommitDateTime, err := time.Parse("2006-01-02", commitDateStr)

	return lastCommitDateTime
}

func main() {
	// Open the JSON file
	file, err := os.Open("commits.json")
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	// Read the file's content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	// Unmarshal the JSON data into a map
	var dates map[string]int
	err = json.Unmarshal(byteValue, &dates)
	if err != nil {
		log.Fatalf("Could not unmarshal file: %v", err)
	}

	lastCommmitDate := findLastCommmitDate()

	// Iterate over the dates and values
	for date, value := range dates {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Printf("Skipping invalid date %s: %v", date, err)
			continue
		}

		if parsedDate.After(lastCommmitDate) {
			continue
		}

		for i := 0; i <= value; i++ {
			err := runGitCommand(
				"-C",
				"./repo",
				"commit",
				"--date", date+" 12:00:00",
				"-m", "add contribution",
				"--allow-empty",
			)
			if err != nil {
				log.Printf("Failed to run git command: %v", err)
			}
		}
	}
}
