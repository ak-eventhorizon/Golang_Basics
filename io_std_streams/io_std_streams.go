package main // определение пакета для текущего файла
import (
	"fmt"
	"io"
	"os"
)

// ********* I/O StdStreams*********
// Пакет os определяет три переменных: os.Stdin, os.Stdout и os.Stderr, которые представляют стандартные потоки ввода, вывода и ошибок соответственно.
// Мы можем использовать функцию io.Copy() для копирования данных из одного потока в другой

func main() {

	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(os.Stdout, file) // передача содержимого файла в стандартный поток вывода системы
}
