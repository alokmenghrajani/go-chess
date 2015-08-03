package selfmate

import (
	"fmt"

	"github.com/alokmenghrajani/go-chess/core"
	"github.com/alokmenghrajani/go-chess/logger"
)

func Find(logger *logger.Logger, board *core.Board, depth int) core.Moves {
	if depth < 0 {
		return nil
	}
	logger = logger.Push()
	possible_moves := board.ListMoves()

	logger.Log(fmt.Sprintf("Entering Find, depth=%d, board=\n%s, possible=%d\n", depth, board, len(possible_moves)))

	// the goal is to force ourselves into a mating position
	if len(possible_moves) == 0 {
		if core.Is_in_check(board) {
			logger.Log(fmt.Sprintf("We are mated\n%s", board))
			return make(core.Moves, 0, 10)
		} else {
			return nil
		}
	}
	for _, m := range possible_moves {
		// if all moves leads to a selfmate, we are good.
		// todo: keep shortest chain
		b := m.Perform_move(board)
		logger.Log(fmt.Sprintf("played %s, board=\n%s\n", m, b))

		more_moves := doesSelfMate(logger, b, depth-1)
		if more_moves != nil {
			logger.Log(fmt.Sprintf("all moves lead to selfmate, board=\n%s", board))
			more_moves = append(more_moves, m)
			return more_moves
		}
	}
	return nil
}

func doesSelfMate(logger *logger.Logger, board *core.Board, depth int) core.Moves {
	if depth < 0 {
		return nil
	}
	logger = logger.Push()
	var longest core.Moves = nil
	// If every move leads to a selfmate, we are good. Return longest chain.
	possible_moves := board.ListMoves()
	logger.Log(fmt.Sprintf("Entering doesSelfMate, depth=%d, board=\n%s\npossible=%d\n", depth, board, len(possible_moves)))
	for _, m := range possible_moves {
		b := m.Perform_move(board)
		logger.Log(fmt.Sprintf("played %s, board=\n%s", m, b))

		more_moves := Find(logger, b, depth-1)
		if more_moves == nil {
			return nil
		}
		more_moves = append(more_moves, m)
		if longest == nil || (len(longest) < len(more_moves)) {
			longest = more_moves
		}
	}
	logger.Log(fmt.Sprintf("all moves lead to selfmate, board=\n%s", board))
	return longest
}
