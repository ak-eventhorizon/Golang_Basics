package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* For *********
// в go одна конструкция для циклов - for
// for может состоять из трех компонентов, разделенных точкой с запятой:

// for <init statement> ; <condition expression> ; <post statement> { ... }

// <init statement> 		- выполняется однократно перед первой итерацией
// <condition expression> 	- вычисляется перед каждой итерацией, если false - то итерации не будет
// <post statement> 		- выполняется в конце каждой итерации

func main() {

	sum := 0
	// базовое использование
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	counter := 1
	// чтобы использовать for как цикл while - <init statement> и <post statement> - могут быть пропущены
	for counter < 100 {
		counter += counter
	}
	fmt.Println(counter)

	c := 0
	// бесконечный цикл, необходимо предусмотреть условие выхода из него
	for {
		c = c + 2
		if c > 10 {
			fmt.Println("Infinite loop exit")
			break
		}
	}
}
