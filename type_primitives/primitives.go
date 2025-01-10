package main // определение пакета для текущего файла
import (
	"fmt"
	"reflect"
)

// ********* Простые типы данных *********

func main() {

	var _ int8 = -1
	var _ uint8 = 2
	var _ byte = 3 // byte - синоним типа uint8, занимает в памяти 1 байт, конвенционально используется чтобы отличать байты от целых чисел
	var _ int16 = -4
	var _ uint16 = 5
	var _ int32 = -6
	var _ rune = 33 // rune - синоним типа int32, представляет Unicode кодировку символов, конвенционально используется чтобы отличать символьные значения от числовых
	var _ uint32 = 8
	var _ int64 = -9
	var _ uint64 = 10

	var _ int = 102  // универсальное представление типа int - длиной по меньшей мере 32 бита
	var _ uint = 105 // универсальное представление типа uint - длиной по меньшей мере 32 бита

	var _ float32 = 18.6
	var _ float64 = 3.14

	var _ complex64 = 1 + 2i
	var _ complex128 = 4 + 3i

	var _ bool = true
	var _ bool = false

	var _ string = "Hello"

	// ********* Определение типа *********

	// Способ 1 - флаг %T для форматной строки
	var amount float32 = 10.4
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)

	// Способ 2 - использование пакета reflect стандартной библиотеки
	var isFound bool = false
	fmt.Printf("variable isFound='%v' is of type %v \n", isFound, reflect.TypeOf(isFound))

	// Способ 3 - реализация функции с использованием type assertions
	var days int = 42
	fmt.Println(typeofObject(days))
}

func typeofObject(variable interface{}) string {
	switch variable.(type) {
	case int:
		return "int"
	case float32:
		return "float32"
	case bool:
		return "boolean"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
