package main // определение пакета для текущего файла
import "fmt"

// ********* Struct *********
// Структуры представляют тип данных, определяемый разработчиком и служащий для представления каких-либо объектов

type person struct {
	name string
	age  int
}

func main() {

	// варианты создание и инициализация структуры
	var vasya person = person{"Vasiliy", 25}
	var petya person = person{age: 43, name: "Peter"}
	bob := person{"Robert", 55}
	nobody := person{} // все свойства структуры получают значения по умолчанию

	fmt.Println(vasya.name, vasya.age)
	fmt.Println(petya.name, petya.age)
	fmt.Println(bob.name, bob.age)
	fmt.Println(nobody.name, nobody.age)

}
