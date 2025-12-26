// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the argument input view update handlers. It processes
// user input for filling command arguments, handles navigation between
// argument fields, and submits the final command.
package ui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Args View Update
// =============================================================================

// updateArgs handles input in the argument input view.
func (m App) updateArgs(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keyCtrlC:
			log.Printf("User cancelled application from args view")
			m.Cancelled = true
			return m, tea.Quit

		case keyEsc:
			log.Printf("User exited args view, returning to search")
			m.state = stateSearch
			m.searchInput.Focus()
			return m, nil

		case keyEnter:
			return m.submitArgs()

		case keyTab, keyDown:
			m = m.nextArg()
			return m, nil

		case keyShiftTab, keyUp:
			m = m.prevArg()
			return m, nil
		}
	}

	m.argInputs[m.argCursor], cmd = m.argInputs[m.argCursor].Update(msg)
	return m, cmd
}

// nextArg moves to the next argument field.
func (m App) nextArg() App {
	m.argInputs[m.argCursor].Blur()
	m.argCursor = (m.argCursor + 1) % len(m.args)
	m.argInputs[m.argCursor].Focus()
	return m
}

// prevArg moves to the previous argument field.
func (m App) prevArg() App {
	m.argInputs[m.argCursor].Blur()
	m.argCursor--
	if m.argCursor < 0 {
		m.argCursor = len(m.args) - 1
	}
	m.argInputs[m.argCursor].Focus()
	return m
}

// submitArgs submits the arguments and builds the final command.
func (m App) submitArgs() (tea.Model, tea.Cmd) {
	for i := range m.args {
		m.args[i].Value = m.argInputs[i].Value()
	}

	if !model.HasEmptyArgs(m.args) {
		m.FinalCommand = model.BuildCommand(m.selectedCheat.Command, m.args)
		log.Printf("Command built from args: %s", m.FinalCommand)
		return m, tea.Quit
	}
	log.Printf("WARNING: Attempted to submit args with empty values")
	return m, nil
}

