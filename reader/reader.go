package main // определение пакета для текущего файла
import (
	"fmt"
	"io"
	"strings"
)

// ********* Reader *********
// Пакет io определяет интерфейс io.Reader, который представляет чтение потока данных
// Интерфейс io.Reader имеет метод Read, который заполняет слайс байтов данными и возвращает количество прочитанных байтов и ошибку
// Когда поток для чтения заканчивается - возвращается ошибка io.EOF (end of file)

// func (T) Read(b []byte) (n int, err error)

func main() {
	r := strings.NewReader("Hello, Reader!") // reader

	b := make([]byte, 8) // buffer, через который reader может читать поток порциями по 8 байт

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
