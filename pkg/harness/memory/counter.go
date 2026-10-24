package memory

// WordCounter counts tokens by whitespace words (simple heuristic).
type WordCounter struct{}

func (WordCounter) Count(m Message) int {
	n := 0
	for _, c := range m.Content {
		if c == ' ' || c == '\n' || c == '\t' {
			n++
		}
	}
	if len(m.Content) > 0 {
		n++
	}
	return n
}
