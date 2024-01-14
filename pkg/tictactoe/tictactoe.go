package tictactoe

import (
	"errors"
)

const (
	X = "X"
	O = "O"
)

func Init_state() [3][3]string {
	return [3][3]string{}
}

func player(board [3][3]string) string{
	moves := 0
	for i:= 0; i < 3; i++ {
		for j:=0; j < 3; j++ {
			if board[i][j] != "" {
				moves++
			}
		}
	}
	if moves % 2 == 0 {
		return X
	} else {
		return O
	}
}

func actions(board [3][3]string) [][2]int{
	actions := make([][2]int, 0)
	for i:= 0; i < 3; i++ {
		for j:=0; j < 3; j++ {
			if board[i][j] != "" {
				actions = append(actions, [2]int{i,j})
			}
		}
	}
	return actions
}

func Result(board [3][3]string, action [2]int) ([3][3]string, error){
	x, y := action[0], action[1]
	if board[x][y] != "" {
		return [3][3]string{}, errors.New("illegal move")
	}
	board[x][y] = player(board)
	return board, nil
}

func winner(board [3][3]string) string {
	for i:= 0; i < 3; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
				return board[i][0]
		}
		
		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
				return board[0][i]
		}

		if board[1][1] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
				return board[0][0]
		}

		if board[1][1] != "" && board[0][2] == board[1][1] && board[2][0] == board[0][2] {
				return board[0][2]
		}
	}
	return ""
}

func terminal(board [3][3]string) bool{
	if winner(board) != "" {
		return true
	}

	return len(actions(board)) == 0
}

func utility(board [3][3]string) int {
	switch winner(board) {
		case X: return 1
		case O: return -1
		default: return 0
	}
}

func maxval(board [3][3]string) (int, error) {
	v := -2
	if terminal(board) {
		return utility(board), nil
	}

	for _, action := range actions(board) {
		res, err := Result(board,action)
		if err != nil {
			return 0, err
		}

		mval, err := minval(res)
		if err != nil {
			return 0, err
		}

		v = max(v, mval)
	}
	return v, nil
}

func minval(board [3][3]string) (int, error) {
	v := +2
	if terminal(board) {
		return utility(board), nil
	}
	for _, action := range actions(board) {
		res, err := Result(board,action)
		if err != nil {
			return 0, err
		}

		mval, err := maxval(res)
		if err != nil {
			return 0, err
		}

		v = max(v, mval)

	}
	return v, nil
}

func Minmax(board [3][3]string) (*[2]int, error) {
	var move [2]int
	var eval func([3][3]string)(int, error)
	var v int
	var isMin bool

	if terminal(board) {
		return nil, nil
	}
	
	if player(board) == X {
		v = -2
		eval = minval
		isMin = true
	} else {
		v = 2
		eval = maxval
		isMin = false
	}

	for _, action:= range actions(board) {
		res, err := Result(board,action)
		if err != nil {
			return nil, err
		}

		n, err := eval(res)
		if err != nil {
			return nil, err
		}

		if isMin {
			if n < v {
				v = n
				move = action
			}
		} else {
			if n > v {
				v = n
				move = action
			}
		}

	}
	return &move, nil
}
