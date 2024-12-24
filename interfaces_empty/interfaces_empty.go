package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* Empty Interfaces *********
// Тип интерфейса, в котором указаны ноль методов - называется пустым интерфейсом - interface{}
// Пустые интерфейсы могут содержать значение любого типа. Используются, когда необходимо работать со значениями заранее неизвестного типа.
// Например fmt.Print принимает аргументы с типом interface{}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
