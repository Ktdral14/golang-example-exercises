// ---ENTRADA POR TECLADO---

package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {

	// Para leer por teclado podemos usar alguna de las funciones Scan como Scanln
	// del paquete fmt
	// Para leer usando las funciones Scan es necesario pasarle el puntero de la variable
	// a la función, el puntero se define con & antes de la variable
	// Esto hace más rápido el proceso de lectura por teclado
	// NOTA: En windows, la funcioón Scanf reconoce el enter como una entrada de teclado
	// por lo que es recomendable usar Scanln en su lugar
	fmt.Println("Ingrese un numero: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese otro numero: ")
	fmt.Scanln(&numero2)

	// Otra forma de leer por teclado es usando un objeto Scanner del paquete bufio
	// para crear este objeto es necesario pasarle un buffer de entrada, el cual se obtiene
	// con la instrucción Stdin del paquete os que devuelve la entrada por teclado de sistema
	fmt.Println("Descripción: ")
	scanner := bufio.NewScanner(os.Stdin)

	// Es necesario evaluar en un if que se está entrando algo por teclado antes de hacer una
	// una asignación de valores, esto se hace con la función Scan del objeto
	if scanner.Scan() {
		// Con la función Text obtenemos lo que se ingresó para guardarlo en una variable
		leyenda = scanner.Text()
	}

	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
