package core

type piece interface {
	String() string
}

type pawn struct {
}
type rook struct {
}
type knight struct {
}
type bishop struct {
}
type queen struct {
}
type king struct {
}

func (pawn) String() string {
	return "p"
}
func (rook) String() string {
	return "r"
}
func (knight) String() string {
	return "n"
}
func (bishop) String() string {
	return "b"
}
func (queen) String() string {
	return "q"
}
func (king) String() string {
	return "k"
}
