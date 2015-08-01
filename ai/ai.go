package ai

import (
	"fmt"

	"github.com/alokmenghrajani/go-chess/core"
)

type Logger struct {
	enabled bool
	offset  int
}

func NewLogger(enabled bool) *Logger {
	l := new(Logger)
	l.enabled = enabled
	l.offset = 0
	return l
}

func (logger *Logger) push() *Logger {
	l := *logger
	l.offset++
	return &l
}

func (logger *Logger) log(msg string) {
	if !logger.enabled {
		return
	}
	for i := 0; i < logger.offset; i++ {
		fmt.Printf("  ")
	}
	for _, c := range msg {
		fmt.Printf("%c", c)
		if c == '\n' {
			for i := 0; i < logger.offset; i++ {
				fmt.Printf("  ")
			}
		}
	}
	fmt.Printf("\n")
}

/**
 * Quickly hacked together AI.
 *
 * I could do a bunch of things better: better debugging, better
 * code re-use, avoid re-computing if a board position is in check or not, etc.
 */
func FindMate(logger *Logger, board *core.Board, depth int) core.Moves {
	if depth < 0 {
		return nil
	}
	logger = logger.push()
	//	logger.log(fmt.Sprintf("entering FindMate, depth=%d, board=\n%s", depth, board))
	possible_moves := board.ListMoves()
	if len(possible_moves) == 0 {
		// We are either mated or in stalemate. In either case we haven't found
		// a solution.
		// Display which case purly for debugging purpose
		if core.Is_in_check(board) {
			//			logger.log("we are checkmated\n")
		} else {
			//			logger.log("we are stalemated\n")
		}
		return nil
	}
	for _, m := range possible_moves {
		// if any move leads to a mate, we are good.
		// TODO: keep shortest chain?
		b := m.Perform_move(board)

		more_moves := doesMate(logger, b, depth-1)
		if more_moves != nil {
			logger.log(fmt.Sprintf("found mate: played %s, board=\n%s", m, b))
			more_moves = append(more_moves, m)
			return more_moves
		}
	}
	//	logger.log("we did not find a mate")
	// If there are no more moves, the current board is either a mate or stalemate
	return nil
}

func doesMate(logger *Logger, board *core.Board, depth int) core.Moves {
	if depth < 0 {
		return nil
	}
	logger = logger.push()
	//	logger.log(fmt.Sprintf("entering doesMate, depth=%d, board=\n%s", depth, board))

	// If every move leads to a mate, we are good. Return longest chain.
	longest := make(core.Moves, 0, 10)
	possible_moves := board.ListMoves()
	if len(possible_moves) == 0 {
		// We are either mated or in stalemate. We might have found a solution
		if core.Is_in_check(board) {
			logger.log(fmt.Sprintf("We are mated\n%s", board))
			return longest
		} else {
			//			logger.log("we are stalemate")
			return nil
		}
	}
	for _, m := range possible_moves {
		b := m.Perform_move(board)
		more_moves := FindMate(logger, b, depth-1)
		if more_moves == nil {
			//			logger.log("all moves don't mate, returning")
			return nil
		}
		more_moves = append(more_moves, m)
		if len(longest) < len(more_moves) {
			longest = more_moves
		}
	}
	logger.log(fmt.Sprintf("all moves lead to mate, board=\n%s", board))
	return longest
}
