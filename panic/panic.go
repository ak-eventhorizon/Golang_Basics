package main // определение пакета для текущего файла
import (
	"fmt" // подключение пакета fmt
)

// Оператор panic позволяет сгенерировать ошибку и выйти из программы

func main() {
	fmt.Println(divide(15, 5))
	fmt.Println(divide(4, 0))
	fmt.Println("Program has been finished")
}
func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}
