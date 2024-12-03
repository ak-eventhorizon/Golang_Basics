package main // определение пакета для текущего файла
import (
	"fmt"
	"net"
)

// ********* Net.Listen *********
// Для прослушивания и приема входящих запросов по сети в пакете net определена функция net.Listen

func main() {

	message := "Hello, I am a GOLANG server" // отправляемое сообщение

	listener, err := net.Listen("tcp4", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")

	// в бесконечном цикле принимаем входящие подключения
	// можно проверить подключение: из cmd-> telnet localhost 4545
	// можно проверить подключение: клиентской утилитой из cmd-> go run ..\client\client.go
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
