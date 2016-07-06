package ports

import (
	"bytes"
	"testing"
)

func TestLaunch(t *testing.T) {
	table := []struct {
		program  string
		params   []string
		expected []byte
	}{
		{"echo", []string{"hello world"}, []byte("hello world\n")},
	}

	for _, test := range table {
		result, err := Launch(test.program, test.params)
		if err != nil {
			t.Errorf("launch returned error: %v\n", err)
		}
		if !bytes.Equal(test.expected, result) {
			t.Errorf("Expected %v, got %v\n", test.expected, result)
		}
	}
}
