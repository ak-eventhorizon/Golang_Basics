package main // определение пакета для текущего файла
import (
	"fmt" // подключение пакета fmt
)

// Рекурсия - на примере вычисления факториала
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(2))
	fmt.Println(factorial(6))
}
