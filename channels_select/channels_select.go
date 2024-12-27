package main // определение пакета для текущего файла
import (
	"fmt"
	"time"
)

// ********* Channels Select *********
// select - позволяет горутине ждать данных на нескольких каналах и обрабатывать их по мере поступления на один из них
// опциональная ветка default - выполняется, если никакие прочие ветки не готовы

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	timeout := make(chan bool) // появление данных на этом канале означает таймаут для выхода из бесконечного цикла

	go func() {
		time.Sleep(10000 * time.Millisecond) // установка таймаута в 10 сек
		timeout <- true
	}()
	go func() {
		time.Sleep(4000 * time.Millisecond)
		c1 <- "ONE"
	}()
	go func() {
		time.Sleep(2000 * time.Millisecond)
		c2 <- "TWO"
	}()

	for i := 0; ; i++ {
		select {
		case <-timeout:
			fmt.Printf("TIMEOUT EXIT! (%vs)\n", i)
			return
		case msg := <-c1:
			fmt.Printf("message from C1: %v (%vs)\n", msg, i)
			time.Sleep(1000 * time.Millisecond)
		case msg := <-c2:
			fmt.Printf("message from C2: %v (%vs)\n", msg, i)
			time.Sleep(1000 * time.Millisecond)
		default:
			fmt.Printf("no messages yet... (%vs)\n", i)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
