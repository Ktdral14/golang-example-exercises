// ---ESTRUCTURA DE UN PROGRAMA EN GO---

// En Go es necesario declararar un paquete main el cual contendrá
// la función principal del programa y siempre debe de ir al inicio
// del programa

package main

// El paquete fmt es el paquete estándar de Go
// contiene todas las funciones básicas

import "fmt"

// La función main es la función principal de nuestro programa, contiene el
// hilo principal de nuestro programa y es la primera función que se ejecuta
// al ejecutar un programa de Go

func main() {
	fmt.Print("Hi world!")
}

// Para ejecutar este programa es necesario:
// 1 - Abrir la consola
// 2 - Ubicarse en la carpeta del archivo
// 3 - Ejecutar el comando << go run main.go >>
