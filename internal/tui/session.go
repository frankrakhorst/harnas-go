package tui

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Message is a chat message.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Session holds chat history.
type Session struct {
	Messages []Message `json:"messages"`
}

// LoadSession reads session from path (returns empty if missing).
func LoadSession(path string) (Session, error) {
	var s Session
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return s, nil
		}
		return s, err
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return s, err
	}
	return s, nil
}

// SaveSession writes session to path.
func SaveSession(path string, s Session) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}
