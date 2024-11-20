package main // определение пакета для текущего файла
import (
	"fmt"
	"io"
	"os"
)

// ********* I/O Files*********
// Все файлы в Go представлены типом os.File.
// Этот тип реализует ряд интерфейсов, например, io.Reader и io.Writer, которые позволяют читать содержимое файла и сохранять данные в файл.

func main() {

	// **************** СОЗДАНИЕ ФАЙЛА И ЗАПИСЬ В НЕГО ****************

	file1, err := os.Create("log.txt") // создаем файл
	if err != nil {                    // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выход из программы
	}
	defer file1.Close() // закрываем файл при завершении функции

	text := "LOG START: "

	file1.WriteString(text) // Для записи текстовой информации в файл можно применять метод WriteString()
	file1.WriteString("log message 1; ")
	file1.WriteString("log message 2; ")

	fmt.Println("\n", file1.Name(), "-- file created")
	fmt.Println("")

	// **************** ОТКРЫТИЕ И ЧТЕНИЕ ИЗ ФАЙЛА ****************

	file2, err := os.Open("log.txt") // открываем файл
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file2.Close() // закрываем файл при завершении функции

	buffer := make([]byte, 64) // подготовка слайса на 64 байта для считывания в него потока из файла

	// бесконечный цикл, считывающий содержимое файла, пока не получит ошибку Конец Файла (EOF)
	for {
		n, err := file2.Read(buffer) // прочитать содержимое файла в побайтно в буфер
		if err == io.EOF {
			break
		}

		// проходим по слайсу, выводя каждый его элемент в виде десятичного значения байта, отформатированного в 4 знака
		for _, v := range buffer[:n] {
			fmt.Printf("%4d", v)
		}

		fmt.Println("")

		// проходим по слайсу, выводя каждый его элемент в виде символа, соответсвующего коду байта, отформатированного в 4 знака
		for _, v := range string(buffer[:n]) {
			fmt.Printf("%4s", string(v))
		}
	}
}
