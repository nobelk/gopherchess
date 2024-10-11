package bitboard

const (
	a8 = iota
	b8 = iota
	c8 = iota
	d8 = iota
	e8 = iota
	f8 = iota
	g8 = iota
	h8 = iota

	a7 = iota
	b7 = iota
	c7 = iota
	d7 = iota
	e7 = iota
	f7 = iota
	g7 = iota
	h7 = iota

	a6 = iota
	b6 = iota
	c6 = iota
	d6 = iota
	e6 = iota
	f6 = iota
	g6 = iota
	h6 = iota

	a5 = iota
	b5 = iota
	c5 = iota
	d5 = iota
	e5 = iota
	f5 = iota
	g5 = iota
	h5 = iota

	a4 = iota
	b4 = iota
	c4 = iota
	d4 = iota
	e4 = iota
	f4 = iota
	g4 = iota
	h4 = iota

	a3 = iota
	b3 = iota
	c3 = iota
	d3 = iota
	e3 = iota
	f3 = iota
	g3 = iota
	h3 = iota

	a2 = iota
	b2 = iota
	c2 = iota
	d2 = iota
	e2 = iota
	f2 = iota
	g2 = iota
	h2 = iota

	a1 = iota
	b1 = iota
	c1 = iota
	d1 = iota
	e1 = iota
	f1 = iota
	g1 = iota
	h1 = iota
)

const (
	white = iota
	black = iota
)

const not_ab_file = 0

var pawn_attacks [2][64]uint

type Bitboard struct {
	value uint64
}

func (board *Bitboard) GetBit(square uint) uint {
	if board.value&(1<<square) == 0 {
		return 0
	} else {
		return 1
	}
}

func (board *Bitboard) SetBit(square uint) {
	board.value |= (1 << square)
}

func (board *Bitboard) PopBit(square uint) {
	if board.value&(1<<square) != 0 {
		board.value ^= (1 << square)
	}
}

func MaskPawnAttacks(square uint, side uint) uint64 {
	var attacks uint64 = 0
	var board = new(Bitboard)

	// set piece on board
	board.SetBit(square)

	if side == white {
		attacks |= (board.value >> 7)
	} else { // black pawns
		attacks |= (board.value << 9)
	}

	return attacks
}
