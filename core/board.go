package core

import (
	"errors"
	"fmt"
	"strings"
)

/**
 * Board related code.
 *
 * For a start, we represent a baord as an array of pieces. An alternative
 * would be to use 64-bit vectors for each type of piece (12x64 bits) and
 * leverage some smart binary operations.
 *
 */

type color int

const (
	White color = 0
	Black color = 1
)

type square struct {
	piece piece
	color color
}

type Board struct {
	board          [8][8]square
	to_play        color
	en_passant     int /* if not -2, en_passant allowed on this column */
	has_moved_king [2]bool
	has_moved_rook [2][2]bool
}

func (board *Board) ListMoves() Moves {
	r := make(Moves, 0, 35)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			point := xy{i, j}
			if !is_empty(board, point) && !is_opponent(board, point) {
				r2 := board.board[i][j].piece.list_moves(board, point)
				r = append(r, r2...)
			}
		}
	}
	return r
}

func Parse(s string, color color) (*Board, error) {
	b := new(Board)
	b.to_play = color
	b.en_passant = -2
	b.has_moved_king[0] = true
	b.has_moved_king[1] = true
	b.has_moved_rook[0][0] = true
	b.has_moved_rook[0][1] = true
	b.has_moved_rook[1][0] = true
	b.has_moved_rook[1][1] = true

	board_offset := 0
	for i := 0; i < len(s); i++ {
		square := byteToSquare(s[i])
		if square == nil {
			continue
		}
		b.board[board_offset%8][7-board_offset/8] = *square
		board_offset++
	}
	if board_offset != 64 {
		return nil, errors.New(fmt.Sprintf("Parse: not enough data? %d", board_offset))
	}
	return b, nil
}

func byteToSquare(c byte) *square {
	ic := ToLower(c)
	var p piece
	switch ic {
	case '.':
		return &square{nil, White}
	case 'p':
		p = new(pawn)
	case 'r':
		p = new(rook)
	case 'n':
		p = new(knight)
	case 'b':
		p = new(bishop)
	case 'k':
		p = new(king)
	case 'q':
		p = new(queen)
	default:
		return nil
	}
	if ic == c {
		return &square{p, Black}
	}
	return &square{p, White}
}

func (board *Board) String() string {
	s := ""
	if board.to_play == White {
		for i := 7; i >= 0; i-- {
			s += fmt.Sprintf("%d ", i+1)
			for j := 0; j < 8; j++ {
				s += fmt.Sprintf("%s ", board.board[j][i])
			}
			s += "\n"
		}
		s += "  "
		for i := 'a'; i <= 'h'; i++ {
			s += fmt.Sprintf("%c ", i)
		}
	} else {
		for i := 0; i < 8; i++ {
			s += fmt.Sprintf("%d ", i+1)
			for j := 7; j >= 0; j-- {
				s += fmt.Sprintf("%s ", board.board[j][i])
			}
			s += "\n"
		}
		s += "  "
		for i := 'h'; i >= 'a'; i-- {
			s += fmt.Sprintf("%c ", i)
		}
	}
	return s
}

func (c square) String() string {
	if c.piece == nil {
		return "."
	}
	s := c.piece.String()
	if c.color == Black {
		return strings.ToLower(s)
	}
	return s
}
