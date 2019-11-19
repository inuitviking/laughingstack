package main

import (
	"github.com/inuitviking/laughingstack/Go/1.12.x/concurrent"
	"github.com/inuitviking/laughingstack/Go/1.12.x/laugher"
	"github.com/inuitviking/laughingstack/Go/1.12.x/simple"
)

func main() {
	// Simple laugh
	allLaughs(simple.SimpleLaugher{})

	// Concurrent laugh
	allLaughs(concurrent.ConcurrentLaugher{})
}

func allLaughs(l laugher.Laugher) {
	l.Chuckle()
	l.Chortle()
	l.Guffaw()
}
