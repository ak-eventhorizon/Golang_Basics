package main // определение пакета для текущего файла
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// ********* net/http.Client *********
// Для осуществления HTTP-запросов также может применяться структура http.Client
// Структура http.Client имеет ряд полей, которые позволяют настроить ее поведение

func main() {
	client := &http.Client{
		Timeout: 6 * time.Second,
	}

	req, err := http.NewRequest(
		"GET", "https://example.com", nil,
	)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
