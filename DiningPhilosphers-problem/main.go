package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Aristotle", leftFork: 0, rightFork: 1},
	{name: "Socrates", leftFork: 1, rightFork: 2},
	{name: "Fermi", leftFork: 2, rightFork: 3},
	{name: "Dalton", leftFork: 3, rightFork: 4},
}

var hunger = 3
var eatTime = 1 * time.Second
var sleepTime = 1 * time.Second
var thinkTime = 2 * time.Second

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("Table is empty")

	dine()
	fmt.Println("Table is empty")

}

func dine() {

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	wg.Wait()

}
func diningProblem(philsopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.", philsopher.name)
	seated.Done()

	for i := hunger; i > 0; i-- {

		//to solve the hidden race condition
		if philsopher.leftFork > philsopher.rightFork {
			forks[philsopher.rightFork].Lock()
			fmt.Printf("\t %s takes the right fork\n", philsopher.name)
			forks[philsopher.leftFork].Lock()
			fmt.Printf("\t %s takes the left fork\n", philsopher.name)
		} else {
			forks[philsopher.leftFork].Lock()
			fmt.Printf("\t %s takes the left fork\n", philsopher.name)
			forks[philsopher.rightFork].Lock()
			fmt.Printf("\t %s takes the right fork\n", philsopher.name)
		}

		fmt.Printf("\t %s is having both the forks and is eating\n", philsopher.name)
		time.Sleep(eatTime)
		fmt.Printf("\t %s is thinking\n", philsopher.name)
		time.Sleep(thinkTime)

		forks[philsopher.leftFork].Unlock()
		forks[philsopher.rightFork].Unlock()

		fmt.Printf("\t %s put down the forks.\n", philsopher.name)

	}
	fmt.Println(philsopher.name, "is satisfied")
	fmt.Println(philsopher.name, "left the table")
}
