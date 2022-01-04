package main

import (
	"log"

	"github.com/mrtarikunal/go-essentials/channels/helpers"
)

const numPool = 1000

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
	//oluşturduğum random sayıyı intChan(channel olarak tanımlanan) değişkinine aktardım
}

func main() {
	intChan := make(chan int)
	//make(chan int) channel olştrr. channel lar programın bir yerinden başka yerine bilgi iletmek için kullanlr
	defer close(intChan)
	//intChan kullanımı biter bitmez bu channelı kapat

	go calculateValue(intChan)
	//go routine gerçekleşyr. aslında önce channel tanımlyrz. sonra burda o channelı uygulatyrz

	num := <-intChan
	//burda channeldan gelen değeri dinliyrz. ve num değişkenine aktyrz
	log.Println(num)
}
