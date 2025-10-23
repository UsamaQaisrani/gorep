package core
import (
	"testing"
	"strings"
)

type mockReadCloser struct {
	*strings.Reader	
}

func (m *mockReadCloser) Close() error {
	return nil	
}

func TestReadLines(t *testing.T) {
	const input = "First line.\nSecond line."
	want := []string{"First line.", "Second line."}

	mockFile := &mockReadCloser{Reader: strings.NewReader(input)}
	lineChan := ReadLines(mockFile)

	got := []string{}
	
	for line := range lineChan {
		got = append(got, line)
	}

	if len(got) != len(want) {
		t.Fatalf("Got %d lines, want %d lines. Got: %v", len(got), len(want), got)
	}
}
