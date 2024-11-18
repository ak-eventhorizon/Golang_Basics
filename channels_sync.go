package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* Channels Sync *********
// канал не обязательно должен нести данные, которые представляют некоторый результат
// это может быть холостой объект, например, пустая структура, которая необходима только для синхронизации горутин

func main() {

	results := make(map[int]int)
	structCh := make(chan struct{})

	go factorial(6, structCh, results)

	<-structCh // ожидание закрытия канала structCh - процесс main не пойдет дальше, пока канал не будет закрыт

	for i, v := range results {
		fmt.Println(i, " - ", v)
	}
}

// функция не отправляет конкретные данные в канал, а просто закрывает его после выполнения всех своих инструкций
func factorial(n int, ch chan struct{}, results map[int]int) {
	defer close(ch) // закрываем канал после завершения горутины
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}
