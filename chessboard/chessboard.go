package chessboard

import "fmt"

// Declare un tipo llamado Rank que almacene si un cuadrado está ocupado por una pieza; esto será un slice de bools
type Rank []bool

// Declare un tipo llamado Chessboard que contiene un mapa de ocho rangos, al que se accede con las teclas de "A" a "H"
type Chessboard map[string][]bool

func newChessboardMain() Chessboard {
	//
	//   1 2 3 4 5 6 7 8
	// A # _ # _ _ _ _ # A
	// B _ _ _ _ # _ _ _ B
	// C _ _ # _ _ _ _ _ C
	// D _ _ _ _ _ _ _ _ D
	// E _ _ _ _ _ # _ # E
	// F _ _ _ _ _ _ _ _ F
	// G _ _ _ # _ _ _ _ G
	// H # # # # # # _ # H
	//   1 2 3 4 5 6 7 8
	return Chessboard{
		"A": Rank{true, false, true, false, false, false, false, true},
		"B": Rank{false, false, false, false, true, false, false, false},
		"C": Rank{false, false, true, false, false, false, false, false},
		"D": Rank{false, false, false, false, false, false, false, false},
		"E": Rank{false, false, false, false, false, true, false, true},
		"F": Rank{false, false, false, false, false, false, false, false},
		"G": Rank{false, false, false, true, false, false, false, false},
		"H": Rank{true, true, true, true, true, true, false, true},
	}
}

// CountInRank devuelve cuántas casillas están ocupadas en el tablero de ajedrez,
// dentro del rank dado
func CountInRank(cb Chessboard, rank string) int {
	counter := 0
	for letter, value := range cb {
		if letter == rank {
			for _, hasBlcok := range value {
				if hasBlcok {
					counter++
				}
			}
		}
	}
	return counter
}

// CountInFile devuelve cuantas casillas están ocupadas en el tablero de ajedrez,
// dentro del archivo dado
func CountInFile(cb Chessboard, file int) int {
	if file < 0 || file > 8 {
		return 0
	}

	counter := 0
	for _, slice := range cb {
		if slice[(file - 1)] {
			counter++
		}

	}
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	couter := 0
	for _, v := range cb {
		for _, a := range v {
			fmt.Println(a)
			couter++
		}
	}
	return couter
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	counter := 0
	for _, letter := range letters {
		counter += CountInRank(cb, letter)
	}
	return counter
}

func main() {
	board := newChessboardMain()

	//fmt.Println(CountInRank(board, "A"))
	//fmt.Println(CountInFile(board, 8))
	//fmt.Println(CountAll(board))
	fmt.Println(CountOccupied(board))
}
