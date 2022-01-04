package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

//printFirstName fonksiyonunu myStruct type na bağladık. yani onu point etti ve reciever oldu
//yani myStruct type nın bir fonksiyonu haline geldi

func main() {
	var myVar myStruct
	//myVar değişkenini myStruct typena eşitledik

	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}
	//myVar2 değişkenini myStruct typena eşitledik. kısa yolu

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
