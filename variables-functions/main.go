package main

import "fmt"

//burda bir değişken tanımlarsam main dışındanda ulaşablyrm. global scope.
//ayrıca değişkenin ilk harfi büyük olursa tüm paketlerede ulaşablyr
//kullanılmayan değişken tanımlayamazsın

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	//whatWasSaid ve theOtherThingThatWasSaid değişkinlerini tanıma ve saySomething func dan dönen değerlere eşitle

	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}

//iki tane string dönüyor
