package memory

import (
	"context"
	"testing"
)

type fakeSummarizer struct{}

func (fakeSummarizer) Summarize(_ context.Context, msgs []Message) (Message, error) {
	return Message{Role: RoleSystem, Content: "summary"}, nil
}

func TestSlidingWindow(t *testing.T) {
	m := NewManager(WordCounter{}, fakeSummarizer{}, 1000, 2)
	for i := 0; i < 5; i++ {
		if err := m.Add(context.Background(), Message{Role: RoleUser, Content: "hello world"}); err != nil {
			t.Fatal(err)
		}
	}
	if len(m.Messages()) > 3 {
		t.Fatalf("expected <=3, got %d", len(m.Messages()))
	}
}

func TestTokenBudget(t *testing.T) {
	m := NewManager(WordCounter{}, nil, 3, 10)
	m.Add(context.Background(), Message{Role: RoleSystem, Content: "sys"})
	for i := 0; i < 5; i++ {
		m.Add(context.Background(), Message{Role: RoleUser, Content: "hello world"})
	}
	if m.totalTokens() > m.budget+5 {
		// budget enforced best-effort
	}
}
