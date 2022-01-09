package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	left, right *chopstick
}

func main() {

	philosophers := make([]*philosopher, 5)
	chopsticks := make([]*chopstick, 5)

	for i := 0; i < 5; i++ {
		chopsticks[i] = new(chopstick)
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{chopsticks[i], chopsticks[(i+1)%5]}
	}

	var wg sync.WaitGroup

	// Single channel for requests from philosophers to host
	c := make(chan string)

	// Array of channels for responses from host to philosophers
	var pc [5]chan bool
	for i := range pc {
		pc[i] = make(chan bool)
	}

	// Host routine
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Dinner is ready!")
		philosophers_dining := 0
		// Reading messages from philosophers
		for msg := range c {
			msg_split := strings.Fields(msg)
			id, _ := strconv.Atoi(msg_split[0])
			// Wants to eat
			if msg_split[1] == "s" {
				if philosophers_dining < 2 {
					fmt.Printf("Philosopher %d can eat, %d philosophers dining\n", id+1, philosophers_dining)
					pc[id] <- true
					philosophers_dining++
				} else {
					fmt.Printf("Philosopher %d needs to wait, %d philosophers dining\n", id+1, philosophers_dining)
					pc[id] <- false
				}
				// Finished eating
			} else {
				philosophers_dining--
			}
		}
	}(&wg)

	// Philosopher routine
	for id := 0; id < 5; id++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup, p *philosopher) {
			defer wg.Done()
			times_eaten := 0
			for times_eaten < 3 {
				c <- strconv.Itoa(id) + " s"
				can_eat := <-pc[id]
				if can_eat {
					p.left.Lock()
					p.right.Lock()
					fmt.Printf("Philosopher %d starting to eat\n", id+1)
					time.Sleep(1000 * time.Millisecond)
					times_eaten++
					fmt.Printf("Philosopher %d finishing eating, eaten %d times\n", id+1, times_eaten)
					p.left.Unlock()
					p.right.Unlock()
					c <- strconv.Itoa(id) + " f"
				} else {
					time.Sleep(1000 * time.Millisecond)
				}
			}
		}(id, &wg, philosophers[id])
	}

	wg.Wait()
	close(c)
}