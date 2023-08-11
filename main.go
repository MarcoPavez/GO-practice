package main

// El punto de entrada para la aplicación de GO
// Lo que queremos hacer ahora es crear un servicio Web dentro de main()
import (
	"context"
	"log"
	"modInit/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// HandleFunc registra una función a un path en un ServidorMux (DefaultServeMux)
	// ServeMux is an HTTP request multiplexer (an Http Handler, pero hace más que eso):
	// It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL. (https://pkg.go.dev/net/http#ServeMux)

	// ResponseWriter puede modificar el header, el statuscode, entre otros.
	// Request puede modificar el path, el metodo, el encoding, el body del request etc.
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// Pero, y como manejamos los timeouts de mis conexiones? creando un http Server()
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Cierre recibido, finalización segura", sig)
	// Para cierres de emergencia o cuando se requiere updatear la aplicacion, se utiliza shutdown()
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
