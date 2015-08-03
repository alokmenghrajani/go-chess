package selfmate

import (
	"fmt"
	"testing"

	"github.com/alokmenghrajani/go-chess/core"
	"github.com/alokmenghrajani/go-chess/logger"
)

/**
 * Grin 1999, selfmate in 2
 */
func TestPuzzle1(t *testing.T) {
	s :=
		"........\n" +
			"........\n" +
			"........\n" +
			"........\n" +
			".p......\n" +
			"k.......\n" +
			"..Q.....\n" +
			"KB......\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 4)
	expecting := "b3→b2,d1→g4,b4→b3,c2→d1"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}

/**
 * Mikan 1963, selfmate in 3
 */
func TestPuzzle2(t *testing.T) {
	s :=
		"........\n" +
			"........\n" +
			"........\n" +
			"...R....\n" +
			".....N.Q\n" +
			".p.Bpk..\n" +
			".p..R..p\n" +
			".N.KB.nr\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 6)
	expecting := "e2→c3,b1→c3,f3→e4,g6→h5,g1→e2,d3→g6"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}

/**
 * Avner & Gordian 1995, selfmate in 3
 */
func TestPuzzle3(t *testing.T) {
	s :=
		"........\n" +
			"Bpp.....\n" +
			".rr..p..\n" +
			"KBk..N..\n" +
			"R.......\n" +
			"p.......\n" +
			"Q....N..\n" +
			"........\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 6)
	expecting := "b6→b5,c4→b5,c5→c6,f2→e4,c6→d6,b5→c4"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}
