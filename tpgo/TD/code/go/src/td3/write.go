package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var myFile *os.File
	var err error
	myFile, err = os.Create("ouiouibaguette.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(myFile, "BacharElSalade")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(myFile, "Parce que c'est notre projet")
	if err != nil {
		log.Fatal(err)
	}

	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
