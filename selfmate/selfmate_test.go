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

/**
 * Mikan 1955, selfmate in 3
 */
func TestPuzzle4(t *testing.T) {
	s :=
		"........\n" +
			"...N...p\n" +
			"....R.BQ\n" +
			".R..bq..\n" +
			"....k..K\n" +
			"....P...\n" +
			"....PP..\n" +
			".......N\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 6)
	expecting := "g6→h7,h6→h7,e4→f5,f2→f3,f5→g6,e6→e8"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}

/**
 * Mikan 1977, selfmate in 3
 */
func TestPuzzle5(t *testing.T) {
	s :=
		"......R.\n" +
			"......NR\n" +
			"........\n" +
			"........\n" +
			".....Pk.\n" +
			"...N.qp.\n" +
			"......Qp\n" +
			"...BB..K\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 6)
	expecting := "e2→h2,g2→h2,g4→h4,h5→g5,f3→e2,h7→h5"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}

/**
 * Bakcsi 1981, selfmate in 3
 */
func TestPuzzle6(t *testing.T) {
	s :=
		"...R....\n" +
			"........\n" +
			".N..p...\n" +
			"...b.B..\n" +
			"..R.r...\n" +
			"...k....\n" +
			".P.p.Q..\n" +
			"...K....\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 6)
	expecting := "d5→b3,b4→b3,e5→d4,f2→d4,e6→e5,c4→b4"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}

/**
 * Labai 2007, selfmate in 3
 */
func TestPuzzle7(t *testing.T) {
	s :=
		"..B.....\n" +
			"...RN...\n" +
			".....p..\n" +
			"........\n" +
			"....knRK\n" +
			"....n.B.\n" +
			"....Q...\n" +
			"........\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 6)
	expecting := "f6→g5,g4→g5,e5→f5,e2→e3,e4→e5,e7→f5"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}

/**
 * Kubbel 1933, selfmate in 3
 */
func TestPuzzle8(t *testing.T) {
	s :=
		"....k.q.\n" +
			"p....np.\n" +
			"....K.B.\n" +
			"..B.QN.p\n" +
			"..NR...p\n" +
			".p......\n" +
			"bP......\n" +
			"....b...\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 6)
	expecting := "f7→d6,d5→c4,e8→d8,e6→d5,a2→b1,c4→b6"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}

/**
 * Brown 1974, selfmate in 3
 */
func TestPuzzle9(t *testing.T) {
	s :=
		"........\n" +
			"........\n" +
			"........\n" +
			"........\n" +
			".....N.p\n" +
			".......B\n" +
			"........\n" +
			"....RKBk\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	selfmate := Find(logger, board, 6)
	expecting := "h3→g2,f4→e2,h4→h3,h3→g2,h1→h2,g1→f2"
	if selfmate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, selfmate))
	}
}
