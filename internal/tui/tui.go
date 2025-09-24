package tui

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/gemini"
	"github.com/rs/zerolog/log"
)

const idle = "idle"
const processing = "processing"
const blue = "#1D56F4"
const purple = "#7D56F4"

type TUIModel struct {
	geminiClient *gemini.Client
	status       string // processing, idle, etc
	message      string
	spinner      spinner.Model
}

func NewTUIModel(geminiClient *gemini.Client) *TUIModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &TUIModel{
		geminiClient: geminiClient,
		status:       idle,
		spinner:      s,
	}
}

func InitTUI(geminiClient *gemini.Client) error {
	model := NewTUIModel(geminiClient)
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

func (m TUIModel) Init() tea.Cmd {
	return nil
}

func (m TUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// its a key press
	case tea.KeyMsg:
		// cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.message != "" && m.status == idle {
				m.status = processing

				// run it async (make sure to create custom commands because bubbletea is not thread-safe)
				ctx := context.Background()
				for chunk := range m.geminiClient.Chat(ctx, m.message) {
					if chunk.Error != nil {
						log.Fatal().Err(chunk.Error).Msg("error to process answer")
					}

					if chunk.Done {
						break
					}

					c := styleTextColor(chunk.Text, blue)
					m.message = fmt.Sprintf("%s%s", m.message, c)
					m.status = idle
				}

				// start the spinner animation
				return m, m.spinner.Tick
			}
		case "backspace":
			if len(m.message) > 0 {
				m.message = m.message[:len(m.message)-1]
			}
		default:
			// Filter out special keys that shouldn't be added to message
			if len(msg.String()) == 1 {
				m.message = fmt.Sprintf("%s%s", m.message, msg.String())
			}
		}
	case spinner.TickMsg:
		if m.status == processing {
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m TUIModel) View() string {
	if m.status == processing {
		return fmt.Sprintf("Asking to gemini... %s", m.spinner.View())
	}

	return fmt.Sprintf("\nMessage is: %s", m.message)
}

func styleTextColor(text string, color string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

	return style.Render(text)
}
