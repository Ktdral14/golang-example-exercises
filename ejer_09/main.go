// --- MAPAS ---

package main

import "fmt"

func main() {
	// Podemos usar la función make para crear mapas
	// En este caso el mapa se declara de la siguiente manera:
	// (map[clave]valor, tamaño inicial)
	paises := make(map[string]string, 5)

	fmt.Println(paises)

	// Se accede al valor de un mapa de la siguiente manera:
	// mapa[clave] = nuevoValor
	paises["Mexico"] = "D. F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)

	// También podemos crear mapas directamente con una asignación
	// en este caso la sintaxis es identica a usar la función make
	// solo agregamos unas llaves al final:
	// variable := map[clave]valor{
	// 		clave: valor
	// }
	// NOTA: De este modo también se puede crear un mapa vacío
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	campeonato["Riverplate"] = 25
	campeonato["Chivas"] = 55

	// La función delete nos permite borrar un valor de un mapa
	// delete(mapa, clave)
	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

	// Podemos iterar un mapa de la siguiente manera
	// for variableClave, variableValor := range campeonato {
	// 		código random
	// }
	// La palabra reservada range hace que el mapa se comportara como si fuera un arreglo
	// y siempre nos da diferentes resultados ya que no tiene un criterio específico para
	// recorrer el mapa
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s tiene un puntaje de %d \n", equipo, puntaje)
	}

	// Podemos obtener la existencia de un valor de la siguiente manera:
	// valor, boolParaComprobarExistencia := mapa[clave]
	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

	puntaje2, existe2 := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje2, existe2)
}
