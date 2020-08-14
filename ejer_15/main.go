// --- GO ROUTINES (FUNCIONES ASÍNCRONAS EN GO) ---

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Para ejecutar una función de manera asíncrona solo es necesario
	// agregar la palabra reservada go antes del llamado a la función
	// go miFunción(parametros)
	go miNombreLento("angel")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
	// La función Split del paquete strings convierte una cadena en un arreglo
	// usando como separador la cadena dada
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		// La función Sleep pone en pausa la ejecución por el tiempo dado
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(letra)
	}
}
