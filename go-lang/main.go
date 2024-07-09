package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	AUDIT LogLevel = "AUDIT"
	WARNING LogLevel = "WARNING"
	DEBUG LogLevel = "DEBUG"
)

type LogRecord struct {
	Level   LogLevel `json:"level"`
	User    string   `json:"user"`
	Message string   `json:"message"`
	Time    string   `json:"time"`
}

func main() {
	users := []string{"user1", "user2", "user3"}
	levels := []LogLevel{INFO, AUDIT, WARNING, DEBUG}
	messages := []string{
		"Started processing request",
		"Login attempt from unknown IP",
		"Low disk space detected",
		"Debug message: Validating data",
	}

	rand.Seed(time.Now().UnixNano())

	for {
		level := levels[rand.Intn(len(levels))]
		user := users[rand.Intn(len(users))]
		message := messages[rand.Intn(len(messages))]
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		logRecord := LogRecord{
			Level:   level,
			User:    user,
			Message: message,
			Time:    currentTime,
		}

		fmt.Println(logRecord)
		time.Sleep(time.Second * 2) // Adjust sleep duration as needed
	}
}
