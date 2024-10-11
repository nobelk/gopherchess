package bitboard

import (
	"fmt"
)

func PrintBitboard(input Bitboard) {
	var mask uint64 = 1

	fmt.Println("")
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			if file == 0 {
				fmt.Printf("%d  ", 8-rank)
			}
			square := rank*8 + file
			if (input.value & (mask << square)) == 0 {
				fmt.Printf(" %d", 0)
			} else {
				fmt.Printf(" %d", 1)
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n    a b c d e f g h")
	fmt.Printf("\n Bitboard: %d \n\n", input.value)
}

func PrintBitboardHelper(board Bitboard) {
	fmt.Println("")
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			if file == 0 {
				fmt.Printf("%d  ", 8-rank)
			}
			square := uint(rank*8 + file)
			if file > 1 {
				board.SetBit(square)
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n    a b c d e f g h")
	fmt.Printf("\n Bitboard: %d \n\n", board.value)
}
