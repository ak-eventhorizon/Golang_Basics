package main // определение пакета для текущего файла
import (
	"fmt" // подключение пакета fmt
)

// Замыкание - возвращает анонимную функцию со своей копией всего контекста родительской функции
func counter(initVal int) func() int {
	var x = initVal
	return func() int {
		x++
		return x
	}
}

func main() {
	c1 := counter(10)
	c2 := counter(100)

	for i := 0; i < 3; i++ {
		fmt.Println("counter1 -", c1())
	}

	for i := 0; i < 3; i++ {
		fmt.Println("counter2 =", c2())
	}

	for i := 0; i < 3; i++ {
		fmt.Println("counter1 -", c1())
	}

}
