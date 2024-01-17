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
				{"X", "", ""},
				{"", "", ""},
				{"", "", ""},
			},
			expectedResult: &[2]int{2, 2},
		},
	}

	for _, tc := range testCases {
		result, err := Minmax(tc.board)

		t.Logf("board %v", tc.board)
		assert.Nil(t, err) 
		assert.Equal(t, tc.expectedResult, result) 
	}
}
