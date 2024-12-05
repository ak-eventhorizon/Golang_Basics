package main // определение пакета для текущего файла
import (
	"fmt"
	"net"
)

// ********* Net.Listen *********
// Для прослушивания и приема входящих запросов по сети в пакете net определена функция net.Listen

var dict = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

func main() {

	listener, err := net.Listen("tcp4", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")

	// в бесконечном цикле принимаем входящие подключения
	// клиентская утилита из cmd-> go run ..\client\client.go
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn) // запускаем горутину для обработки подключения клиента
	}
}

// обработка подключения
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// выделяем буфер в 4кб для считывания данных из запроса клиента
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		// данные из ввода пользователя
		source := string(input[0:n])
		// на основании полученных данных получаем из словаря перевод
		target, ok := dict[source]
		if ok == false { // если данные не найдены в словаре
			target = "undefined"
		}
		// выводим на консоль сервера диагностическую информацию
		fmt.Println(source, "-", target)
		// отправляем данные клиенту
		conn.Write([]byte(target))
	}
}
