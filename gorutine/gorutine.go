package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* Gorutine *********
// Горутины представляют параллельные операции, которые могут выполняться независимо от функции, в которой они запущены

func factorial(n int) {
	if n < 1 {
		fmt.Println("Unvalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
}

func main() {

	for i := 1; i < 7; i++ {
		go factorial(i) // запуск функции в параллельном процессе
	}
	fmt.Scanln() // ждем ввода пользователя, чтобы процесс main не завершился и не завершил вместе с собой все порожденные горутины
	fmt.Println("The End")
}
