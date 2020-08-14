// --- MANEJO DE ARCHIVOS ---

package main

// Para el manejo de archivos utilizamos los paquetes
// bufio, io/ioutil, os
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

// La primera forma de leer archivos es usando el paquete ioutil
func leoArchivo() {
	// Cuando se leen archivos siempre retorna dos valores
	// el archivo en sí y el error si e que ocurre
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error", err.Error())
	} else {
		fmt.Println(string(archivo))
	}
}

// El paquete os también nos permite abrir archivos
func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error", err.Error())
	} else {
		// Con el paquete bufio podemos crear un scanner para leer el archivo línea por linea
		scanner := bufio.NewScanner(archivo)
		// La funcion Scan() hace que el buffer lea una línea del archivo y retorna true mientras
		// haya líneas
		for i := 1; scanner.Scan(); i++ {
			registro := scanner.Text()
			fmt.Printf("Linea " + fmt.Sprint(i) + " > " + registro + "\n")
		}
	}
	// También es necesario cerrar el buffer del archivo
	archivo.Close()
}

func graboArchivo() {
	// Con la función Create del paquete os nosotros podemos crear archivos nuevos
	archivo, err := os.Create("./newFile.txt")
	if err != nil {
		fmt.Println("Hubo un error", err.Error())
		return
	}
	// La función Fprint del paquete fmt nos permite escribir contenido dentro de un archivo
	// solo es necesario pasarle el archivo y el contenido que se va a grabar
	// dentro del archivo
	fmt.Fprint(archivo, "Esta es un archivo nuevo")
	archivo.Close()
}

func graboArchivo2() {
	fileName := "./newFile.txt"
	if appendToFile(fileName, "\nEsta es una nueva linea") == false {
		fmt.Println("Error en la 2da linea")
	}

}

func appendToFile(file string, text string) bool {
	// Podemos mandar algunos parámetros adicionales para decir que el archivo será de solo escritura
	// y que el contenido se va a agregar al final del archivo
	archivo, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error", err.Error())
		return false
	}

	// La función WriteString nos permite escribir en archivos
	_, err = archivo.WriteString(text)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	archivo.Close()
	return true
}
