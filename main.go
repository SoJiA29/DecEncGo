package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Bienvenido al programa de encriptación/desencriptación.")

	// Loop hasta que el usuario escriba "salir"
	for {
		fmt.Print("\nIngrese un texto para encriptar o desencriptar (o escriba 'salir' para salir): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		texto := scanner.Text()

		if texto == "salir" {
			break
		}

		fmt.Print("Seleccione una opción: 1 para encriptar, 2 para desencriptar: ")
		scanner.Scan()
		opcion := scanner.Text()

		if opcion == "1" {
			fmt.Print("Ingrese una clave para encriptar: ")
			scanner.Scan()
			clave := scanner.Text()
			for len(clave) < 32 {
				clave += " " // agrega espacios hasta que la clave tenga 32 bytes
			}
			clave_en_bytes := []byte(clave)[:32]
			block, _ := aes.NewCipher(clave_en_bytes)
			nonce := make([]byte, 12)
			aesgcm, _ := cipher.NewGCM(block)
			texto_en_bytes := aesgcm.Seal(nil, nonce, []byte(texto), nil)
			fmt.Printf("El texto encriptado es: %x\n", texto_en_bytes)
		} else if opcion == "2" {
			fmt.Print("Ingrese la clave de encriptación: ")
			scanner.Scan()
			clave := scanner.Text()
			for len(clave) < 32 {
				clave += " " // agrega espacios hasta que la clave tenga 32 bytes
			}
			clave_en_bytes := []byte(clave)[:32]
			block, _ := aes.NewCipher(clave_en_bytes)
			nonce := make([]byte, 12)
			aesgcm, _ := cipher.NewGCM(block)
			texto_en_bytes, _ := hex.DecodeString(texto)
			texto_desencriptado, _ := aesgcm.Open(nil, nonce, texto_en_bytes, nil)
			fmt.Printf("El texto desencriptado es: %s\n", texto_desencriptado)
		} else {
			fmt.Println("Opción no válida. Por favor seleccione 1 o 2.")
		}
	}

	fmt.Println("Gracias por utilizar el programa. ¡Hasta pronto!")
}
