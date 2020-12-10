package gsm_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/user3301/gsm"
	"testing"
)

func TestLevenshtein(t *testing.T) {
	t.Parallel()
	testCases := []struct{
		testName string
		a, b string
		expected int
	} {
		{
			testName: "Two Equal Strings",
			a: "hello world",
			b: "hello world",
			expected: 0,
		},
		{
			testName: "One Of Them Is Empty",
			a: "",
			b:"hello world",
			expected: 11,
		},
		{
			testName: "Two Equal Length Strings Have Different Character At Different Position",
			a: "a quick brown fox jumps over the lazy dog",
			b: "a quick brown cat jumps over the lazy dog",
			expected: 3,
		},
		{
			testName: "Two Difference Strings With Different Length And Characters",
			a: "no lollygagging",
			b: "no lollipop",
			expected: 8,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			gotDistance := gsm.LevenshteinDistanceByteCount(tc.a, tc.b)
			assert.Equal(t, tc.expected, gotDistance)
		})
	}
}
