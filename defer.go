package main // определение пакета для текущего файла
import (
	"fmt" // подключение пакета fmt
)

// Оператор defer позволяет выполнить функцию в конце программы, при этом не важно, где в реальности вызывается эта функция
// Если несколько функций вызываются с оператором defer, то функции, которые вызываются раньше, будут выполняться позже всех.

func main() {
	defer finish()
	defer preFinish()
	fmt.Println("Step 1")
	fmt.Println("Step 2")
}

func finish() {
	fmt.Println("Step Finish")
}

func preFinish() {
	fmt.Println("Step Pre-Finish")
}
