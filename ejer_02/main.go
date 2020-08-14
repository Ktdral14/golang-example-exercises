// ---TIPOS DE DATOS Y VARIABLES EN GO---

package main

// De esta forma podemos importar varios paquetes

// El paquete strconv tiene varias funciones muy útiles para convertir strings
// en otro tipos de datos

import (
	"fmt"
	"strconv"
)

// En go se declaran variables globales de la siguiente manera:

// var nombreVariale tipoDeDato = valor

// En donde el tipo de dato y la asignación de un valor son opcionales
// Los tipos de datos pueden ser:
// int
// uint
// bool
// string
// int32
// int64
// float32
// float64
// uint32
// uint32
// Los tipos de datos terminados en 32 o 64 hacen referencia a la cantidad máxima de 
// bits que se reservarán en memoria para esa variable

// Cuando creamos variables sin un valor, Go les asigna un valor por defecto, evitando así el uso del null
// Los enteros y flotantes se inician en 0
// Los booleanos se incian en false
// los strings se inician en comillas vacias ("")
var numero int
var texto string
var status bool = true

// Dentro de una funcion podemos crear variables locales de la misma forma que una variable global.
// Pero además podemos usar el operador := para crear una variable y asignarle su valor al mismo tiempo
// en este caso Go asume el tipo de dato segun el valor asignado.

// Además Go tiene la capacidad de asignar multiples valores a multiples variables como se ve en el ejemplo.
func main() {
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Este es mi texto", false
	var variable 
	var moneda float64 = 0

	// Go es un lenguaje fuertemente tipado
	// No es posible pasar de un tipo de dato a otro con un simple asignación
	// Tampoco es posible pasar de un int64 a un int (por ejemplo) de esta manera
	// Por lo que es necesario usar funciones para lograr esto
	// Los tipos de datos no solo funcionan como lo que son, sino que también son funciones que permiten
	// convertir entre tipos de datos
	numero2 = int(moneda)

	// Para convertir datos numéricos a string no es posible hacerlo solo poniendo string(valor)
	// Hay varias funciones para esto una de ellas es Itoa del paquete strconv
	// Notese que como moneda es float64 es necesario convertirlo a int primero
	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)
	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(status)
}
