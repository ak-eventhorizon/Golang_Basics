package main // определение пакета для текущего файла
import (
	"fmt"
)

// ********* Generics *********
// Обобщенные типы - с их помощью можно реализовывать конструкции, которые будут работать с любым типом данных из указанного набора

// определяем интерфейс с ограничением типов (type constraint) для возможности его переиспользования в разных участках кода
type Number interface {
	int64 | float64
}

func main() {
	// Инициализация map с целыми числами
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Инициализация map с числами с плавающей точкой
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	// тут явно указываются параметры типов для аргумента функции
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// тут опускается указание параметров типов, компилятор выведет их из аргументов самой функции
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// Складывает значения всех элементов m (типа int64)
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// Складывает значения всех элементов m (типа float64)
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Параметры типов, в дополнение к обычным аргументам функции - делают эту функцию generic, позволяя работать с аргументами разных типов.
// Сomparable - предопределенный интерфейс в Go, означает любой тип, значения которого поддерживают операции сравнения == и !=
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Используем интерфейс Number для ограничения типов (type constraint)
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
