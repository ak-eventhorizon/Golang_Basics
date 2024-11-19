package main // определение пакета для текущего файла
import (
	"fmt"
	// "slices" // библиотека для работы со срезами
)

func main() {
	// ********* Срез (Slice) *********

	// Срез определяется также, как и массив, за тем исключением, что у него не указывается длина
	var slice1 []string                  // []
	var slice2 = []string{"a", "b", "c"} // [a b c]
	slice3 := []string{"x", "y", "z"}    // [x y z]
	var slice4 []int = make([]int, 5)    // [0 0 0 0 0]

	fmt.Println(slice1, slice2, slice3, slice4)

	// добавление в срез
	letters := []string{"a", "b", "c"}       // [a b c]
	letters = append(letters, "x", "y", "z") // [a b c x y z]
	fmt.Println(letters)

	// оператор среза
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6] // с 2-го индекса по 6-й (невключительно)
	users2 := initialUsers[:4]  // с начала по 4-й индекс (невключительно)
	users3 := initialUsers[3:]  // с 3-го индекса до конца

	fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
	fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
	fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]

	// удаление элемента из среза (например с индексом 3 - "Sam")
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

}
