// ---FUNCIONES---

package main

import "fmt"

// Podemos crear funciones de diversar formas en Go
// Go no permite la sobrecarga, por lo que solo podemos tener una función
// con el mismo nombre

func main() {
	fmt.Println(uno(5))

	numero, estado := dos(2)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(tres(5))

	fmt.Println(cuatro(1, 2, 3, 4, 5, 6))
}

// La primera es la estructura más común para crear funciones
// func nombreFuncion(parametro tipoDeDato, etc tipoEtc...) tipoDeRetorno {
// 		CODIGO
// 		return valor
// }

func uno(numero int) int {
	return numero * 2
}

// En Go, la funciones pueden retornar más de un valor, solo es necesario definir el tipo
// de retorno de cada variable al inicio
// func nombreFuncion(parametro tipoDeDato, etc tipoEtc...) (retorno1, retorno2, etc...){
// 		CODIGO
// 		return valor1, valor2, etc...
// }
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// También podemos definir el nombre de la variable de retorno al inicio de la función
// (esto también aplica a multiples retornos) y solo es necesaria la instrucción return
// para devolver todos los valores
// func nombreFuncion(parametro tipoDeDato, etc tipoEtc...) (variable tipoDeRetorno){
// 		CODIGO
// 		return
// }
func tres(numero int) (z int) {
	z = numero * 3
	return
}

// Dada la imposibilidad de usar sobrecarga podemos definir que una funcion reciba un
// numero n de valores, agregando tres puntos antes del tipo de dato, a la hora de
// trabajar con estos valores dentro de la función, se comportarán como un arreglo
// func nombreFuncion(listaDeValores ...tipoDeDato) tipoDeRetorno{
// 		CODIGO
// 		return valor
// }
func cuatro(numero ...int) int {
	total := 0
	// Este es un for especial para recorrer arreglos, la instrucción range nos devuelve
	// el índice y el valor de manera iterativa en un arreglo
	// NOTA: Cuando guardamos valores en _ Go no los guarda en memoria,
	// es decir, son valores que no vamos a utilizar en nuestro código
	for _, num := range numero {
		total += num
	}
	return total
}
