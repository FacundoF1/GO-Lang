package examples

import (
	"fmt"
	"time"
)

func TryRoutines() {

	go DelayedAnnouncement("This will print later", 10)

	finished := make(chan int)
	go DelayedAnnouncementWithChannel("This will also print later", 10, finished)
	<-finished

	bufferedChannel := make(chan int, 10)
	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3
	readData := <-bufferedChannel
	fmt.Println(readData)

	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)
	quit := make(chan int)
	go ReadFromChannels(chan1, chan2, chan3, quit)
	chan2 <- 1
	chan3 <- 1
	chan3 <- 1
	chan1 <- 1
	quit <- 1

	data := make(chan string)
	go RangeOverChannel(data)
	data <- "10"
	data <- "20"
	data <- "30"
	close(data)
}

func DelayedAnnouncement(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(message)
}

func DelayedAnnouncementWithChannel(message string, delay time.Duration, finished chan int) {
	time.Sleep(delay)
	fmt.Println(message)
	finished <- 1
}

func ReadFromChannels(chan1, chan2, chan3, quit chan int) {
	for {
		select {
		case <-chan1:
			fmt.Println("Read from channel 1")
		case <-chan2:
			fmt.Println("Read from channel 2")
		case <-chan3:
			fmt.Println("Read from channel 3")
		case <-quit:
			return
		}
	}
}

func RangeOverChannel(data <-chan string) {
	for elem := range data {
		fmt.Println(elem)
	}
}
