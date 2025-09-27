package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/gemini"
)

const idle = "idle"
const streaming = "streaming"
const answered = "answered"

// const cyan = "#00d7d7"
const cyanBlue = "#87d7ff"

var activeGeminiStreamChat <-chan gemini.GeminiChatStreamChunk

type chat struct {
	// question string
	answer    string
	textInput textinput.Model
}

type TUIModel struct {
	geminiClient *gemini.Client
	status       string // streaming, idle, done
	chats        []*chat
	ready        bool
	// components
	viewport viewport.Model
	spinner  spinner.Model
}

func NewTUIModel(geminiClient *gemini.Client) *TUIModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	ti := createTextInputModel()

	defaultChat := chat{
		// question: "",
		answer:    "",
		textInput: ti,
	}

	return &TUIModel{
		geminiClient: geminiClient,
		status:       idle,
		spinner:      s,
		chats: []*chat{
			&defaultChat,
		},
	}
}

func InitTUI(geminiClient *gemini.Client) error {
	model := NewTUIModel(geminiClient)
	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion(), tea.WithMouseAllMotion())
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

	currentChat := m.chats[len(m.chats)-1]

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ready {
			m.viewport = viewport.New(msg.Height, msg.Height)
			m.viewport.Style = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62"))
			m.viewport.MouseWheelEnabled = true
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height
		}

		currentChat.textInput.Width = msg.Width - 10

		return m, nil

	// its a key press
	case tea.KeyMsg:
		// cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "pgup":
			m.viewport.HalfPageUp()
			return m, nil
		case "pgdown":
			m.viewport.HalfPageDown()
			return m, nil
		case "home":
			m.viewport.GotoTop()
			return m, nil
		case "end":
			m.viewport.GotoBottom()
			return m, nil
		case tea.KeyEnter.String():
			v := strings.TrimSpace(currentChat.textInput.Value())
			if len(v) > 0 && m.status != streaming {
				// remove focus from the current chat
				currentChat = m.chats[len(m.chats)-1]
				currentChat.textInput.Blur()

				return m, m.startGeminiChatStreamCmd()
			}

			return m, nil
		}

		currentChat.textInput, cmd = currentChat.textInput.Update(msg)
		return m, cmd

	case tea.MouseMsg:
		switch msg.Button {
		case tea.MouseButtonWheelUp, tea.MouseButtonWheelDown:
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
		return m, nil

	case spinner.TickMsg:
		if m.status == streaming {
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	case geminiStreamingStarted:
		m.status = streaming
		return m, tea.Batch(m.spinner.Tick, m.readNextChunk())

	case geminiStreamingChunk:
		currentChat.answer = fmt.Sprintf("%s%s", currentChat.answer, msg.text)

		m.updateViewportContent()
		m.viewport.GotoBottom()

		return m, m.readNextChunk()

	case geminiStreamingDone:
		m.status = answered

		ti := createTextInputModel()
		// create/initialize a new chat for the next question
		m.chats = append(m.chats, &chat{answer: "", textInput: ti})
	}

	return m, cmd
}

type geminiStreamingStarted struct{}

func (m TUIModel) startGeminiChatStreamCmd() tea.Cmd {
	return func() tea.Msg {
		currentChat := m.chats[len(m.chats)-1]
		// run it async (make sure to create custom commands because bubbletea is not thread-safe)
		ctx := context.Background()
		activeGeminiStreamChat = m.geminiClient.Chat(ctx, currentChat.textInput.Value())

		return geminiStreamingStarted{}
	}
}

type geminiStreamingError struct{}
type geminiStreamingDone struct{}
type geminiStreamingChunk struct {
	text string
}

func (m TUIModel) readNextChunk() tea.Cmd {
	return func() tea.Msg {
		chunk := <-activeGeminiStreamChat
		if chunk.Error != nil {
			return geminiStreamingError{}
		}

		if chunk.Done {
			return geminiStreamingDone{}
		}

		return geminiStreamingChunk{text: chunk.Text}
	}
}

func (m TUIModel) View() string {
	if !m.ready {
		return "Initializing..."
	}

	m.updateViewportContent()

	return m.viewport.View()
}

func (m *TUIModel) updateViewportContent() {
	var msg strings.Builder

	for i, chat := range m.chats {
		isCurrentChat := false
		if i == len(m.chats)-1 {
			isCurrentChat = true
		}

		header := chat.textInput.View()
		body := fmt.Sprintf(`
%s`, chat.answer)

		out, err := glamour.Render(body, "dark")
		if err != nil {
			// msg = fmt.Sprintf("%s\n%s", msg, "failed to process this question.")
			msg.WriteString("Failed to process this question")
		}

		if m.status == streaming && isCurrentChat {
			out = fmt.Sprintf(`%s
%s processing...`, out, m.spinner.View())
		}

		msg.WriteString(fmt.Sprintf("%s%s", styleTextColor(header, cyanBlue), out))
	}

	m.viewport.SetContent(msg.String())
}

func styleTextColor(text string, color string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

	return style.Render(text)
}

func createTextInputModel() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "How can I help you today?"
	// initially we were using the tea.WindowSizeMsg but that is only dispatched at the start
	// we need to always set this because we are creating multiple inputs for each chat conversation
	ti.Width = 100
	ti.Focus()

	return ti
}
