package main

import (
	"bytes"
	"time"
	"fmt"
	"math/rand"
)

// data structure Game containing all of the integers
type Game struct {
	a [][]bool
	w, h  int
}

// data structure Life containing two states of Game for comparison
type Life struct {
	a, b *Game
	w, h int
}

func NewLife(w, h int) *Life{
	a := NewGame(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a, b: NewGame(w, h),
		w: w, h: h,
	}
}

func NewGame(w, h int) *Game {
	//push row
	s := make([][]bool, h)
	for i := 0; i < w; i++ {
		s[i] = make([]bool, w)
	}
	return &Game{a: s, w: w, h: h}
}

func (g *Game) Set(x, y int, b bool){
	// loop through fields
	g.a[x][y] = b
}

func (g *Game) Alive(x, y int) bool {
	x += g.w
	x %= g.w
	y += g.h
	y %= g.h
	return g.a[y][x]
}

func (g *Game) isAlive(x, y int) bool {
	countAlive := 0
	for d := -1; d <= 1; x++ {
		for k := -1; k <= 1; k++ {
			if(d != 0 || k != 0) && (g.Alive(x+d, y+d)){
				countAlive++
			}
		}
	}
	return countAlive == 3 || countAlive == 2 && g.isAlive(x, y)
}

func (l *Life) Step() {
	//
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.isAlive(x, y))
		}
	}
	l.a, l.b = l.b, l.a
}

func (l *Life) PrintGame() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.isAlive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	// create new board
	a := NewLife(30, 30)

	// loop through n times
	for i := 0; i < 10; i++ {
		a.Step()

		//Print board
		z := a.PrintGame()
		fmt.Print("\x0c", z)

		//Sleep between steps
		time.Sleep(time.Second / 30)
	}
	// print board
}
