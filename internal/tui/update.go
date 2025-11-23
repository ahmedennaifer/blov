package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.cancel()
			return m, tea.Quit

		case "tab":
			m.activePane = (m.activePane + 1) % 3
			return m, nil

		case "up", "k":
			return m.handleUp(), nil

		case "down", "j":
			return m.handleDown(), nil

		case "enter":
			return m.handleEnter()

		case "esc":
			m.activePane = 0
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case BucketsLoadedMsg:
		m.buckets = msg.buckets
		m.loading = false
		m.status = "Connected"
		return m, nil

	case ObjectsLoadedMsg:
		m.objects = msg.objects
		m.loading = false
		m.status = fmt.Sprintf("Loaded %d objects", len(msg.objects))
		return m, nil

	case ErrorMsg:
		m.loading = false
		m.status = "Error: " + msg.err.Error()
		return m, nil
	}

	return m, nil
}

func (m Model) handleUp() Model {
	switch m.activePane {
	case 0:
		if m.menuIndex > 0 {
			m.menuIndex--
		}
	case 1:
		if m.bucketIndex > 0 {
			m.bucketIndex--
		}
	case 2:
		if m.objectIndex > 0 {
			m.objectIndex--
		}
	}
	return m
}

func (m Model) handleDown() Model {
	switch m.activePane {
	case 0:
		if m.menuIndex < len(menuItems)-1 {
			m.menuIndex++
		}
	case 1:
		if m.bucketIndex < len(m.buckets)-1 {
			m.bucketIndex++
		}
	case 2:
		if m.objectIndex < len(m.objects)-1 {
			m.objectIndex++
		}
	}
	return m
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	switch m.activePane {
	case 0:
		switch m.menuIndex {
		case 0:
			m.activePane = 1
			return m, nil
		case 1:
			m.status = "Login functionality not implemented"
			return m, nil
		case 2:
			m.status = "Config functionality not implemented"
			return m, nil
		case 3:
			m.status = "Dashboard functionality not implemented"
			return m, nil
		}

	case 1:
		if m.bucketIndex < len(m.buckets) {
			bucketName := m.buckets[m.bucketIndex].Name
			m.currentPath = "/" + bucketName
			m.activePane = 2
			m.loading = true
			m.status = "Loading objects..."
			return m, m.loadObjects(bucketName)
		}

	case 2:
		if m.objectIndex < len(m.objects) {
			obj := m.objects[m.objectIndex]
			m.status = "Selected: " + obj.Name
			return m, nil
		}
	}

	return m, nil
}

