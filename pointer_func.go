package main // определение пакета для текущего файла
import "fmt"

func main() {
	// ********* Pointer with Functions *********

	// По умолчанию все параметры передаются в функцию по значению (копия переменной)
	x := 15
	fmt.Println("Before(by value):", x)
	changeValue1(x)
	fmt.Println("After(by value):", x) // значение x не изменилось

	// В данном случае функция принимает указатель (адрес переменной)
	fmt.Println("Before(by pointer):", x)
	changeValue2(&x)
	fmt.Println("After(by pointer):", x) // значение x изменилось

}

// функция работает по значению
func changeValue1(val int) {
	val += 1
}

// функция работает по указателю
func changeValue2(val *int) {
	*val += 1
}
