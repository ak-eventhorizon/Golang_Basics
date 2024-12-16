package main // определение пакета для текущего файла
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// ********* net/http *********
// все соответствующие функции по работе с http были выделены в отдельный пакет net/http
// для отправки запросов в пакете net/http определен ряд функций: Get(), Post(), Head(), PostForm()

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// for {
	// 	buffer := make([]byte, 1024)
	// 	n, err := resp.Body.Read(buffer)
	// 	if n == 0 || err != nil {
	// 		break
	// 	}
	// 	fmt.Println(string(buffer[:n]))
	// }

	io.Copy(os.Stdout, resp.Body)
}
