// --- ARREGLOS, MATRICES Y SLICES ---

package main

import (
	"fmt"
)

// Para crear arreglos estáticos lo hacemos de la siguiente manera:
// var nombre [tamaño]tipoDeDato
var arr [10]int

// Y para crear matrices solo agregamos una dimensión más a nuestro arreglo:
// var nombre [filas][columnas]tipoDeDato
var matriz [5][7]int

// Al igual que con las variables, los arreglos son inicializados de forma automática,
// por ejemplo, un arreglo de enteros (int) estará inicializado en 0

func main() {

	// Para acceder a un elemento específico solo es necesario brindar el índice del
	// elemento al cual se quiere acceder
	arr[0] = 1
	arr[5] = 15
	fmt.Println(arr)

	// También es posible crear arreglos con todos sus valores definidos, encerrando estos
	// entre llaves
	arr2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr2)

	// Podemos recorrer arreglos de la forma tradicional con un for de i
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("element at: %d : ", i)
		fmt.Printf(" %d\n", arr2[i])
	}

	// Para acceder a los elementos de una matriz solo es necesario brindar el índice
	// de la fila y columna a la cual se quiere acceder
	matriz[3][5] = 1
	fmt.Println(matriz)

	// Los slices son las listas de Go, arreglos con tamaño dinámico
	// Se crean de la misma forma que los arreglos pero sin brindarles un tamaño por defecto
	slice := []int{2, 4, 5}
	fmt.Println(slice)
	variante2()
	variante3()
	variante4()
}

// Es posible crear slices a partir de un arreglo
func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	// Obtiene todos los elemenos desde el índice 3 hasta el final
	porcion1 := elementos[3:]
	// Obtiene todos los elementos desde el inicio hasta el índice 3
	porcion2 := elementos[:3]
	// Obtiene todos los elementos desde el índice 2 al índice 4
	porcion3 := elementos[2:4]

	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

// Otra forma de crear slices es mediante la función make
// La función make recibe 2 parámetros y un tercer parametro opcional
// - El primer parámetro es lo que creará, que en este caso es un slice ([]tipoDeDato)
// - El segundo parámetro es el tamaño inicial del slice
// - El tercer parámetro (opcional) es lq capacidad por si se rebasa el tamaño brindado
// - en el segundo parámetro, será el espacio que se reserve en memoria para el slice
func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

// Podemos crear un slice con la función make de tamaño 0 y capacidad 0
// Esto provocará que el tamaño y la capacidad aumenten dinámicamente conforme
// se agregan elementos al arreglo. El tamaño irá variando según el número de
// elementos, y la capacidad será la potencia de 2 mínima mayor al tamaño actual
// (por ejemplo, si creamos un slice de tamaño 100, su capacidad será de 128; y
// si aumentaramos ese tamaño hasta 129, la capacidad sería ahora de 256)
func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		// La función apend nos devuelve una copia de nuestro slice con el nuevo
		// elemento agregado al final
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
