package tui

import (
	"context"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hmans/beans/internal/ui"
	"github.com/hmans/beans/pkg/bean"
	"github.com/hmans/beans/pkg/config"
)

// beanCreatedMsg is sent when a new bean is created
type beanCreatedMsg struct {
	title    string
	parentID string
}

// closeCreateModalMsg is sent when the create modal is cancelled
type closeCreateModalMsg struct{}

// openCreateModalMsg requests opening the create modal
type openCreateModalMsg struct {
	parentID string
}

// beanGetter looks up a bean by ID for parent display
type beanGetter interface {
	Bean(ctx context.Context, id string) (*bean.Bean, error)
}

type createFocus int

const (
	focusTitle createFocus = iota
	focusParent
)

// createModalModel is the model for the create bean modal
type createModalModel struct {
	textInput textinput.Model
	lookup    beanGetter
	config    *config.Config
	parentID  string
	focus     createFocus
	width     int
	height    int
}

func newCreateModalModel(lookup beanGetter, cfg *config.Config, parentID string, width, height int) createModalModel {
	ti := textinput.New()
	ti.Placeholder = "Enter bean title..."
	ti.CharLimit = 200
	ti.Width = 50
	ti.Focus()
	ti.PromptStyle = lipgloss.NewStyle().Foreground(ui.ColorPrimary)
	ti.TextStyle = lipgloss.NewStyle()
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(ui.ColorMuted)
	ti.Prompt = ""

	return createModalModel{
		textInput: ti,
		lookup:    lookup,
		config:    cfg,
		parentID:  parentID,
		focus:     focusTitle,
		width:     width,
		height:    height,
	}
}

func (m createModalModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *createModalModel) toggleFocus() {
	if m.focus == focusTitle {
		m.focus = focusParent
		m.textInput.Blur()
		return
	}
	m.focus = focusTitle
	m.textInput.Focus()
}

func (m createModalModel) openParentPicker() tea.Msg {
	return openParentPickerMsg{
		beanTitle:     m.textInput.Value(),
		beanTypes:     []string{m.config.GetDefaultType()},
		currentParent: m.parentID,
	}
}

func (m createModalModel) Update(msg tea.Msg) (createModalModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg {
				return closeCreateModalMsg{}
			}

		case "tab", "shift+tab":
			m.toggleFocus()
			return m, nil

		case "enter":
			if m.focus == focusParent {
				return m, func() tea.Msg {
					return m.openParentPicker()
				}
			}
			title := m.textInput.Value()
			if title == "" {
				return m, func() tea.Msg {
					return closeCreateModalMsg{}
				}
			}
			parentID := m.parentID
			return m, func() tea.Msg {
				return beanCreatedMsg{title: title, parentID: parentID}
			}
		}
	}

	if m.focus != focusTitle {
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m createModalModel) parentLabel() string {
	if m.parentID == "" {
		return ""
	}
	if m.lookup != nil {
		parent, err := m.lookup.Bean(context.Background(), m.parentID)
		if err == nil && parent != nil {
			return parent.Type + " " + parent.Title
		}
	}
	return m.parentID
}

func (m createModalModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	modalWidth := max(40, min(60, m.width*50/100))

	header := lipgloss.NewStyle().Bold(true).Render("Create New Bean")

	inputBorder := ui.ColorMuted
	if m.focus == focusTitle {
		inputBorder = ui.ColorPrimary
	}
	inputBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(inputBorder).
		Padding(0, 1).
		Width(modalWidth - 6).
		Render(m.textInput.View())

	cursor := "  "
	if m.focus == focusParent {
		cursor = lipgloss.NewStyle().Foreground(ui.ColorPrimary).Bold(true).Render("▌") + " "
	}
	parentValue := m.parentLabel()
	if parentValue == "" {
		parentValue = ui.Muted.Render("(none)")
	}
	parentLine := cursor + ui.Muted.Render("Parent") + "  " + parentValue
	parentBlock := lipgloss.NewStyle().
		Width(modalWidth - 6).
		MaxHeight(2).
		Render(parentLine)

	var help string
	if m.focus == focusTitle {
		help = helpKeyStyle.Render("enter") + " " + helpStyle.Render("create") + "  " +
			helpKeyStyle.Render("tab") + " " + helpStyle.Render("parent") + "  " +
			helpKeyStyle.Render("esc") + " " + helpStyle.Render("cancel")
	} else {
		help = helpKeyStyle.Render("enter") + " " + helpStyle.Render("change") + "  " +
			helpKeyStyle.Render("tab") + " " + helpStyle.Render("title") + "  " +
			helpKeyStyle.Render("esc") + " " + helpStyle.Render("cancel")
	}

	content := header + "\n\n" + inputBox + "\n\n" + parentBlock + "\n\n" + help

	border := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ui.ColorPrimary).
		Padding(1, 2).
		Width(modalWidth)

	return border.Render(content)
}

// ModalView returns the modal rendered as a centered overlay
func (m createModalModel) ModalView(bgView string, fullWidth, fullHeight int) string {
	modal := m.View()
	return overlayModal(bgView, modal, fullWidth, fullHeight)
}
