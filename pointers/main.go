package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	//&myString ile myString variable nın memoryde yerine referans veryrm
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	//:= ile direk değişken tanımlayablyrz. diğer türlü önce tanımlayıp sonra değer atamamız gerekyrdu
	*s = newValue
}

//*string type of pointer string
//değişkenin memorydeki yerine referans verme için & kullanıyrz
//point etmek için de * kullanıyrz
//changeUsingPointer func da bir değer dönmedik ama myString değişkenine pointer ettiğimiz için onu değiştirebildik
