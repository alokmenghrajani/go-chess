package mate

import (
	"fmt"
	"testing"

	"github.com/alokmenghrajani/go-chess/core"
	"github.com/alokmenghrajani/go-chess/logger"
)

func TestSimpleMateWithRook(t *testing.T) {
	s :=
		"k.......\n" +
			"...K....\n" +
			"........\n" +
			"......R.\n" +
			"........\n" +
			"........\n" +
			"........\n" +
			"........\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 3)
	expecting := "g5→a5,a8→a7,d7→c7"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Fokin & Pochtarev 1991, mate in 3
 */
func TestPuzzle1(t *testing.T) {
	s :=
		"........\n" +
			"........\n" +
			"........\n" +
			"........\n" +
			".......p\n" +
			"...kN..K\n" +
			"...PNQ..\n" +
			"........\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "f5→d5,h4→g3,f2→f5,d3→d4,e2→g3"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Fokin 1983, mate in 3
 */
func TestPuzzle2(t *testing.T) {
	s :=
		"........\n" +
			".....Q.R\n" +
			"........\n" +
			"K.......\n" +
			"..N.k...\n" +
			".......p\n" +
			"........\n" +
			"........\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "h7→h3,d3→c3,c4→d6,e4→d3,f7→f2"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Jakab 1950, mate in 3
 */
func TestPuzzle3(t *testing.T) {
	s :=
		".B.Q.R.n\n" +
			".....p..\n" +
			"B...k...\n" +
			"..r.P.pP\n" +
			"..p.....\n" +
			"......K.\n" +
			"........\n" +
			".......b\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "d1→g4,e6→f5,f8→e8,c4→c3,d8→d1"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Breuer 1951, mate in 3
 */
func TestPuzzle4(t *testing.T) {
	s :=
		"......q.\n" +
			"........\n" +
			"........\n" +
			"...pB...\n" +
			"...NpN..\n" +
			"....k...\n" +
			".pQ.....\n" +
			".Kn.....\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "h2→e2,e3→f3,d4→f5,c1→a2,c2→h2"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Moravec 1953, mate in 3
 */
func TestPuzzle5(t *testing.T) {
	s :=
		".....qn.\n" +
			"r.nNb...\n" +
			"pQ.p.p..\n" +
			"p..k....\n" +
			"....R..N\n" +
			"..p...P.\n" +
			"..PP...p\n" +
			"K..B....\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "b6→b4,d5→e4,f5→e3,a5→a4,h4→f5"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

///**
// * Sheel 1954, mate in 3 (TODO no solution?)
// */
//func TestPuzzle6(t *testing.T) {
//  s :=
//      "...N....\n" +
//      "r..pB..p\n" +
//      "........\n" +
//      "...p.B.p\n" +
//      "r....k..\n" +
//      "p..Q....\n" +
//      "Kb.P....\n" +
//      "........\n"
//  board, _ := core.Parse(s, core.White)
//  logger := NewLogger(false)
//  mate := Find(logger, board, 5)
//  expecting := "h2→e2,e3→f3,d4→f5,c1→a2,c2→h2"
//  if mate.String() != expecting {
//    t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
//  }
//}

/**
 * Palkoska 1951, mate in 3
 */
func TestPuzzle7(t *testing.T) {
	s :=
		"........\n" +
			"....p...\n" +
			"...pN.nb\n" +
			"K......p\n" +
			"nNk.....\n" +
			"..p...Q.\n" +
			"..B..pp.\n" +
			".....q.b\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "f3→e4,d5→c4,g3→f3,c4→d5,b4→d5"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Havel 1954, mate in 3
 */
func TestPuzzle8(t *testing.T) {
	s :=
		"....K...\n" +
			"..B....p\n" +
			"...N....\n" +
			"....k...\n" +
			"..r.....\n" +
			"...Q....\n" +
			"......B.\n" +
			"b.......\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "d3→e4,c4→c7,d6→e8,a1→b2,e8→f8"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Dombrovskis 1955, mate in 3
 */
func TestPuzzle9(t *testing.T) {
	s :=
		".....n.B\n" +
			"........\n" +
			"..N.....\n" +
			"..p....R\n" +
			".Rbb.k..\n" +
			"Q...pp.P\n" +
			".n..PP..\n" +
			".......K\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "e2→f3,f4→e4,h5→f5,b2→a4,c6→e7"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Bolton 1842, mate in 3
 */
func TestPuzzle10(t *testing.T) {
	s :=
		"r....q..\n" +
			"rn.p..p.\n" +
			".....k..\n" +
			"..NPN...\n" +
			"....p..P\n" +
			"....P.PQ\n" +
			"PP......\n" +
			".K......\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 5)
	expecting := "d7→e6,e7→f6,f5→d7,f6→e7,h3→f5"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Fokin 2004, mate in 4
 */
func TestPuzzle11(t *testing.T) {
	s :=
		"K.......\n" +
			"........\n" +
			"........\n" +
			"........\n" +
			"....k...\n" +
			"........\n" +
			"B..Q....\n" +
			"......N.\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 7)
	expecting := "d5→d3,e4→e3,d2→d5,e5→e4,a2→c4,e4→e5,g1→h3"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}

/**
 * Fokin 2003, mate in 4
 */
func TestPuzzle12(t *testing.T) {
	s :=
		"...N....\n" +
			"...N....\n" +
			"....B...\n" +
			"........\n" +
			".....k.B\n" +
			"...K....\n" +
			"........\n" +
			"........\n"
	board, _ := core.Parse(s, core.White)
	logger := logger.New(false)
	mate := Find(logger, board, 7)
	expecting := "d7→e5,f4→f3,d8→e6,f3→f4,h4→e1,f4→f3,e6→h3"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}
