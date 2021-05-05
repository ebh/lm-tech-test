package lm_tech_test_test

import (
	lm_tech_test "lm-tech-test"
	"testing"

	"github.com/stretchr/testify/assert"
)

const longTextToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"

func TestSearchCaseInsensitive(t *testing.T) {
	tcs := []struct {
		textToSearch string
		subtext      string
		expected     []int
	}{
		// These test cases are from the problem description
		{
			textToSearch: longTextToSearch,
			subtext:      "Peter",
			expected:     []int{1, 20, 75},
		},
		{
			textToSearch: longTextToSearch,
			subtext:      "peter",
			expected:     []int{1, 20, 75},
		},
		{
			textToSearch: longTextToSearch,
			subtext:      "pick",
			expected:     []int{30, 58},
		},
		{
			textToSearch: longTextToSearch,
			subtext:      "pi",
			expected:     []int{30, 37, 43, 51, 58},
		},
		{
			textToSearch: longTextToSearch,
			subtext:      "z",
			expected:     nil,
		},
		{
			textToSearch: longTextToSearch,
			subtext:      "Peterz",
			expected:     nil,
		},

		// These test cases are ones I've added
		// Note: The specs were unclear what is the desired outcome
		// for some of these case. So I made my own decision and
		// documented using tests
		{
			textToSearch: "",
			subtext:      "",
			expected:     nil,
		},
		{
			textToSearch: "",
			subtext:      "a",
			expected:     nil,
		},
		{
			textToSearch: "a",
			subtext:      "",
			expected:     nil,
		},
		{
			textToSearch: "a",
			subtext:      "a",
			expected:     []int{1},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.subtext, func(t *testing.T) {
			assert.NotPanics(t, func() {
				r := lm_tech_test.SearchCaseInsensitive(tc.textToSearch, tc.subtext)

				assert.Equal(t, tc.expected, r)
			})
		})
	}
}
