// ---MIDDLEWARES---
/* Son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben y
devuelven los mismos tipos de variables */

package main

import "fmt"

var result int

func main() {
	fmt.Println("INICIO")

	// Le mandamamos la función que queremos ejecutar y los valores con los que se ejecutará esa funcion
	result = operacionesMid(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMid(restar)(20, 8)
	fmt.Println(result)
	result = operacionesMid(multiplicar)(5, 5)
	fmt.Println(result)
	result = operacionesMid(dividir)(9, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}
func dividir(a, b int) int {
	return a / b
}

// Un middleware no es más que una función que recibe otra función y retorna la función deseada
// Podría pensarse como un menú que decide que función se utilizará
func operacionesMid(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("===Inicio de operación===")
		return f(a, b)
	}
}
