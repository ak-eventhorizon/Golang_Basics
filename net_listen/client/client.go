package main // определение пакета для текущего файла
import (
	"fmt"
	"io"
	"net"
	"os"
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

	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone")
}
