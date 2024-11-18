package main // определение пакета для текущего файла
import "fmt"

// ********* Struct Nested*********
// Поля одних структур могут представлять другие структуры

type contact struct {
	email string
	phone string
}

type account struct {
	login  string
	domain string
}

type person struct {
	name        string
	age         int
	contactInfo contact
	account     // сокращенное определение - фактически эквивалент свойству account account
}

func main() {

	var tom = person{
		name: "Thomas",
		age:  43,
		contactInfo: contact{
			email: "thomas@gmail.com",
			phone: "+5556789012",
		},
		account: account{
			login:  "t.anderson",
			domain: "earth.local",
		},
	}

	fmt.Println(tom.contactInfo.email)

	tom.contactInfo.phone = "+1234567890"
	fmt.Println(tom.contactInfo.phone)

	fmt.Println(tom.account.login)
	fmt.Println(tom.login) // эквивалентно предыщей форме записи
}
