// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"charm.land/bubbles/v2/textinput"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/variables"
)

// =============================================================================
// Model State
// =============================================================================

// Model holds the state of the argument input view.
type Model struct {
	// Selected cheat being filled in.
	Cheat *model.Cheat
	// Command arguments with current values.
	Args []model.Argument

	// Per-argument text inputs.
	Inputs []textinput.Model
	// Focused field index.
	Cursor int
}

// =============================================================================
// Model Operations
// =============================================================================

// New creates an args model for the cheat, pre-filling values from globals
// and argument defaults. Focus lands on the first empty field.
func New(cheat *model.Cheat, globals *variables.Store) Model {
	args := model.ParseArguments(cheat.Command)

	// Globals win over defaults so saved values carry forward.
	for i, arg := range args {
		if globalVal, ok := globals.Get(arg.Name); ok {
			args[i].Value = globalVal
		} else if arg.DefaultValue != "" {
			args[i].Value = arg.DefaultValue
		}
	}

	m := Model{Cheat: cheat, Args: args}
	m.Inputs = make([]textinput.Model, len(args))
	firstEmpty := -1

	for i, arg := range args {
		ti := textinput.New()
		ti.Placeholder = arg.Name
		ti.CharLimit = config.ArgCharLimit
		ti.SetWidth(40)
		ti.SetValue(arg.Value)

		if arg.Value == "" && firstEmpty == -1 {
			firstEmpty = i
			ti.Focus()
		}
		m.Inputs[i] = ti
	}

	// All filled means focus starts on the first field.
	if firstEmpty == -1 {
		m.Inputs[0].Focus()
		m.Cursor = 0
	} else {
		m.Cursor = firstEmpty
	}

	return m
}

// Next moves to the next argument field.
func (m Model) Next() Model {
	m.Inputs[m.Cursor].Blur()
	m.Cursor = (m.Cursor + 1) % len(m.Args)
	m.Inputs[m.Cursor].Focus()
	return m
}

// Prev moves to the previous argument field.
func (m Model) Prev() Model {
	m.Inputs[m.Cursor].Blur()
	m.Cursor--
	if m.Cursor < 0 {
		m.Cursor = len(m.Args) - 1
	}
	m.Inputs[m.Cursor].Focus()
	return m
}
