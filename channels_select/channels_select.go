package main // определение пакета для текущего файла
import (
	"fmt"
	"time"
)

// ********* Channels Select *********
// select - позволяет ждать данных на нескольких каналах и обрабатывать их по мере поступления на один из них

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
