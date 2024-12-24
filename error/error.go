package main // определение пакета для текущего файла
import (
	"fmt"
	"time"
)

// ********* Error *********
// Программы на Go выражают состояние ошибки значением типа error

// type error interface {
//     Error() string
// }

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v: %s", e.When.Format(time.DateTime), e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"что-то пошло не так!",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
