package main

import (
	"github.com/mrtarikunal/go-essentials/packages/helpers"
	//go mod init github.com/tsawler/myniceprogram ile kendi paketimizi olştrdk.
	"log"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}
