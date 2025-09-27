package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
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
	// viewport     viewport.Model
	// components
	spinner spinner.Model
}

func NewTUIModel(geminiClient *gemini.Client) *TUIModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	// vp := viewport.New(80, 20)
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

	currentChat := m.chats[len(m.chats)-1]

	switch msg := msg.(type) {
	// its a key press
	case tea.KeyMsg:
		// cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case tea.KeyEnter.String():
			v := strings.TrimSpace(currentChat.textInput.Value())
			if len(v) > 0 && m.status != streaming {
				return m, m.startGeminiChatStreamCmd()
			}
		}

		var cmd tea.Cmd
		// newQuestion := fmt.Sprintf("%s%s", currentChat.textInput.View(), msg.String())
		// currentChat.textInput, cmd = currentChat.textInput.Update(newQuestion)
		currentChat.textInput, cmd = currentChat.textInput.Update(msg)

		return m, cmd

	// case tea.WindowSizeMsg:
	// 	currentChat.textInput.Width = msg.Width - 10

	case spinner.TickMsg:
		if m.status == streaming {
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	case geminiStreamingStarted:
		m.status = streaming
		return m, tea.Batch(m.spinner.Tick, m.readNextChunk())

	case geminiStreamingChunk:
		// c := styleTextColor(msg.text, blue)
		currentChat.answer = fmt.Sprintf("%s%s", currentChat.answer, msg.text)

		// update viewport
		// m.viewport.SetContent(currentChat.answer)

		return m, m.readNextChunk()

	case geminiStreamingDone:
		m.status = answered
		ti := createTextInputModel()
		// create/initialize a new chat for the next question
		m.chats = append(m.chats, &chat{answer: "", textInput: ti})
	}

	return m, nil
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
	// msg := fmt.Sprintf("%s\n", m.viewport.View())
	msg := ""

	for i, chat := range m.chats {
		isCurrentChat := false
		if i == len(m.chats)-1 {
			isCurrentChat = true
		}

		header := chat.textInput.View()

		if m.status == streaming && isCurrentChat {
			header = fmt.Sprintf(`%s
%s processing...`, header, m.spinner.View())
		}

		body := fmt.Sprintf(`
%s`, chat.answer)

		out, err := glamour.Render(body, "dark")
		if err != nil {
			msg = fmt.Sprintf("%s\n%s", msg, "failed to process this question.")
		}

		msg = fmt.Sprintf("%s\n%s%s", msg, styleTextColor(header, cyanBlue), out)
	}

	return msg
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
