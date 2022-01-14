package main

import (
	"log"
	"math/rand"
	"time"
)

type Strategy int

const (
	Stay   Strategy = 0
	Switch Strategy = 1
	Games           = 1000
	Doors           = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Printf("Lazy monty (select random door after selections are made) test with %d runs for both strategies", Games)
	stayWin := 0
	for i := 0; i < Games; i++ {
		if Monty(Stay) {
			stayWin++
		}
	}
	switchWin := 0
	for i := 0; i < Games; i++ {
		if Monty(Switch) {
			switchWin++
		}
	}
	log.Printf("Stay over %d wins: %d times %0.2f%%", Games, stayWin, float64(stayWin)/float64(Games)*100.0)
	log.Printf("Switch over %d wins: %d times %0.2f%%", Games, switchWin, float64(switchWin)/float64(Games)*100.0)
}

func Monty(s Strategy) bool {
	selectedDoor := rand.Intn(Doors)
	montyDoor := rand.Intn(Doors)
	for montyDoor == selectedDoor {
		montyDoor = rand.Intn(Doors)
	}
	switch s {
	case Stay:
	case Switch:
		previousSelection := selectedDoor
		for selectedDoor == montyDoor || selectedDoor == previousSelection {
			selectedDoor = rand.Intn(Doors)
		}
	}
	winningDoor := rand.Intn(Doors)
	return winningDoor == selectedDoor
}
