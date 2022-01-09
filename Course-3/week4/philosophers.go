package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const COUNT = 5
const MEALS = 3
const MAX_EATERS = 2

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	id             int
	leftChopStick  *ChopStick
	rightChopStick *ChopStick
}

var wg sync.WaitGroup
var ch = make(chan bool, MAX_EATERS)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	chopSticks := make([]*ChopStick, COUNT)
	philosophers := make([]*Philosopher, COUNT)

	for i := 0; i < COUNT; i++ {
		chopSticks[i] = new(ChopStick)
	}

	for i := 0; i < COUNT; i++ {
		leftChopStick := chopSticks[i]
		rightChopStick := chopSticks[(i+1)%COUNT]
		philosophers[i] = &Philosopher{i + 1, leftChopStick, rightChopStick}

		for j := 0; j < MEALS; j++ {
			wg.Add(1)
			go philosophers[i].eat(j + 1)
		}
	}

	wg.Wait()
}

func (philosopher Philosopher) eat(meal int) {
	getPermission()

	philosopher.leftChopStick.Lock()
	philosopher.rightChopStick.Lock()

	fmt.Printf("Philosopher %d started to eat meal %d...\n", philosopher.id, meal)
	randomWait(2)
	fmt.Printf("Philosopher %d finished eating meal %d!\n", philosopher.id, meal)

	philosopher.rightChopStick.Unlock()
	philosopher.leftChopStick.Unlock()

	freePermission()
	wg.Done()
}

func getPermission() {
	ch <- true
}

func freePermission() {
	<-ch
}

func randomWait(seconds int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(seconds*1000)))
}
