package main // определение пакета для текущего файла
import "fmt"

// ********* Struct Methods*********
// Метод представляет функцию, связанную с определенным типом, например со структурой

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println(p.name, "is", p.age, "years old")
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}

func main() {

	mike := person{name: "Michael", age: 34}
	mike.print()
	mike.eat("борщ")

}
