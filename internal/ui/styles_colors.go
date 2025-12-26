// Package ui provides the terminal user interface for arsenal-ng.
//
// This file defines color constants and initializes the random logo color
// at application startup. It contains the base color palette used throughout
// the application.
package ui

import (
	"math/rand"
	"time"

	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Color Definitions
// =============================================================================

var (
	primaryColor   = lipgloss.Color(config.ColorPrimary)
	secondaryColor = lipgloss.Color(config.ColorSecondary)
	accentColor    = lipgloss.Color(config.ColorAccent)
	dimColor       = lipgloss.Color(config.ColorDim)
	brightColor    = lipgloss.Color(config.ColorBright)
)

// Random logo color - selected once at startup
var logoColor lipgloss.Color

func init() {
	rand.Seed(time.Now().UnixNano())
	logoColor = lipgloss.Color(config.TagColors[rand.Intn(len(config.TagColors))])
	logoStyle = lipgloss.NewStyle().Foreground(logoColor).Bold(true)
}

