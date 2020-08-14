// --- PANIC, DEFER Y RECOVER ---

// Los panic y defer es la manera de Go de tratar las excepciones
// haciendo un paralelismo con java, un panic es como un throw,
// defer es como un finally
// y un defer junto con el recover es como un catch

// Panic genera excepciones que nosotros como desarrolladores deseemos
// El defer permite que una función se ejecute
// El recover permite recuperarse de las excepciones

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	// Defer necesita una función que será la que se ejecute al final
	// del código, independientemente de donde se declare
	// defer funciónAEjecutar
	defer f.Close()

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}

	ejemploPanic()
}

func ejemploPanic() {
	// Para simular un catch se usa defer en conjunto con
	// una función anónima de la siguiente manera:
	// defer func() {
	// 		código en caso de excepción
	// }
	defer func() {
		// Recover es una función que trata de recuperar el estado del programa
		// para continuar con la ejecución del código
		reco := recover()
		if reco != nil {
			log.Fatalf("Se recuperó de un panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		// Para usar panic solo se llama como si fuera una función
		// panic("mensaje de error")
		panic("Se encontró el valor de 1")
	}
}
