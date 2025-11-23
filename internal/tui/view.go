package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1).
		Align(lipgloss.Center)

	menuStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(20).
		Height(16).
		Padding(1)

	bucketStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(30).
		Height(16).
		Padding(1)

	objectStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(40).
		Height(16).
		Padding(1)

	selectedStyle = lipgloss.NewStyle().
		Reverse(true).
		Bold(true)

	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Align(lipgloss.Center)

	pathStyle = lipgloss.NewStyle()

	statusActiveStyle = lipgloss.NewStyle().
		Bold(true)

	statusLoadingStyle = lipgloss.NewStyle().
		Bold(true)

	keybindStyle = lipgloss.NewStyle().
		Bold(true)
)

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Resizing..."
	}

	title := titleStyle.Width(90).Render(fmt.Sprintf("BLOV Storage Manager [%s]", m.currentPath))

	var menuContent strings.Builder
	menuContent.WriteString(headerStyle.Width(18).Render("MENU") + "\n\n")
	for i, item := range menuItems {
		if i == m.menuIndex && m.activePane == 0 {
			menuContent.WriteString(selectedStyle.Render(fmt.Sprintf("> %s", item)) + "\n")
		} else {
			menuContent.WriteString(fmt.Sprintf("  %s\n", item))
		}
	}

	var bucketContent strings.Builder
	bucketContent.WriteString(headerStyle.Width(28).Render("BUCKETS") + "\n\n")
	if len(m.buckets) == 0 {
		bucketContent.WriteString("  No buckets found\n")
	} else {
		for i, bucket := range m.buckets {
			name := bucket.Name
			if len(name) > 22 {
				name = name[:19] + "..."
			}
			if i == m.bucketIndex && m.activePane == 1 {
				bucketContent.WriteString(selectedStyle.Render(fmt.Sprintf("> %s", name)) + "\n")
			} else {
				bucketContent.WriteString(fmt.Sprintf("  %s\n", name))
			}
		}
	}

	var objectContent strings.Builder
	objectContent.WriteString(headerStyle.Width(38).Render("OBJECTS") + "\n\n")
	if len(m.objects) == 0 {
		objectContent.WriteString("  No objects found\n")
	} else {
		for i, obj := range m.objects {
			name := obj.Name
			if len(name) > 30 {
				name = name[:27] + "..."
			}
			size := fmt.Sprintf("(%s)", obj.Size)
			if i == m.objectIndex && m.activePane == 2 {
				objText := fmt.Sprintf("> %s %s", name, size)
				objectContent.WriteString(selectedStyle.Render(objText) + "\n")
			} else {
				objectContent.WriteString(fmt.Sprintf("  %s %s\n", name, size))
			}
		}
	}

	menuBox := menuStyle.Render(menuContent.String())
	bucketBox := bucketStyle.Render(bucketContent.String())
	objectBox := objectStyle.Render(objectContent.String())

	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, menuBox, bucketBox, objectBox)

	var statusText string
	if m.loading {
		statusText = statusLoadingStyle.Render("Loading...")
	} else {
		statusText = statusActiveStyle.Render("Status: " + m.status)
	}

	providerInfo := pathStyle.Render(fmt.Sprintf("Provider: %s | Region: %s", m.provider, m.region))

	keybinds := lipgloss.JoinHorizontal(lipgloss.Left,
		keybindStyle.Render("[Tab]"),
		" Switch  ",
		keybindStyle.Render("[Enter]"),
		" Select  ",
		keybindStyle.Render("[q]"),
		" Quit")

	statusBar := lipgloss.JoinHorizontal(lipgloss.Left, statusText, "  |  ", providerInfo, "  |  ", keybinds)

	return lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		mainContent,
		"",
		statusBar,
	)
}