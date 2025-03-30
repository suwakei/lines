package pathHandler

import (
	"testing"
)

func TestMakeIgnoreList(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected map[string][]string
		err      bool
	}{
		{
			input:    "",
			expected: map[string][]string{"file": nil, "dir": nil},
			err:      false,
		},
		{
			input:    []string{"../testdata/dummy/dummy.txt", "../testdata/dummy/dummyDir"},
			expected: map[string][]string{"file": {"../testdata/dummy/dummy.txt"}, "dir": {"../testdata/dummy/dummyDir"}},
			err:      false,
		},
		{
			input:    []string{},
			expected: nil,
			err:      true,
		},
		{
			input:    "invalid.txt",
			expected: nil,
			err:      true,
		},
	}

	for _, test := range tests {
		switch v := test.input.(type) {
		case string:
			result, err := MakeIgnoreList(v)
			if (err != nil) != test.err {
				t.Errorf("expected error: %v, got: %v", test.err, err)
			}
			if !equal(result, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, result)
			}

		case []string:
			result, err := MakeIgnoreList(v)
			if (err != nil) != test.err {
				t.Errorf("expected error: %v, got: %v", test.err, err)
			}
			if !equal(result, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, result)
			}
		}
	}
}

func equal(a, b map[string][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for key, valA := range a {
		valB, exists := b[key]
		if !exists || len(valA) != len(valB) {
			return false
		}

		for i, v := range valA {
			if v != valB[i] {
				return false
			}
		}
	}
	return true
}
