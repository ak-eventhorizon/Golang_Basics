package main // определение пакета для текущего файла
import (
	"fmt"
)

func main() {
	// ********* Массив (Array) *********

	var arr1 [5]int                         // [0 0 0 0 0]
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // [1 2 3 4 5]
	var arr3 [5]int = [5]int{1, 2}          // [1 2 0 0 0]
	arr4 := [5]int{7, 8, 9}                 // [7 8 9 0 0]
	arr5 := [...]int{7, 8, 9}               // [7 8 9]

	fmt.Println(arr1, arr2, arr3, arr4, arr5)

}
