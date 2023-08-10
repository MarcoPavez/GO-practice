package main

// El punto de entrada para la aplicación de GO
// Lo que queremos hacer ahora es crear un servicio Web dentro de main()
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Lo que se usa en GO es el paquete HTTP, el cual contiene muchas capacidades para la creación de servicios

	// HandleFunc registra una función a un path en un ServidorMux (DefaultServeMux)
	// ServeMux is an HTTP request multiplexer (an Http Handler, pero hace más que eso):
	// It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL. (https://pkg.go.dev/net/http#ServeMux)

	// Uno de los parámetros de HandleFunc es 'handler', que es un tipo en GO y también una interfaz que tiene un método ServeHttp(ResponseWriter, *Request), por lo tanto, HandleFunc implementa el método de la interfaz Handler.

	// ResponseWriter puede modificar el header, el statuscode, entre otros.
	// Request puede modificar el path, el metodo, el encoding, el body del request etc.
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// Leemos todo desde el body con ioutil.ReadAll() y lo asignamos a la variable data
		data, err := ioutil.ReadAll(r.Body)

		// Los errores se manejan utilizando el método Error del paquete http, y Error solicita un responseWriter, el mensaje de error y el statusCode
		if err != nil {
			http.Error(rw, "Jodiste wacho", http.StatusBadRequest)
			return
		}
		// Ahora, ¿cómo entregamos la respuesta al usuario? ¿Cómo modificamos la respuesta? Eso se realiza utilizando fmt.Fprintf y la interfaz ResponseWriter
		// fmt.Fprintf le da formato y lo escribe en el stream correspondiente (ResponseWriter en este caso)
		fmt.Fprintf(rw, "Hello %s", data)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// nil utiliza el defaultServeMux
	http.ListenAndServe(":9090", nil)
}
