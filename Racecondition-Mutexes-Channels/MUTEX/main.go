//simple program to demonstrate the race conditions in go and it s solution

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var msg string
// var wg sync.WaitGroup

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()
// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

// func main() {
// 	msg = "Hello world"

// 	var mutex sync.Mutex

// 	wg.Add(2)
// 	go updateMessage("Hello cosmos", &mutex)
// 	go updateMessage("Hello universe", &mutex)
// 	wg.Wait()

// 	fmt.Println(msg)

// }

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Initial account balance $ %d.00", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 200},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("on week %d you have earned %d.00 from %s\n", week, income.Amount, income.Source)
			}

		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Final bank balance is $ %d.00", bankBalance)

}
