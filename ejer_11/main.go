// --- INTERFACES ---

// Las interfaces nos ayudan a tener un esqueleto que puede definir el
// comportamiento de cualquier cosa
// Pensemos, por ejemplo, en un animal existen infinidad de especies de animales
// pero todos tienen algunos comportamientos en común, como el hecho de respirar
// (ya sea con pulmones o con branquias) o el mero hecho de estar vivo ya es un comportamiento
// por lo que podríamos decir que un animal es una interfaz que define el comportamiento
// de todas las especies de animales
package main

import (
	"fmt"
)

// Para definir interfaces en Go usamos la siguiente sintaxis
// type NombreInterfaz interface {
// 		nombreMetodo(parametros) retorno
// }
// NOTA: Puede o no puede haber un tipo de retorno (void) y parametros
type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

// Para que una clase implemente una interfaz solo es necesario redefinir
// todos los métodos con un puntero a la clase que implementa la interfaz
// La estructura es la misma que la de un método de una clase
func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	}
	return "Mujer"
}
func (h *hombre) estaVivo() bool { return h.vivo }

func humanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un(a) %s, y ya estoy respirando \n", hu.sexo())
}

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) esCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func animalesRespirar(an animal) {
	an.respirar()
	fmt.Printf("Soy un animal y estoy respirando \n")
}

func animalesCarnivoros(an animal) int {
	if an.esCarnivoro() {
		return 1
	}
	return 0
}

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

func main() {
	kt := new(hombre)
	kt.esHombre = true
	humanosRespirando(kt)
	fmt.Printf("Estoy vivo = %t \n", estoyVivo(kt))

	sara := new(mujer)
	humanosRespirando(sara)
	fmt.Printf("Estoy vivo = %t \n", estoyVivo(sara))

	totalCarnivoros := 0
	mila := new(perro)
	mila.carnivoro = true
	mila.vivo = true
	animalesRespirar(mila)
	totalCarnivoros = +animalesCarnivoros(mila)
	fmt.Printf("Total de carnivoros: %d \n", totalCarnivoros)

	fmt.Printf("Estoy vivo = %t \n", estoyVivo(mila))
}
