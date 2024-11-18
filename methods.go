package main // определение пакета для текущего файла
import "fmt"

// ********* Methods*********
// Метод представляет функцию, связанную с определенным типом

type library []string // определяем тип library, являющийся срезом строк

func (lib library) print() { // определяем метод print для типа library
	for _, v := range lib {
		fmt.Println(v)
	}
}

func main() {

	var myLib library = []string{"book1", "book2", "book3"}
	myLib.print()

}
