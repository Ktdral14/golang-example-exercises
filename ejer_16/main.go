// --- CHANNELS ---

// Los channels nos permiten enviar información entre funciones que se ejecutan
// de manera asíncrona

package main

import (
	"fmt"
	"time"
)

func main() {
	// Para crear channels es necesaria la función make
	// channelVar := make(chan tipoDeDato)
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegó hasta aquí")

	// Aquí se obtiene la información guaradada en el channnel
	// variableParaGuardarInfo := <-channelVar
	msg := <-canal1

	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000000000; i++ {
	}
	final := time.Now()

	// Aquí se guarda información dentro del channel
	// channelVar <- valorAGuardar
	canal1 <- final.Sub(inicio)
	// La funcion sub nos sirve para restar tiempo
}
