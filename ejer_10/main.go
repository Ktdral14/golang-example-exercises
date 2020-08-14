// --- POO Y PAQUETES

// Para definir paquetes en go es necesario meter todos los archivos

package main

import (
	"fmt"
	"time"

	// Para importar un paquete personalizado es necesario poner el nombre
	// del paquete dentro de nuestro código y la ruta al paquete en el proyecto
	// nombrePaquete "ruta/paquete"
	user "./user"
)

// Para hacer herencia en Go solo es necesario llamar a la clase padre dentro de clase hija
// type ClaseHija {
// 		ClasePadre
// }
type kt struct {
	user.Usuario
}

// Nota en este caso utilizo user.Usuario porque la clase Usuario
// se encuentra dentro del paquete user

func main() {
	// La función new crea objetos de clases propias
	User := new(kt)
	User.AltaUsuario(10, "Angel", time.Now(), true)
	fmt.Println(User.Usuario)
}
