// ---FUNCIONES ANONIMAS Y CLOSURES---

// Las funciones anonimas son, como su nombre lo indica, funciones
// que no tienen nombre. Su mayor ventaja es que podemos modificar 
// su comportamiento en tiempo de ejecución

package main

import "fmt"

// Las funciones anónimas se declaran como variables
// En este caso se declaran de la siguiente manera:
// var nombre func(parámetros) retorno = func(parametros) retorno {
// 		CODIGO
// 		return
// }
// Lo que se encuentra después del signo de igual es opcional
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma de 5 + 7 = %d \n", calculo(5, 7))

	// Como se ve, es posible modificar el comportamiendo de la función
	// mientras se va ejecutando el código. La única regla es que siempre
	// debe de respetarse la cantidad de parámetro y el tipo de retorno de la
	// declaración inicial
	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resta de 6 - 4 = %d \n", calculo(6, 4))

	calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Division de 12 / 3 = %d \n", calculo(12, 3))

	operaciones()

	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(miTabla())
	}
}

// Por otro lado los closures son funciones que retornan 
// funciones anónimas lo cual nos da un nivel extra de encapsulamiento.
// Puesto que la función principal no sabe lo que ocurre dentro del closure.
// Y ps la neta no doy más explicación poruque no termino de entenderlo del todo xd

func operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 14
		return a + b
	}
	fmt.Println(resultado())
}

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
