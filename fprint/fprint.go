package main // определение пакета для текущего файла
import (
	"fmt"
	"os"
)

// ********* Fprint *********
// Ряд возможностей по чтению и записи файлов предоставляет пакет fmt.
// Этот пакет предоставляет ряд функций для записи данных в произвольный объект, который реализует интерфейс io.Writer: fmt.Fprint(), fmt.Fprintln() и fmt.Fprintf()

type person struct {
	name   string
	age    int32
	weight float64
}

func main() {

	tom := person{
		name:   "Thomas",
		age:    34,
		weight: 68.5,
	}

	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprint(file, "Сегодня ")                                              // вывод в файл строки
	fmt.Fprintln(file, "хорошая погода")                                      // вывод в файл строки с переводом строки
	fmt.Fprintf(file, "|%-10s|%-10d|%-10.3f|", tom.name, tom.age, tom.weight) // вывод в файл форматной строки
}
