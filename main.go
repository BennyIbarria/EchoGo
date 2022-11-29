package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	//Instanciamos echo
	e := echo.New()
	//Creamos la ruta / y le pasamos una función anonima
	e.GET("/", func(c echo.Context) error {
		//Creamos un objeto de tipo http.StatusOK y le enviamos el mensaje que queremos imprimir en pantalla
		return c.String(http.StatusOK, "Hola mundo desde Golang!")
	})
	//Le decimos a echo que escuche en el puerto 1323
	e.Logger.Fatal(e.Start(":1323"))
}
