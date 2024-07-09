package main

import (
        "fmt"
        "log"
        "os"
        "time"
)

func main() {
        duration := 5 * time.Minute // Default duration
        if len(os.Args) > 1 {
                argDuration, err := time.ParseDuration(os.Args[1] + "m")
                if err != nil {
                        log.Fatalf("Invalid duration format: %v", err)
                }
                duration = argDuration
        }

        users := []string{"User1", "User2", "User3"}

        start := time.Now()
        end := start.Add(duration)

        fmt.Printf("Logging for %v minutes\n", duration.Minutes())

        for time.Now().Before(end) {
                for _, user := range users {
                        logMessage(user, "info", "This is an info message.")
                        logMessage(user, "debug", "This is a debug message.")
                        logMessage(user, "audit", "This is an audit message.")
                        logMessage(user, "trace", "This is a trace message.")
                }
                time.Sleep(1 * time.Second) // Simulate some delay between logs
        }
}

func logMessage(user, level, message string) {
        timestamp := time.Now().Format("2006-01-02 15:04:05")
        fmt.Printf("[%s] [%s] %s: %s\n", timestamp, level, user, message)
}
