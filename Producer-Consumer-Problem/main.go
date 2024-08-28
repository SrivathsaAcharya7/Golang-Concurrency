package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order #%d\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}
		total++

		fmt.Printf("making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** we ran out of ingrediants for pizza #%d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook has quit making pizza #%d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order is ready #%d", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {

	var i = 0

	for {
		currentOrder := makePizza(i)

		if currentOrder != nil {
			i = currentOrder.pizzaNumber
			select {
			case pizzaMaker.data <- *currentOrder:

			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return

			}
		}
	}

}

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	color.Cyan("The Pizzeria is open for business")
	color.Cyan("*********************************")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Customer is really mad")
			}
		} else {
			color.Cyan("Done making pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing channel", err)
			}
		}
	}

	color.Cyan("*********************************")
	color.Cyan("Done for the day")

	color.Cyan("We made %d pizzas but failed to make %d, with %d attempts in total", pizzaMade, pizzaFailed, total)

	switch {
	case pizzaFailed > 9:
		color.Red("It was an awful day")
	case pizzaFailed >= 6:
		color.Red("It was a pretty bad day")
	case pizzaFailed >= 4:
		color.Yellow("It was an ok day")
	case pizzaFailed >= 2:
		color.Yellow("It was a pretty good day")
	default:
		color.Green("It was a great day")
	}

}
