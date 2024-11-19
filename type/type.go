package main // определение пакета для текущего файла
import "fmt"

// ********* Type *********
// Оператор type позволяет определять именнованный тип на основе другого.

type mile uint
type kilometer uint

// Или псевдоним для дугого типа
type PositiveInteger = uint

func distanceToEnemy(n mile) {
	// передаваемые данные должны быть явным образом определены как значение типа mile, а не типа uint или типа kilometer
	// это позволяет уменьшить вероятность передачи некорректных данных
	fmt.Println("Distance to enemy:", n, "mile(s)")
}

func checkUInt(v uint) {
	fmt.Println(v, "is uint type")
}

func main() {

	var dist1 mile = 15
	distanceToEnemy(dist1)

	// var dist2 uint = 13
	// distanceToEnemy(dist2) // будет ошибка

	// var dist3 kilometer = 22
	// distanceToEnemy(dist3) // будет ошибка

	// псевдоним НЕ определяет нового типа и все псевдонимы одного и того же типа считаются идентичными
	var v1 uint = 15
	checkUInt(v1) // норм

	var v2 PositiveInteger = 22
	checkUInt(v2) // норм

}
