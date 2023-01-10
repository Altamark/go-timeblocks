package timeblocks_test

import (
	"testing"
	"time"

	timeblocks "github.com/altamark/go-timeblocks/timeblocks"
)

var (
	locNewYork, _ = time.LoadLocation("America/New_York")
	// locParis, _   = time.LoadLocation("Europe/Paris")
)

func TestBlockIsOverlap(t *testing.T) {
	cases := []struct {
		block1   timeblocks.Block
		block2   timeblocks.Block
		expected bool
		actual   bool
	}{
		// Case 1
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
			},
			expected: false,
		},
		// Case 2
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			expected: false,
		},
		// Case 3
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			expected: true,
		},
		// Case 4
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 3, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 4, 1, 0, 0, locNewYork),
			},
			expected: false,
		},
		// Case 5
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 3, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 4, 1, 0, 0, locNewYork),
			},
			expected: true,
		},
		// Case 6
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 3, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 4, 1, 0, 0, locNewYork),
			},
			expected: true,
		},
		// Case 7
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 3, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			expected: false,
		},
		// Case 8
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 3, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			expected: false,
		},
		// Case 9
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 3, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			expected: true,
		},
		// Case 10
		{
			block1: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 3, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 2, 1, 0, 0, locNewYork),
			},
			block2: timeblocks.Block{
				Start: time.Date(2023, 1, 1, 1, 1, 0, 0, locNewYork),
				End:   time.Date(2023, 1, 1, 3, 1, 0, 0, locNewYork),
			},
			expected: true,
		},
	}

	for i, testCase := range cases {
		if testCase.actual = testCase.block1.IsOverlap(&testCase.block2); testCase.actual != testCase.expected {
			t.Errorf("Test Case %d: expected: %v, got: %v", i+1, testCase.expected, testCase.actual)
			return
		}
	}
}
