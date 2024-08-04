package symphonycloudtools

import (
	"fmt"
	"net"
)

func Init() {
	fmt.Println("Init syr2uploader")
}

func HashGenerator() {
	// Obtener todas las interfaces de red
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Iterar sobre cada interfaz y obtener su dirección MAC
	for _, iface := range interfaces {
		mac := iface.HardwareAddr.String()
		if mac != "" {
			fmt.Printf("Interfaz: %s, MAC: %s\n", iface.Name, mac)
		}
	}
}
