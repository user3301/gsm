package gsm_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/user3301/gsm"
	"testing"
)

func TestHammingDistanceSuccess(t *testing.T) {
	t.Parallel()
	testCases := []struct{
		testName string
		a,b string
		expectedByteCount int
		expectedBitCount int
	} {
		{
			testName: "Two Equal Strings",
			a: "cyberpunk2077",
			b: "cyberpunk2077",
			expectedByteCount: 0,
			expectedBitCount: 0,
		},
		{
			testName: "Two Different String With Equal Length",
			a: "cyberpunk2077",
			b: "cyberpunk2088",
			expectedByteCount: 2,
			expectedBitCount: 8,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			byteCount, err := gsm.HammingDistanceByteCount(tc.a, tc.b)
			bitCount, err := gsm.HammingDistanceBitCount(tc.a, tc.b)
			assert.NoError(t, err)
			assert.Equal(t, byteCount, tc.expectedByteCount)
			assert.Equal(t, bitCount, tc.expectedBitCount)
		})
	}
}

func TestHammingDistanceError(t *testing.T) {
	t.Parallel()
	a, b := "length12", "length8"
	_, err := gsm.HammingDistanceBitCount(a, b)
	assert.EqualError(t, err,"two string has different length")
	_, err = gsm.HammingDistanceByteCount(a, b)
	assert.EqualError(t, err, "two string has different length")
}