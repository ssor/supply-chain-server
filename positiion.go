package main

import "fmt"

// func newPositionWithSN(sn, id, x, y int) Position {
// 	p := Position{
// 		Id: id,
// 	}
// 	p.X = x
// 	p.Y = y
// 	p.SerailNumber = sn
// 	return p
// }

func newPosition(id, x, y int) Position {
	p := Position{
		Id: id,
	}
	p.X = x
	p.Y = y
	return p
}
func newEmptyPosition(x, y int) Position {
	p := Position{
		Id: 0,
	}
	p.X = x
	p.Y = y
	return p
}

type Position struct {
	Id int
	Coord
}

func (p Position) WithSN(sn int) Position {
	p.SerailNumber = sn
	return p
}

func (p Position) String() string {
	return fmt.Sprintf("%d:%s", p.Id, p.Coord)
}

func (p Position) IsEmpty() bool {
	return p.Id <= 0
}

func (p Position) Equal(pos Position) bool {
	return p.Coord.Equal(pos.Coord)
}

type Positions []Position
type MaybePosition struct {
	Maybe bool
	Position
}

func (positions Positions) find(id int) MaybePosition {
	for _, c := range positions {
		if c.Id == id {
			return MaybePosition{Maybe: true, Position: c}
		}
	}
	return MaybePosition{}
}

func (positions Positions) add(id, x, y int) Positions {
	return append(positions, newPosition(id, x, y))
}

func newPositionLine() *PositionLine {
	return &PositionLine{
		SN: 0,
		// NextIndex: -1,
	}
}

type PositionLine struct {
	NextIndex int
	SN        int
	Positions
}

func (pl *PositionLine) Shift() MaybePosition {
	if pl.SN <= 0 {
		return MaybePosition{}
	}
	if pl.NextIndex >= pl.SN {
		return MaybePosition{}
	}
	p := pl.Positions[pl.NextIndex]
	pl.NextIndex++
	return MaybePosition{Maybe: true, Position: p}
}

func (pl *PositionLine) AddPosition(p Position) {
	pl.Positions = append(pl.Positions, p.WithSN(pl.SN))
	pl.SN++
}
