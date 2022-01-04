package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse"}

	for i, animal := range animals {
		log.Println(i, animal)

	}

	vegetables := make(map[string]string)

	vegetables["first"] = "garlic"
	vegetables["second"] = "spinach"

	for vegetableType, vegetable := range vegetables {
		log.Println(vegetableType, vegetable)

	}

	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}
	//string aslında her karekter slice içinde tutulyr. ve aslında immutable.

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})
	log.Println(users)

	for _, l := range users {
		//i index olan kullanmıyorsan yereine _ koyablyrsn
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
