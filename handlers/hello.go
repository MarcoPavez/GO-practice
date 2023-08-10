package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	// Declaramos un field
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Uno de los parámetros de HandleFunc es 'handler', interfaz que tiene un método ServeHttp(ResponseWriter, *Request), por lo tanto, HandleFunc implementa el método de la interfaz Handler.
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// log.Println() se reemplaza por:
	h.l.Println("Hello World")
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
}
