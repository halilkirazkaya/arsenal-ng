// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

// =============================================================================
// List Navigation
// =============================================================================

// MoveUp moves the cursor up in the list.
func (m Model) MoveUp() Model {
	if m.Cursor > 0 {
		m.Cursor--
		if m.Cursor < m.Offset {
			m.Offset = m.Cursor
		}
	} else if m.Cursor == 0 {
		// Backing off the first item clears the selection.
		m.Cursor = -1
	}
	return m
}

// MoveDown moves the cursor down in the list.
func (m Model) MoveDown() Model {
	visible := m.MaxVisible()
	if m.Cursor < 0 {
		// Advancing from no selection lands on the first item.
		if len(m.Filtered) > 0 {
			m.Cursor = 0
			m.Offset = 0
		}
	} else if m.Cursor < len(m.Filtered)-1 {
		m.Cursor++
		if m.Cursor >= m.Offset+visible {
			m.Offset = m.Cursor - visible + 1
		}
	}
	return m
}

// PageUp moves the cursor one page up.
func (m Model) PageUp() Model {
	if m.Cursor < 0 {
		// No selection is already the top.
		return m
	}
	m.Cursor -= m.MaxVisible()
	if m.Cursor < 0 {
		// Overshooting the top clears the selection.
		m.Cursor = -1
		m.Offset = 0
	} else {
		m.Offset = m.Cursor
	}
	return m
}

// PageDown moves the cursor one page down.
func (m Model) PageDown() Model {
	visible := m.MaxVisible()
	if m.Cursor < 0 {
		// Advancing from no selection lands on the first item.
		if len(m.Filtered) > 0 {
			m.Cursor = 0
			m.Offset = 0
		}
		return m
	}
	m.Cursor += visible
	if m.Cursor >= len(m.Filtered) {
		m.Cursor = len(m.Filtered) - 1
	}
	if m.Cursor >= m.Offset+visible {
		m.Offset = m.Cursor - visible + 1
	}
	return m
}
