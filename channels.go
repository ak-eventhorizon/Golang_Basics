package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* Channels *********
// Каналы представляют инструменты коммуникации между горутинами.
// Как правило, отправителем данных является одна горутина, а получателем - другая.

func main() {

	var intCh chan int = make(chan int) // функция make используется для инициализации пустого канала

	go func() {
		fmt.Println("Gorutine starts")
		intCh <- 5
		intCh <- 6
		intCh <- 7
	}()

	fmt.Println(<-intCh)
	<-intCh
	fmt.Println(<-intCh)
}
