package main

/*
 * Implement the dining philosopher’s problem with the following
 * constraints/modifications.
 *
 * 1. There should be 5 philosophers sharing chopsticks, with one
 *    chopstick between each adjacent pair of philosophers.
 * 2. Each philosopher should eat only 3 times (not in an infinite
 *    loop as we did in lecture)
 * 3. The philosophers pick up the chopsticks in any order, not
 *    lowest-numbered first (which we did in lecture).
 * 4. In order to eat, a philosopher must get permission from a host
 *    which executes in its own goroutine.
 * 5. The host allows no more than 2 philosophers to eat concurrently.
 * 6. Each philosopher is numbered, 1 through 5.
 * 7. When a philosopher starts eating (after it has obtained
 *    necessary locks) it prints “starting to eat <number>” on a line
 *    by itself, where <number> is the number of the philosopher.
 * 8. When a philosopher finishes eating (before it has released its
 *    locks) it prints “finishing eating <number>” on a line by itself,
 *    where <number> is the number of the philosopher.
 */

import (
	"fmt"
	"sync"
	"time"
)

var eatWG sync.WaitGroup

type chopStick struct{ sync.Mutex }
type phil struct {
	id              int
	leftCS, rightCS *chopStick
}

// The method eat for the philosophers
func (p phil) eat() {
	defer eatWG.Done()
	for i := 0; i < 3; i++ {

		// The philosopher is eating
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("The philosopher %v is eating round %v\n", p.id+1, i+1)
		time.Sleep(time.Millisecond * 500)

		// The philosopher has finish
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fmt.Printf("The philosopher %v has finished round %v\n", p.id+1, i+1)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// 5 philosophers
	number := 5

	// Creating the chopsticks
	chopSticks := make([]*chopStick, number)
	for i := 0; i < number; i++ {
		chopSticks[i] = new(chopStick)
	}

	// Create thr philosopher and assign the two chopsticks
	//and sit the at the table
	philosophers := make([]*phil, number)
	for i := 0; i < number; i++ {
		philosophers[i] = &phil{
			id:      i,
			leftCS:  chopSticks[i],
			rightCS: chopSticks[(i+1)%number],
		}
		eatWG.Add(1)
		go philosophers[i].eat()
	}
	eatWG.Wait()
}
