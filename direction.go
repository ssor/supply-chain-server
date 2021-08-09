package main

type Direction int

const (
	NoDirection Direction = -1
	North       Direction = 0
	East        Direction = 1
	South       Direction = 2
	West        Direction = 3
)

func (d Direction) String() string {
	switch d {
	case NoDirection:
		return "unknown"
	case North:
		return "North"
	case South:
		return "South"
	case West:
		return "West"
	case East:
		return "East"
	}
	return "unknown"
}
