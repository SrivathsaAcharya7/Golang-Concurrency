package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HaircutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	isSleeping := false
	color.Yellow("%s goes to waiting room to check for clients", barber)
	for {
		if len(shop.ClientChan) == 0 {
			color.Yellow("There is no client so barber %s goes to sleep", barber)
			isSleeping = true
		}
		client, shopOpen := <-shop.ClientChan

		if shopOpen {
			if isSleeping {
				color.Yellow("%s wakes up %s", client, barber)
				isSleeping = false
			}
			shop.cutHair(barber, client)
		} else {
			shop.sendBarberHome(barber)
			return
		}
	}

}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HaircutDuration)
	color.Green("%s finished cutting %s's hair.", barber, client)
}
func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.BarbersDoneChan <- true
}
func (shop *BarberShop) closeShopFOrDay() {
	color.Cyan("Closing shop for the day")
	close(shop.ClientChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)
	color.Green("---------------------")
	color.Green("The shop has now closed for the day and everyone has gone home")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("*** %s arrives!***", client)

	if shop.Open {
		select {
		case shop.ClientChan <- client:
			color.Blue("%s takes a seat in the waiting room", client)
		default:
			color.Red("The waiting room is full so %s leaves", client)

		}
	} else {
		color.Red("The shop is already closed, so %s leaves.", client)
	}

}
