package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* Polymorphism *********

/* Полиморфизм - это способность принимать многообразные формы.
В данном примере - в зависимости от реального типа структуры, реализующей интерфейс Vehicle,
динамически определяется, какая именно реализация метода move должна вызываться. */

type Vehicle interface {
	move()
}

type Car struct{ model string }
type Aircraft struct{ model string }
type Ship struct{ model string }

func (c Car) move() {
	fmt.Println(c.model, "едет")
}
func (a Aircraft) move() {
	fmt.Println(a.model, "летит")
}
func (s Ship) move() {
	fmt.Println(s.model, "плывет")
}

func main() {

	tesla := Car{"Tesla"}
	volvo := Car{"Volvo"}
	boeing := Aircraft{"Boeing"}
	bismark := Ship{"Bismark"}

	vehicles := [...]Vehicle{tesla, bismark, volvo, boeing}
	for _, vehicle := range vehicles {
		vehicle.move()
	}
}
