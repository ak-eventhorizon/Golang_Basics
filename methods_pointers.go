package main // определение пакета для текущего файла
import "fmt"

// ********* Pointers Methods*********

type person struct {
	name string
	age  int
}

// При вызове метода, объект структуры, для которого определен метод, передается в него по значению (копия структуры)
func (p person) updateAge1(newAge int) {
	p.age = newAge
}

// Если необходимо изменять состояние структуры - следует использовать указатели на нее (адрес структуры типа person в памяти)
func (p *person) updateAge2(newAge int) {
	(*p).age = newAge // с помощью операции разыменования получаем значение по этому адресу в памяти и меняем поле age
}

func main() {

	mike := person{name: "Michael", age: 34}
	fmt.Println("Mike before", mike.age) // 34
	mike.updateAge1(41)
	fmt.Println("Mike after", mike.age) // 34 - не изменился

	var tom = person{name: "Tom", age: 20}
	var tomPointer *person = &tom
	fmt.Println("Tom before", tom.age) // 20
	tomPointer.updateAge2(33)
	fmt.Println("Tom after", tom.age) // 33 - изменился

}
