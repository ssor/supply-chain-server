package main

import (
	"strconv"
	"time"
)

type Game struct {
	Id string
	StartTime time.Time
}

func NewGame() *Game{
	g:=&Game{}
	g.Id = strconv.FormatInt(time.Now().UnixNano(),10)
	g.StartTime = time.Now()
	return g
}