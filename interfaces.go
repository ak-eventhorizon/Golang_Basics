package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* Interfaces *********
// Интерфейсы представляют абстракцию поведения других типов - они позволяют определять функции, которые не привязаны к конкретной реализации.
// То есть интерфейсы определяют некоторый функционал, но не реализуют его.

type Vehicle interface {
	move()
}

func drive(v Vehicle) {
	v.move()
}

// структура Автомобиль, неявно реализует интерфейс Vehicle, поскольку реализует все его методы
type Car struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}

// структура Самолет, неявно реализует интерфейс Vehicle, поскольку реализует все его методы
type Aircraft struct{}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {

	tesla := Car{}
	boing := Aircraft{}

	drive(tesla)
	drive(boing)

}
