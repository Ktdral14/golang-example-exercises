// ---ESTRUCTURAS DE CONTROL---

package main

import "fmt"

var estado bool

func main() {

	// La estructa if en Go no lleva parentesis, pero sí se pueden usar parentesis para agrupar
	// expresiones booleanas como si de una ecuación algebraica se tratase
	// Hay dos formas de hacer un if en Go:

	// La primera forma es:
	// if condicionBooleana {
	// 		CODIGO
	// }

	// Es importante resaltar que el if, su condición y las llaves de apertura deben estar en
	// una misma línea, de la misma forma, el else y su llave de apaertura deben estar en la misma
	// linea en la que se cierra el if

	if estado {
		fmt.Println("El estado es true")
	} else {
		fmt.Println("El estado es false")
	}

	// La segunda forma de hacer un if es haciendo una asignación dentro del mismo if:

	// if variable = valor; condicionBooleana {
	// 		CÓDIGO
	// } else {
	// 		CÓDIGO
	// }

	if estado = false; estado == true {
		fmt.Println(estado)
	} else {
		fmt.Println("Es falso")
	}

	// Para hacer un switch tenemos las mismas reglas que en el if
	// pero con la estructura de un switch común dentro de las llaves
	// Cabe resaltar que en Go no es necesario el uso de la palabra reservada break
	// para finalizar un caso en un switch

	numero1 := 10

	switch numero1 {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")
	}

	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")
	}
}
