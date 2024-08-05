package symphonycloudtools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
)

func HashChecker(hash string) bool {
	// Obtener todas las interfaces de red
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	// Iterar sobre cada interfaz y obtener su dirección MAC
	for _, iface := range interfaces {
		mac := iface.HardwareAddr.String()
		if mac != "" {
			// Crear el hash de la dirección MAC
			hashMAC := hashmac(mac)
			// Comparar el hash de la dirección MAC con el hash proporcionado
			if hashMAC == hash {
				return true
			}
		}
	}
	return false
}

func HashGenerator() (string, error) {
	// Obtener todas las interfaces de red
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	var MAC string
	// Iterar sobre cada interfaz y obtener su dirección MAC
	for _, iface := range interfaces {
		mac := iface.HardwareAddr.String()
		if mac != "" {
			MAC = mac
			break
		}
	}
	// Crear el hash de la dirección MAC
	hash := hashmac(MAC)
	return hash, nil
}

func hashmac(mac string) string {
	// Crear el hash de la variable mac
	// Aquí puedes usar cualquier algoritmo de hash, como MD5, SHA1, SHA256, etc.
	// Por ejemplo, usando el algoritmo SHA256:
	hash := md5.Sum([]byte(mac))
	// Convertir el hash a una cadena hexadecimal
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
