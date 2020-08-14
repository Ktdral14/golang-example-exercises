// --- RECURSIVIDAD EN GO ---

// Una función es recursiva cuando esta tiene una llamada a sí misma
// lo que genera un ciclo de llamadas hasta llegar al conocido
// "caso base"

// La recursividad en Go es igual a la de otros lenguajes de propósito general

package main

import "fmt"

func main() {
	exponencial(2)
}
func exponencial(numero int) {
	if numero > 1000000 {
		return
	}
	fmt.Println(numero)
	exponencial(numero * 2)
}
