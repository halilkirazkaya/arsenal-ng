// Package main is the entry point for arsenal-ng, a modern pentest command launcher.
//
// This program loads cheat files, initializes the TUI, and outputs the selected
// command to the terminal. Inspired by https://github.com/Orange-Cyberdefense/arsenal
package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/loader"
	"github.com/halilkirazkaya/arsenal-ng/internal/output"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Setup logging to file (same directory as variables.json)
	logPath, err := config.GetLogPath()
	if err != nil {
		// Log error but don't fail - logging is optional
		fmt.Fprintf(os.Stderr, "Warning: failed to get log path: %v\n", err)
	} else {
		logFile, err := tea.LogToFile(logPath, config.AppName)
		if err != nil {
			// Log error but don't fail - logging is optional
			fmt.Fprintf(os.Stderr, "Warning: failed to setup logging: %v\n", err)
		} else {
			defer logFile.Close()
			log.Printf("Application started, logging to: %s", logPath)
		}
	}

	// Load cheat files
	log.Printf("Loading cheat files...")
	cheats, err := loader.Load()
	if err != nil {
		log.Printf("ERROR: Failed to load cheats: %v", err)
		return fmt.Errorf("failed to load cheats: %w", err)
	}

	log.Printf("Loaded %d cheat(s) successfully", len(cheats))

	if len(cheats) == 0 {
		log.Printf("ERROR: No cheats found")
		return fmt.Errorf("no cheats found")
	}

	// Run TUI
	log.Printf("Starting TUI...")
	app := ui.New(cheats)
	program := tea.NewProgram(app, tea.WithAltScreen())

	result, err := program.Run()
	if err != nil {
		log.Printf("ERROR: TUI error: %v", err)
		return fmt.Errorf("TUI error: %w", err)
	}

	// Handle result
	model := result.(ui.App)

	if model.Cancelled {
		log.Printf("Application cancelled by user")
		return nil
	}

	if model.FinalCommand == "" {
		log.Printf("No command selected")
		return nil
	}

	// Output command to terminal
	log.Printf("Outputting command to terminal: %s", model.FinalCommand)
	output.ToTerminal(model.FinalCommand)
	log.Printf("Application completed successfully")

	return nil
}

