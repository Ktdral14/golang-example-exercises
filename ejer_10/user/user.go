// Podemos crear paquetes personalizados con la palabra reservada package
// package nombrepaquete
// Cada paquete debe tener una carpeta única para todos sus archivos
// El nombre del paquete debe de coincidir con el nombre de la carpeta que lo contiene
// Los paquetes pueden ser útiles para crear código reutilizable y mantenible de manera sencilla
package user

import (
	"time"
)

// Podemos definir una especie de clase en Go con la palabra reservada struct
// de la siguiente manera:
// type NombreClase struct {
// 		propiedad1 tipoDeDato
// 		propiedadn tipoDeDato
// }
type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

// Dentro de un struct no se pueden definir métodos, pero podemos crear métodos para un struct
// específico haciendo uso de punteros
// func (this *NombreClase) nombreMétodo(parametros) {
// 		código random
// }
func (this *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status
}
