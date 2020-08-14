// ---ESTRUCTURAS ITERATIVAS---

package main

import "fmt"

// En Go solo existe la palabra reservada for para hacer ciclos, pero viene supercargada
// para no echar de menos a las otras

func main() {

	// Forma 1: Esta forma actua como si fuera un ciclo while
	// for condicionBooleana {
	//		CÓDIGO
	// }
	a := 1
	for a < 10 {
		fmt.Println(a)
		a++
	}

	// Forma 2: Estructura clásica de un for, pero sin los paréntesis
	// for variable := valor; condicionBooleana; incremento {
	// 		CODIGO
	// ]
	for b := 1; b < 10; b++ {
		fmt.Println(b)
	}

	// Forma 3: Ciclo infinito. Si escribimos la instrucción for sin más ocasionará
	// que se genere un ciclo infinito.
	// for {
	// 		CODIGO
	// ]
	numero := 0
	c := 1
	for {
		fmt.Println("Repeticion: ", c)
		fmt.Println("Ingresa el número secreto")
		fmt.Scanln(&numero)
		c++
		if numero == 100 {
			// Podemos terminar cualquier ciclo con la palabra reservada break
			break
		}
	}

	var d = 0
	for d < 10 {
		fmt.Printf("\n Valor de d: %d", d)
		if d == 5 {
			fmt.Printf("   Multiplicacion por 2 \n")
			d = d * 2
			// La instrucción continue salta a una nueva iteración del for omitiendo todo el codigo
			// posterior a esta instrucción
			continue
		}
		fmt.Printf("   Pasó por aqui \n")
		d++
	}

	// Otra forma de hacer ciclos es usando a instrucción goto LABEL
	// Para poder usarla es necesario definir labels en nuestro código
	// de la siguiente manera:
	// NOMBRELABEL:
	var e int = 0

RUTINA:
	for e < 10 {
		if e == 4 {
			e = e + 2
			fmt.Println("Voy a rutina sumando 2 a e")
			// La instrucción goto nos regresará al punto en nuestro código
			// posterior a la declaración de a etiqueta
			goto RUTINA
		}
		// Con Printf podemos usar verbos (como %d por ejemplo) para poder escribir todo el texto
		// de nuestro mensaje de manera seguida, después separamos por coma y escribimos todas nuestras
		// variables en orden
		// Algunos verbos son:
		// * %d - Decimal, se usa con números
		// * %s - String, se usa para cadenas de texto
		// * %t - Se usa para valores booleanos
		// * %f - Valores de punto flotante
		// * %v - Cualquier tipo de valor
		// Se usa %% para escribir un signo de tanto porciento (%) literal
		fmt.Printf("Valor de e: %d\n", e)
		e++
	}
}
