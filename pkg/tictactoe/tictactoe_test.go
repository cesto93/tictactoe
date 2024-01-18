package tictactoe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinmax(t *testing.T) {
	testCases := []struct {
		board          [3][3]string
		expectedResult *[2]int
	}{
		{
			board: [3][3]string{
				{"X", "X", ""},
				{"O", "O", ""},
				{"", "", ""},
			},
			expectedResult: &[2]int{0, 2},
		},
	}

	for _, tc := range testCases {
		result, err := Minmax(tc.board)

		assert.Nil(t, err)
		assert.Equal(t, tc.expectedResult, result)
	}
}

func TestPlayer(t *testing.T) {
	testCases := []struct {
		board          [3][3]string
		expectedResult string
	}{
		{
			board: [3][3]string{
				{"X", "X", ""},
				{"O", "O", ""},
				{"", "", ""},
			},
			expectedResult: X,
		},
		{
			board: [3][3]string{
				{"X", "X", ""},
				{"O", "", ""},
				{"", "", ""},
			},
			expectedResult: O,
		},
	}

	for _, tc := range testCases {
		result := player(tc.board)

		assert.Equal(t, tc.expectedResult, result)
	}
}

func TestUtility(t *testing.T) {
	testCases := []struct {
		board          [3][3]string
		expectedResult int
	}{
		{
			board: [3][3]string{
				{"X", "X", "X"},
				{"O", "O", ""},
				{"", "", ""},
			},
			expectedResult: 1,
		},
		{
			board: [3][3]string{
				{"X", "X", ""},
				{"O", "O", "O"},
				{"X", "", ""},
			},
			expectedResult: -1,
		},
		{
			board: [3][3]string{
				{"X", "X", "O"},
				{"O", "O", "X"},
				{"X", "O", "X"},
			},
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		result := utility(tc.board)

		assert.Equal(t, tc.expectedResult, result)
	}
}
