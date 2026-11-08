// Package main is the entry point for arsenal-ng, a modern pentest command launcher.
//
// This program loads cheat files, initializes the TUI, and outputs the selected
// command to the terminal. Inspired by https://github.com/Orange-Cyberdefense/arsenal
package main

import (
	"fmt"
	"log"
	"os"
)

// =============================================================================
// Entry Point
// =============================================================================

// main runs the launcher and exits nonzero on failure.
func main() {
	if err := run(production{}); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// =============================================================================
// Run Pipeline
// =============================================================================

// run loads cheats, runs the TUI, and prefills the chosen command.
func run(d deps) error {
	// Logging is optional: a broken log path only warns and continues.
	logPath, err := d.logPath()
	if err != nil {
		d.warn("Warning: failed to get log path: %v", err)
	} else {
		logFile, err := d.setupLog(logPath)
		if err != nil {
			d.warn("Warning: failed to setup logging: %v", err)
		} else {
			defer logFile.Close()
			log.Printf("Application started, logging to: %s", logPath)
		}
	}

	log.Printf("Loading cheat files...")
	cheats, err := d.load()
	if err != nil {
		log.Printf("ERROR: Failed to load cheats: %v", err)
		return fmt.Errorf("failed to load cheats: %w", err)
	}

	log.Printf("Loaded %d cheat(s) successfully", len(cheats))

	if len(cheats) == 0 {
		log.Printf("ERROR: No cheats found")
		return fmt.Errorf("no cheats found")
	}

	log.Printf("Starting TUI...")
	app, err := d.interact(cheats)
	if err != nil {
		log.Printf("ERROR: TUI error: %v", err)
		return fmt.Errorf("TUI error: %w", err)
	}

	if app.Cancelled {
		log.Printf("Application cancelled by user")
		return nil
	}

	if app.FinalCommand == "" {
		log.Printf("No command selected")
		return nil
	}

	// A failed prefill only warns: the command is still printed by the TUI.
	log.Printf("Outputting command to terminal: %s", app.FinalCommand)
	if err := d.prefill(app.FinalCommand); err != nil {
		log.Printf("Terminal prefill unavailable (%v)", err)
		d.warn("")
		d.warn("arsenal-ng: could not prefill your shell (terminal injection is disabled by the kernel).")
		d.warn("")
		d.warn("To enable auto-prefill, see:")
		d.warn("  https://github.com/halilkirazkaya/arsenal-ng#linux-configuration")
		d.warn("")
	}
	log.Printf("Application completed successfully")

	return nil
}
