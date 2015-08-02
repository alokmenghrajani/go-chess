package mate

import (
	"fmt"

	"github.com/alokmenghrajani/go-chess/core"
	"github.com/alokmenghrajani/go-chess/logger"
)

/**
 * Quickly hacked together AI.
 *
 * I could do a bunch of things better: better debugging, better
 * code re-use, avoid re-computing if a board position is in check or not, etc.
 */
func Find(logger *logger.Logger, board *core.Board, depth int) core.Moves {
	if depth < 0 {
		return nil
	}
	logger = logger.Push()
	possible_moves := board.ListMoves()
	for _, m := range possible_moves {
		// if any move leads to a mate, we are good.
		// TODO: keep shortest chain?
		b := m.Perform_move(board)

		more_moves := doesMate(logger, b, depth-1)
		if more_moves != nil {
			logger.Log(fmt.Sprintf("found mate: played %s, board=\n%s", m, b))
			more_moves = append(more_moves, m)
			return more_moves
		}
	}
	// If there are no more moves, the current board is either a mate or stalemate
	return nil
}

func doesMate(logger *logger.Logger, board *core.Board, depth int) core.Moves {
	if depth < 0 {
		return nil
	}
	logger = logger.Push()

	// If every move leads to a mate, we are good. Return longest chain.
	longest := make(core.Moves, 0, 10)
	possible_moves := board.ListMoves()
	if len(possible_moves) == 0 {
		// We are either mated or in stalemate. We might have found a solution
		if core.Is_in_check(board) {
			logger.Log(fmt.Sprintf("We are mated\n%s", board))
			return longest
		} else {
			return nil
		}
	}
	for _, m := range possible_moves {
		b := m.Perform_move(board)
		more_moves := Find(logger, b, depth-1)
		if more_moves == nil {
			return nil
		}
		more_moves = append(more_moves, m)
		if len(longest) < len(more_moves) {
			longest = more_moves
		}
	}
	logger.Log(fmt.Sprintf("all moves lead to mate, board=\n%s", board))
	return longest
}
