package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	duration := 5 // Default duration in minutes
	if len(os.Args) > 1 {
		// Parse duration from command line argument
		var err error
		duration, err = parseDuration(os.Args[1])
		if err != nil {
			log.Fatalf("Invalid duration: %v", err)
		}
	}

	// Calculate end time based on duration
	endTime := time.Now().Add(time.Duration(duration) * time.Minute)

	// Print logs continuously until endTime
	for time.Now().Before(endTime) {
		printLogs()
		time.Sleep(time.Second) // Adjust frequency of logs by changing sleep duration
	}
}

func parseDuration(arg string) (int, error) {
	duration, err := time.ParseDuration(arg)
	if err != nil {
		// If parsing as duration fails, try parsing as an integer (assuming minutes)
		var minutes int
		_, err := fmt.Sscanf(arg, "%d", &minutes)
		if err != nil {
			return 0, fmt.Errorf("invalid duration format")
		}
		return minutes, nil
	}
	return int(duration.Minutes()), nil
}

func printLogs() {
	fmt.Println("Logging Info message...")
	fmt.Println("Debug message...")
	fmt.Println("Audit message...")
	fmt.Println("Tracing message...")
	fmt.Println("--------------------------------------")
}
