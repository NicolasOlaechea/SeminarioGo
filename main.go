package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {
	cadena := "TX06ABCDEF"
	result, err := model.NewResult(cadena)
	if err == nil { //Si no hay error, muestro el resultado. Si hay error, muestro el error
		fmt.Println(result)
	} else {
		panic(err)
	}
}

/*
Corré los siguientes comandos:

go test ./... -coverprofile=test.out -v -coverpkg=./...
go tool cover -html=test.out

El primer comando ejecuta los test y genera un resultado en el file test.out y el segundo comando toma
ese resultado y lo muestra como html.
*/

/*
panic: Errores imprevistos que hacen que el programa finalice de forma espontánea y que se cierre el programa Go en ejecución.
*/
