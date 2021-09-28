package model

import (
	"errors"
	"strconv"
	"unicode"
)

//Creo la estructura
type Result struct {
	Type   string
	Length int
	Value  string
}

//Funcion que retorna una instancia de Result y le pone los valores que recibe por parametro
func NewResult(cadena string) (Result, error) {
	min := 4 //Variable con el valor minimo que tiene que tener la cadena para ser valida.
	longitudCadena := len(cadena)
	if longitudCadena >= min {
		resultado := modificarValores(cadena)                         //Modifico los valores de la estructura y la guardo en una variable
		if verificarTipo(resultado) && verificarLongitud(resultado) { //Si es string o int y ademas la longitud es valida...
			return resultado, nil //Retorno la nueva estructura
		}
	}
	return Result{}, errors.New("cadena invalida") //Si no se comple lo anterior, retorno el error
}

func verificarTipo(resultado Result) bool {
	tipoNumerico := "NN" //Variable para identificar el tipo numerico

	//Retorno si el tipo del resultado que recibo por parametro es numerico o string
	if resultado.Type == tipoNumerico {
		return esEntero(resultado.Value) //Funcion que verifica si un valor es un numero entero
	} else {
		return esCadena(resultado.Value) //Funcion que verifica si un valor es una letra
	}
}

func verificarLongitud(resultado Result) bool {
	//Compruebo si la longitud del valor del resultado que recibo por parametro es igual a la longitud del resultado
	return len(resultado.Value) == resultado.Length
}

func modificarValores(cadena string) Result {
	//El tipo son los primeros dos valores de la cadena, desde la posicion 0 hasta la 2
	tipo := cadena[0:2]

	//La longitud se encuentra desde la posicion 2 hasta la 4
	longitud := cadena[2:4]

	//El valor se encuentra desde la posicion 4 hasta el final, teniendo tantos valores como la longitud lo indica
	valor := cadena[4:]

	//Convierto a entero la longitud, utilizando la funcion ParseInt
	longitudEntero, _ := strconv.ParseInt(longitud, 0, 8)

	//Retorno el nuevo resultado, con sus valores modificados
	return Result{tipo, int(longitudEntero), valor}
}

//Compruebo si el tipo que recibo por parametro es numerico
func esEntero(tipo string) bool {
	//Recorro el tipo que recibo por parametro, y si encuentra una letra retorna falso. Si encuentra un numero retorna verdadero
	for _, c := range tipo {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func esCadena(tipo string) bool {
	//Recorro el tipo que recibo por parametro, y si encuentra un numero retorna falso. Si encuentra una letra retorna verdadero
	for _, r := range tipo {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

//Si inicia con mayuscula es publico - Minuscula privado
