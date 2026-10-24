package memory

import "context"

// Role identifies a message participant.
type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)

// Message is a single conversation turn.
type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// TokenCounter estimates tokens for a message.
type TokenCounter interface {
	Count(m Message) int
}

// Summarizer compresses a slice of messages into one summary message.
type Summarizer interface {
	Summarize(ctx context.Context, msgs []Message) (Message, error)
}

// Manager maintains conversation context within a token budget.
type Manager struct {
	counter    TokenCounter
	summarizer Summarizer
	budget     int
	window     int
	messages   []Message
}

// NewManager creates a Manager. window is max messages before summarization; budget is token cap.
func NewManager(counter TokenCounter, summarizer Summarizer, budget, window int) *Manager {
	return &Manager{counter: counter, summarizer: summarizer, budget: budget, window: window}
}

// Add appends a message and trims context.
func (m *Manager) Add(ctx context.Context, msg Message) error {
	m.messages = append(m.messages, msg)
	return m.trim(ctx)
}

// Messages returns a copy of current context.
func (m *Manager) Messages() []Message {
	out := make([]Message, len(m.messages))
	copy(out, m.messages)
	return out
}

func (m *Manager) trim(ctx context.Context) error {
	// Sliding window: summarize oldest when exceeding window.
	for len(m.messages) > m.window && m.summarizer != nil {
		cut := len(m.messages) - m.window
		if cut < 1 {
			cut = 1
		}
		summary, err := m.summarizer.Summarize(ctx, m.messages[:cut])
		if err != nil {
			return err
		}
		m.messages = append([]Message{summary}, m.messages[cut:]...)
		if len(m.messages) <= m.window {
			break
		}
	}
	// Token budget enforcement.
	for m.budget > 0 && m.totalTokens() > m.budget && len(m.messages) > 1 {
		// drop oldest non-system
		idx := 0
		if m.messages[0].Role == RoleSystem {
			idx = 1
		}
		if idx >= len(m.messages) {
			break
		}
		m.messages = append(m.messages[:idx], m.messages[idx+1:]...)
	}
	return nil
}

func (m *Manager) totalTokens() int {
	total := 0
	for _, msg := range m.messages {
		total += m.counter.Count(msg)
	}
	return total
}
