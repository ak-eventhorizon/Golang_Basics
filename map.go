package main // определение пакета для текущего файла
import (
	"fmt"
)

func main() {
	// ********* Map *********

	var myMap = make(map[string]int)
	fmt.Println(myMap) //map[]

	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   3,
		"Alice": 4,
	}

	fmt.Println(people)        // map[Alice:4 Bob:2 Sam:3 Tom:1]
	fmt.Println(people["Bob"]) // 2

	people["Alice"] = 45
	fmt.Println(people["Alice"]) // 45

	// перебор в цикле
	for key, val := range people {
		fmt.Println(key, val)
	}

	// проверка элемента
	if val, isPresent := people["Nobody"]; isPresent {
		fmt.Println(val)
	} else {
		fmt.Println("No such element in map!")
	}

	// добавление элемента
	people["Vasiliy"] = 125
	fmt.Println(people) // map[Alice:45 Bob:2 Sam:3 Tom:1 Vasiliy:125]

	// удаление элемента
	delete(people, "Bob")
	fmt.Println(people) // map[Alice:45 Sam:3 Tom:1 Vasiliy:125]

}
