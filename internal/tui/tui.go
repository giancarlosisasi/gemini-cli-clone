package tui

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/gemini"
)

const idle = "idle"
const streaming = "streaming"
const answered = "answered"
const blue = "#1D56F4"
const purple = "#7D56F4"

var activeGeminiStreamChat <-chan gemini.GeminiChatStreamChunk

type chat struct {
	question string
	answer   string
}

type TUIModel struct {
	geminiClient *gemini.Client
	status       string // streaming, idle, done
	spinner      spinner.Model
	chats        []*chat
}

func NewTUIModel(geminiClient *gemini.Client) *TUIModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	defaultChat := chat{question: "", answer: ""}
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
		case "enter":
			if currentChat.question != "" && m.status != streaming {
				return m, m.startGeminiChatStreamCmd()
			}
		case "backspace":
			if len(currentChat.question) > 0 {
				currentChat.question = currentChat.question[:len(currentChat.question)-1]
			}
		default:
			// Filter out special keys that shouldn't be added to message
			if len(msg.String()) == 1 {
				currentChat.question = fmt.Sprintf("%s%s", currentChat.question, msg.String())
			}
		}

	case spinner.TickMsg:
		if m.status == streaming {
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	case geminiStreamingStarted:
		m.status = streaming
		return m, tea.Batch(m.spinner.Tick, m.readNextChunk())

	case geminiStreamingChunk:
		c := styleTextColor(msg.text, blue)
		currentChat.answer = fmt.Sprintf("%s%s", currentChat.answer, c)
		return m, m.readNextChunk()

	case geminiStreamingDone:
		m.status = answered
		// create/initialize a new chat for the next question
		m.chats = append(m.chats, &chat{question: "", answer: ""})
	}

	return m, nil
}

type geminiStreamingStarted struct{}

func (m TUIModel) startGeminiChatStreamCmd() tea.Cmd {
	return func() tea.Msg {
		currentChat := m.chats[len(m.chats)-1]
		// run it async (make sure to create custom commands because bubbletea is not thread-safe)
		ctx := context.Background()
		activeGeminiStreamChat = m.geminiClient.Chat(ctx, currentChat.question)

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
	msg := ""

	for _, chat := range m.chats {
		header := fmt.Sprintf(`-----------------------------------------
> %s
-----------------------------------------`, chat.question)

		if m.status == streaming {
			header = fmt.Sprintf(`%s
%s processing...`, header, m.spinner.View())
		}

		body := fmt.Sprintf(`
%s
`, chat.answer)

		msg = fmt.Sprintf("%s%s%s", msg, header, body)
	}

	return msg
}

func styleTextColor(text string, color string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

	return style.Render(text)
}
