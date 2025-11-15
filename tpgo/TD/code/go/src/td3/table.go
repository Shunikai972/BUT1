package main 
import (
	"fmt"
	"log"
	"os"
)

func main (){
	n := 9999999999999999999999999
	myFile, err := os.Create("table.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()

for i :=0; i<=999999999999999999999999999999;  i++{
	_, err = fmt.Fprint(myFile, i, "x", n, "=", i*n, "\n")
	if err != nil{
		log.Fatal(err)
	}	
}
}

 