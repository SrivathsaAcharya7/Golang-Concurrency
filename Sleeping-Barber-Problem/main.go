//channel basic demonstration program

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func shout(ping chan string, pong chan string) {

// 	for {
// 		s := <-ping

// 		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
// 	}

// }

// func main() {

// 	ping := make(chan string)
// 	pong := make(chan string)

// 	go shout(ping, pong)

// 	fmt.Println("Type anything and press enter or press Q to quit")
// 	for {
// 		fmt.Println("->")

// 		var userInput string

// 		_, _ = fmt.Scanln(&userInput)

// 		if userInput == strings.ToLower("q") {
// 			break
// 		}
// 		ping <- userInput
// 		result := <-pong
// 		fmt.Println("Response: ", result)
// 	}

// 	fmt.Println("All Done. Closing Channels")
// 	close(ping)
// 	close(pong)

// }

//channels with select statements in action

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func server1(ch chan string) {
// 	for {
// 		time.Sleep(6 * time.Second)
// 		ch <- "This is from server 1"
// 	}

// }
// func server2(ch chan string) {
// 	for {
// 		time.Sleep(3 * time.Second)
// 		ch <- "This is from server 2"
// 	}
// }

// func main() {

// 	fmt.Println("Channels with select statement demo")
// 	fmt.Println("-----------------------------------")

// 	channel1 := make(chan string)
// 	channel2 := make(chan string)

// 	go server1(channel1)
// 	go server2(channel2)
// 	for {
// 		select {
// 		case s1 := <-channel1:
// 			fmt.Println("Case one: ", s1)
// 		case s2 := <-channel1:
// 			fmt.Println("Case two: ", s2)
// 		case s3 := <-channel2:
// 			fmt.Println("Case three: ", s3)
// 		case s4 := <-channel2:
// 			fmt.Println("Case four: ", s4)
// 		}
// 	}

// }

//Buffered channels

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func listenToChan(ch chan int) {
// 	for {
// 		i := <-ch
// 		fmt.Println("Got", i, "from channel")
// 		time.Sleep(1 * time.Second)
// 	}

// }

// func main() {

// 	ch := make(chan int, 50)

// 	go listenToChan(ch)

// 	for i := 0; i <= 100; i++ {
// 		fmt.Println("Sending", i, "to channel")
// 		ch <- i
// 		fmt.Println("Sent", i, "to channel")

// 	}
// 	fmt.Println("Done")
// 	close(ch)

// }

//Sleeping barber problem

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var cutDuration = 1000 * time.Millisecond
var arrivalRate = 100
var timeOpen = 10 * time.Second

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	color.Yellow("Sleeping Barbers Problem")
	color.Yellow("------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HaircutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientChan:      clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The barber shop is open now!")

	shop.addBarber("Frank")
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopFOrDay()
		closed <- true

	}()

	i := 1

	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	<-closed

}
