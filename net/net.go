package main // определение пакета для текущего файла
import (
	"fmt"
	"io"
	"net"
	"os"
)

// ********* Net *********
// Основной функционал по работе с сетью представлен пакетом net
// Этот пакет предоставляет различные низкоуровневые сетевые примитивы, через которые идет взамодействие по сети

func main() {
	httpRequest := "GET / HTTP/1.1\n" + "Host: golang.org\n\n"
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}
