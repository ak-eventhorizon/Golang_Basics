package main // определение пакета для текущего файла
import (
	"fmt"
	"net"
	"time"
)

// ********* Net.Listen *********
// Клиентская часть для проверки сервера, запускаемого из cmd-> go run ..\server\server.go

func main() {

	conn, err := net.Dial("tcp4", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		var source string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}

		// завершение работы клиента, если пользователь ввел "exit"
		if source == "exit" {
			break
		}

		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		// получаем ответ
		fmt.Print("Перевод: ")
		conn.SetReadDeadline(time.Now().Add(time.Second * 5)) // установка начального таймаута - клиент может ожидать данные от сервера в течении 5 сек, по истечении этого времени операция чтения сгенерирует ошибку
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[0:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 200)) // установка таймаута на каждую следующую итерацию цикла получения данных от сервера
		}
		fmt.Println()
	}
}
