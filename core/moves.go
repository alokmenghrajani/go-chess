package core

import (
	"fmt"
	"strings"
)

// We could have one struct with enough information to represent the different
// types of moves, but it felt cleaner to use multiple types.
type move interface {
	Perform_move(board *Board) *Board
	String() string
}

type xy struct {
	x, y int
}

type abstract_move struct {
	from, to xy
}

type normal_move struct {
	abstract_move
}

/* Used for keeping track of en-passant */
type pawn_push_two struct {
	abstract_move
}

type pawn_promote struct {
	abstract_move
	piece piece
}

type pawn_en_passant struct {
	abstract_move
}

type king_castle struct {
	king abstract_move
	rook abstract_move
}

type Moves []move

func perform_move_helper(board *Board, m abstract_move) *Board {
	b := *board // make a copy of the board
	if is_king(&b, m.from) {
		b.has_moved_king[b.to_play] = true
	}
	if is_rook(&b, m.from) {
		if m.from.x == 0 {
			b.has_moved_rook[b.to_play][0] = true
		}
		if m.from.x == 7 {
			b.has_moved_rook[b.to_play][1] = true
		}
	}
	b.en_passant = -2
	b.board[m.to.x][m.to.y] = b.board[m.from.x][m.from.y]
	b.board[m.from.x][m.from.y] = square{nil, White}
	b.to_play = flip_color(b.to_play)
	return &b
}

func (m normal_move) Perform_move(board *Board) *Board {
	return perform_move_helper(board, m.abstract_move)
}

func (m pawn_push_two) Perform_move(board *Board) *Board {
	b := perform_move_helper(board, m.abstract_move)
	assert(m.to.x == m.from.x, "pawn push_push_to not in straight line")
	b.en_passant = m.to.x
	return b
}

func (m pawn_promote) Perform_move(board *Board) *Board {
	b := perform_move_helper(board, m.abstract_move)
	b.board[m.to.x][m.to.y].piece = m.piece
	return b
}

func (m pawn_en_passant) Perform_move(board *Board) *Board {
	b := perform_move_helper(board, m.abstract_move)
	b.board[m.to.x][m.from.y] = square{nil, White}
	return b
}

func (m king_castle) Perform_move(board *Board) *Board {
	b := perform_move_helper(board, m.king)
	return perform_move_helper(b, m.rook)
}

func is_empty(board *Board, point xy) bool {
	return board.board[point.x][point.y].piece == nil
}

func is_opponent(board *Board, point xy) bool {
	return !is_empty(board, point) &&
		board.board[point.x][point.y].color != board.to_play
}

func is_in_bounds(x int, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func is_opponent_king(board *Board, x int, y int) bool {
	if !is_in_bounds(x, y) {
		return false
	}
	if !is_opponent(board, xy{x, y}) {
		return false
	}
	return is_king(board, xy{x, y})
}

func is_king(board *Board, point xy) bool {
	_, ok := board.board[point.x][point.y].piece.(*king)
	return ok
}
func is_rook(board *Board, point xy) bool {
	_, ok := board.board[point.x][point.y].piece.(*rook)
	return ok
}

func append_if_not_in_check(board *Board, moves Moves, move move) (Moves, bool) {
	b := move.Perform_move(board)
	if is_checking(b) {
		return moves, false
	}
	return append(moves, move), true
}

func is_checking(board *Board) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			point := xy{i, j}
			if !is_empty(board, point) && !is_opponent(board, point) {
				if board.board[i][j].piece.can_capture_king(board, point) {
					return true
				}
			}
		}
	}
	return false
}

func Is_in_check(board *Board) bool {
	b := *board
	b.to_play = flip_color(b.to_play)
	return is_checking(&b)
}

/**
 * Handles piece movement for all pieces except pawns
 */
func list_moves_common(board *Board, point xy, offsets []xy, repeat bool) Moves {
	r := make(Moves, 0, 30)
	for _, offset := range offsets {
		r = list_moves_direction(r, board, point, offset, repeat)
	}
	return r
}

func list_moves_direction(r Moves, board *Board, point xy, offset xy, repeat bool) Moves {
	for i := 1; i < 8; i++ {
		to := xy{point.x + offset.x*i, point.y + offset.y*i}
		if !is_in_bounds(to.x, to.y) {
			return r
		}
		if is_empty(board, to) || is_opponent(board, to) {
			// we are moving into an empty square or capturing a piece
			r, _ = append_if_not_in_check(board, r,
				normal_move{abstract_move{point, to}})
			// If we wanted to implement the moving piece causes check optimization,
			// we would need to check that the initial position was not in check.
		}
		if !repeat || !is_empty(board, to) {
			// we are done
			return r
		}
	}
	return r
}

func can_capture_king_common(board *Board, point xy, offsets []xy, repeat bool) bool {
	for _, offset := range offsets {
		if can_capture_king_direction(board, point, offset, repeat) {
			return true
		}
	}
	return false
}

func can_capture_king_direction(board *Board, point xy, offset xy, repeat bool) bool {
	for i := 1; i < 8; i++ {
		to := xy{point.x + offset.x*i, point.y + offset.y*i}
		if !is_in_bounds(to.x, to.y) {
			return false
		}
		if is_opponent_king(board, to.x, to.y) {
			return true
		}
		if !repeat || !is_empty(board, to) {
			// we are done
			return false
		}
	}
	return false
}

func flip_color(c color) color {
	if c == White {
		return Black
	} else {
		return White
	}
}

/**
 * Returns the direction of play (dir), the position where pawns start (end-dir)
 * and the position where pawns end.
 */
func get_dir(board *Board) (int, int, int) {
	if board.to_play == White {
		return 1, 1, 7
	} else {
		return -1, 6, 0
	}
}

func (m normal_move) String() string {
	return fmt.Sprintf("%c%d→%c%d", 'a'+m.from.x, m.from.y+1, 'a'+m.to.x,
		m.to.y+1)
}

func (m pawn_push_two) String() string {
	return fmt.Sprintf("%c%d→%c%d", 'a'+m.from.x, m.from.y+1, 'a'+m.to.x,
		m.to.y+1)
}

func (m pawn_promote) String() string {
	return fmt.Sprintf("%c%d→%c%d%s", 'a'+m.from.x, m.from.y+1, 'a'+m.to.x,
		m.to.y+1, m.piece.String())
}

func (m pawn_en_passant) String() string {
	return fmt.Sprintf("%c%d→%c%de.p.", 'a'+m.from.x, m.from.y+1, 'a'+m.to.x,
		m.to.y+1)
}

func (m king_castle) String() string {
	if m.rook.from.x == 0 {
		return "0-0-0"
	}
	assert(m.rook.from.x == 7, "rook isn't in expected position")
	return "0-0"
}

func (moves Moves) String() string {
	r := make([]string, len(moves))
	for k, m := range moves {
		r[k] = m.String()
	}
	return strings.Join(r, ",")
}
