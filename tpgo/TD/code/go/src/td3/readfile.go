package main

import (
	"io/ioutil"
	"log"
)

func main() {
	var filePath string = "/var/home/E256190S/reseau/Perso/Bureau/BUT1/tpgo/TD/code/go/src/td3/table.txt"
	var data []byte
	var err error
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("J'ai lu : ", string(data))
}
