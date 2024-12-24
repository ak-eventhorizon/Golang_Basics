package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* Stringer *********
// Описывает тип, который умеет представлять себя как строку с помощью метода String()
// Пакет fmt (и многие другие) ищут этот интерфейс, чтобы выводить в консоль значения

// type Stringer interface {
// 	String() string
// }

type IPAddr [4]byte

func (a IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
