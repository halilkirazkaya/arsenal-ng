// Package keys centralizes the keyboard bindings and special command
// literals used across the TUI views.
package keys

// =============================================================================
// Navigation Bindings
// =============================================================================

// Up moves the selection up in the finder list and to the previous field in the args form.
const Up = "up"

// Down moves the selection down in the finder list and to the next field in the args form.
const Down = "down"

// Left shows the previous page in the tools view.
const Left = "left"

// Right shows the next page in the tools view.
const Right = "right"

// VimLeft mirrors Left on the tools pages.
const VimLeft = "h"

// VimRight mirrors Right on the tools pages.
const VimRight = "l"

// CtrlP moves the selection up, mirroring Up.
const CtrlP = "ctrl+p"

// CtrlN moves the selection down, mirroring Down.
const CtrlN = "ctrl+n"

// PgUp jumps one page upward in long lists.
const PgUp = "pgup"

// PgDown jumps one page downward in long lists.
const PgDown = "pgdown"

// =============================================================================
// Action Bindings
// =============================================================================

// Enter confirms the selection in the finder, args, and tools views, and returns from the variables and help views.
const Enter = "enter"

// Esc quits the search view or returns to it from other views.
const Esc = "esc"

// CtrlC cancels the finder search and the args form.
const CtrlC = "ctrl+c"

// Help opens the help view from an empty search box.
const Help = "?"

// Quit leaves the tools, variables, and help views.
const Quit = "q"

// Tab advances to the next argument field.
const Tab = "tab"

// ShiftTab returns to the previous argument field.
const ShiftTab = "shift+tab"

// =============================================================================
// Special Commands
// =============================================================================

// Set stores a variable typed as set key=value.
const Set = "set"

// Unset removes a variable typed as unset key.
const Unset = "unset"

// Variables opens the variables view.
const Variables = "variables"

// Tools opens the tools view.
const Tools = "tools"

// HelpCmd opens help when typed into the search box.
const HelpCmd = "help"
