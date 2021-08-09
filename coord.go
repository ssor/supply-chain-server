package main

import "fmt"

func newCoord(x, y int) Coord {
	return Coord{X: x, Y: y}
}

// func newCoord2(sn, x, y int) Coord {
// 	return Coord{X: x, Y: y, SerailNumber: sn}
// }

type Coord struct {
	SerailNumber int `json:"sn"`
	X            int `json:"x"`
	Y            int `json:"y"`
}

func (coord Coord) String() string {
	return fmt.Sprintf("(%d, %d)", coord.X, coord.Y)
}

func (coord Coord) Equal(c Coord) bool {
	return coord.X == c.X && coord.Y == c.Y
}

type Coords []Coord
