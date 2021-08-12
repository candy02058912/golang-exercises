package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	ID              int
	leftCS, rightCS *ChopS
}

var wg sync.WaitGroup
// the host can allow 2 philosophers to eat at once
var host = make(chan bool, 2)

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		// ask host for permission to eat
    // when host is full, this will block
		host <- true
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Philosopher %v is eating, currently there are %v philosopher(s) eating.\n", p.ID, len(host))

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		<-host

	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		// (i+1)%5 for the case of the 4th philosopher
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}
