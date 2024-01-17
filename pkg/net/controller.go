package net

import (
	"tictactoe/pkg/tictactoe"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func cpuPlay(game *Game) error {
	if game.Terminal {
		return nil
	}

	move, err := tictactoe.Minmax(game.Board)
	if err != nil {
		return errors.WithStack(err)
	}

	return playerPlay(game, *move)
}

func playerPlay(game *Game, move [2]int) error {
	var err error

	if game.Terminal {
		return nil
	}

	game.Board, err = tictactoe.Result(globalGame.Board, move)
	if err != nil {
		return errors.WithStack(err)
	}

	logrus.Infof("board %v", globalGame.Board)

	game.Terminal = tictactoe.Terminal(game.Board)
	if game.Terminal {
		game.Terminal = true
		game.Winner = tictactoe.Winner(game.Board)
	}
	return nil
}
