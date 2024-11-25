package main // определение пакета для текущего файла
import (
	"fmt"
	"os"
)

// ********* Stdin *********
// Объект os.Stdin реализует интерфейс io.Reader и позволяет считывать данные с консоли
// например, мы можем использовать функцию fmt.Fscan() для считывания с консоли с помощью os.Stdin

func main() {
	var name string
	var age int

	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	fmt.Println(name, age)

	// альтернативный вариант
	// fmt.Println("Введите имя и возраст:")
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Println(name, age)
}
