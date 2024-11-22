package main // определение пакета для текущего файла
import (
	"fmt"
	"io"
)

// ********* I/O *********
// Go имеет свою модель работы с потоками ввода-вывода, которая позволяет получать данные из различных источников - файлов, сетевых интерфейсов, объектов в памяти и т.д
// Поток данных в Go представлен байтовым срезом ([]byte), из которого можно считывать байты или в который можно заносить данные.
// Ключевыми типами для работы с потоками являются интерфейсы io.Reader и io.Writer

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}
	return count, io.EOF
}

type phoneWriter struct{}

func (p phoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}

	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]), " ")
		}
	}
	fmt.Println()
	return len(bs), nil
}

func main() {
	phone1 := phoneReader("+7(333)444 5566")
	phone2 := phoneReader("+8-495-777-88-99")

	buffer := make([]byte, len(phone1))
	phone1.Read(buffer)
	fmt.Println(string(buffer)) // 73334445566

	buffer = make([]byte, len(phone2))
	phone2.Read(buffer)
	fmt.Println(string(buffer)) // 84957778899

	bytes := []byte("+1(498)345-12-33")
	writer := phoneWriter{}
	writer.Write(bytes)
}
