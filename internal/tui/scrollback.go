package tui

// Scrollback holds lines with a viewport.
type Scrollback struct {
	Lines    []string
	Capacity int
	Offset   int
}

func NewScrollback(cap int) *Scrollback { return &Scrollback{Capacity: cap} }

func (s *Scrollback) Add(line string) {
	s.Lines = append(s.Lines, line)
	if s.Capacity > 0 && len(s.Lines) > s.Capacity {
		s.Lines = s.Lines[len(s.Lines)-s.Capacity:]
	}
}

func (s *Scrollback) View(height int) []string {
	if height <= 0 || len(s.Lines) == 0 {
		return nil
	}
	start := len(s.Lines) - height - s.Offset
	if start < 0 {
		start = 0
	}
	end := start + height
	if end > len(s.Lines) {
		end = len(s.Lines)
	}
	out := make([]string, end-start)
	copy(out, s.Lines[start:end])
	return out
}