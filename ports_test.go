package ports

import (
	"bytes"
	"testing"
)

func TestLaunch(t *testing.T) {
	table := []struct {
		program  string
		expected []byte
	}{
		{`echo "hello world"`, []byte("hello world")},
	}

	for _, test := range table {
		ch := make(chan []byte)
		err := Launch(test.program, ch)
		if err != nil {
			t.Errorf("launch returned error: %v\n", err)
		}
		ret := <-ch
		if !bytes.Equal(test.expected, ret) {
			t.Errorf("Expected %v, got %v\n", test.expected, ret)
		}
	}
}
