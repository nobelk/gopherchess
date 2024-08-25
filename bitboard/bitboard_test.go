package bitboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitboard(t *testing.T) {
	board := new(Bitboard)
	board.SetBit(e4)
	board.SetBit(f2)
	board.SetBit(c3)

	assert := assert.New(t)
	var boardValue uint64 = 9011666020728832
	assert.Equal(board.value, boardValue)
}

func TestPopBit(t *testing.T) {
	board := new(Bitboard)
	board.SetBit(e4)
	board.SetBit(f2)
	board.SetBit(c3)
	board.PopBit(e4)
	board.PopBit(e4)

	assert := assert.New(t)
	var boardValue uint64 = 9011597301252096
	assert.Equal(board.value, boardValue)
}

func TestPrintBitboard(t *testing.T) {
	board := new(Bitboard)
	board.SetBit(e4)
	board.SetBit(f2)
	board.SetBit(c3)
	PrintBitboard(*board)

	//for rank := 8; rank >= 1; rank-- {
	//	fmt.Printf("a%d = iota\n b%d = iota\n c%d=iota\n d%d=iota\n e%d=iota\n f%d=iota\n g%d=iota\n h%d=iota\n \n", rank, rank, rank, rank, rank, rank, rank, rank)
	//}
}
