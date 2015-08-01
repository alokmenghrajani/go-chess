package ai

import (
	//	"fmt"

	"github.com/alokmenghrajani/go-chess/core"
)

/**
 * Quickly hacked together AI.
 *
 * I could do a bunch of things better: better debugging, better
 * code re-use, avoid re-computing if a board position is in check or not, etc.
 */
func FindMate(board *core.Board, depth int) core.Moves {
	//	fmt.Printf("DEBUG: FindMate, depth=%d\nboard=\n%s\n", depth, board)
	if depth < 0 {
		return nil
	}
	possible_moves := board.ListMoves()
	if len(possible_moves) == 0 {
		// We are either mated or in stalemate. In either case we haven't found
		// a solution.
		// Display which case purly for debugging purpose
		if core.Is_in_check(board) {
			//			fmt.Printf("we are checkmated\n")
		} else {
			//			fmt.Printf("we are stalemated\n")
		}
		return nil
	}
	for _, m := range possible_moves {
		// if any move leads to a mate, we are good.
		// TODO: keep shortest chain?
		//		fmt.Printf("playing %s\n", m)
		b := m.Perform_move(board)
		more_moves := doesMate(b, depth-1)
		if more_moves != nil {
			//			fmt.Printf("we found a mate: %s", m)
			more_moves = append(more_moves, m)
			return more_moves
		}
	}
	//	fmt.Printf("we didn't find a mate, returning\n")
	// If there are no more moves, the current board is either a mate or stalemate
	return nil
}

func doesMate(board *core.Board, depth int) core.Moves {
	//	fmt.Printf("DEBUG: doesMate, depth=%d\nboard=\n%s\n", depth, board)
	if depth < 0 {
		return nil
	}

	// If every move leads to a mate, we are good. Return longest chain.
	longest := make(core.Moves, 0, 10)
	possible_moves := board.ListMoves()
	if len(possible_moves) == 0 {
		// We are either mated or in stalemate. We might have found a solution
		if core.Is_in_check(board) {
			//			fmt.Printf("solution found!\n")
			return longest
		} else {
			//			fmt.Printf("stalemate found.\n")
			return nil
		}
	}
	for _, m := range possible_moves {
		//		fmt.Printf("playing %s\n", m)
		b := m.Perform_move(board)
		more_moves := FindMate(b, depth-1)
		if more_moves == nil {
			return nil
		}
		more_moves = append(more_moves, m)
		if len(longest) < len(more_moves) {
			longest = more_moves
		}
	}
	return longest
}
