package main // определение пакета для текущего файла
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ********* Bufio *********
// Большиство встроенных операций ввода-вывода не используют буфер. Это может иметь отрицательный эффект для производительности приложения
// Для буферизации потоков чтения и записи в Go определены ряд возможностей, которые сосредоточены в пакете bufio

func main() {
	writeFile()
	readFile()
}

func writeFile() {
	rows := []string{
		"Hello!",
		"Welcome to Golang",
		"This is",
		"Example of",
		"Bufio usage",
		"With file",
	}

	file, err := os.Create("some.dat")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Пишем в файл через буфер
	writer := bufio.NewWriter(file)
	for _, row := range rows {
		writer.WriteString(row)  // запись строки
		writer.WriteString("\n") // перевод строки
	}
	writer.Flush() // сбрасываем данные из буфера в файл
	fmt.Println("File saved!")
}

func readFile() {
	file, err := os.Open("some.dat")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	// Читаем из файла через буфер
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // разделителем строк считать перевод строки
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Print(line)
	}
}
