// --- SERVIDOR WEB CON GO ---

package main

// Para crear nuestro servidor necesitaremos el paquete http
import (
	"net/http"
)

func main() {
	// La función HandleFunc nos ayuda a definir las rutas de la aplicación
	// recibe la ruta y una función que nos devolverá el archivo de nuestro página
	http.HandleFunc("/", home)
	// La función ListenAndServe deja un puerto escuchando para levantar el servidor web
	http.ListenAndServe(":3000", nil)
}

// Esta función nos ayuda a mostrar archivos html de nuestra aplicación web
// La estructura de este tipo de funciones siempre es la misma, solo cambiando
// el archivo que abre
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
