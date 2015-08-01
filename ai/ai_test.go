package ai

import (
	"fmt"
	"testing"

	"github.com/alokmenghrajani/go-chess/core"
)

//func TestFindMate(t *testing.T) {
//	s :=
//		"k.......\n" +
//			"...K....\n" +
//			"........\n" +
//			"......R.\n" +
//			"........\n" +
//			"........\n" +
//			"........\n" +
//			"........\n"
//	board, _ := core.Parse(s, core.White)
//	mate := FindMate(board, 3)
//	expecting := "g5→a5,a8→a7,d7→c7"
//	if mate.String() != expecting {
//		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
//	}
//}
//
///**
// * Fokin & Pochtarev 1991, mate in 3
// */
//func TestPuzzle1(t *testing.T) {
//	s :=
//		"........\n" +
//			"........\n" +
//			"........\n" +
//			"........\n" +
//			".......p\n" +
//			"...kN..K\n" +
//			"...PNQ..\n" +
//			"........\n"
//	board, _ := core.Parse(s, core.White)
//	mate := FindMate(board, 5)
//	expecting := "f5→d5,h4→g3,f2→f5,d3→d4,e2→g3"
//	if mate.String() != expecting {
//		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
//	}
//}
//
///**
// * Fokin 1983, mate in 3
// */
//func TestPuzzle2(t *testing.T) {
//	s :=
//		"........\n" +
//			".....Q.R\n" +
//			"........\n" +
//			"K.......\n" +
//			"..N.k...\n" +
//			".......p\n" +
//			"........\n" +
//			"........\n"
//	board, _ := core.Parse(s, core.White)
//	mate := FindMate(board, 5)
//	expecting := "h7→h3,d3→c3,c4→d6,e4→d3,f7→f2"
//	if mate.String() != expecting {
//		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
//	}
//}

///**
// * Jakab 1950, mate in 3
// */
//func TestPuzzle3(t *testing.T) {
//  s :=
//      ".B.Q.R.n\n" +
//      ".....p..\n" +
//      "B...k...\n" +
//      "..r.P.pP\n" +
//      "..p.....\n" +
//      "......K.\n" +
//      "........\n" +
//      ".......b\n"
//  board, _ := core.Parse(s, core.White)
//  logger := NewLogger(false)
//  mate := FindMate(logger, board, 5)
//  expecting := "d1→g4,e6→f5,f8→e8,c4→c3,d8→d1"
//  if mate.String() != expecting {
//    t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
//  }
//}

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
	logger := NewLogger(false)
	mate := FindMate(logger, board, 5)
	expecting := "h2→e2,e3→f3,d4→f5,c1→a2,c2→h2"
	if mate.String() != expecting {
		t.Error(fmt.Printf("Expecting: %s, but got %s", expecting, mate))
	}
}
