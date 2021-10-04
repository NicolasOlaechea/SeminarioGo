# SeminarioGo
Entrega del trabajo practico realizado para el seminario de go

Explicacion:
Creo el archivo "entities.go" donde creo la estructura "Result" y ademas en este archivo tengo las siguientes funciones:
-newResult() que retorna una instancia de Result y le pone los valores que recibe por parametro
-verificarTipo() que retorna si el valor que recibe es un numero o una letra
-verificarLongitud() que comprueba la longitud del resultado que recibe
-modificarValores() que setea los valores del Result y lo retorna
-esEntero() y esCadena() que verifican de que tipo es el valor que reciben

Para el testeo creo el archivo "entities_test.go" donde tengo la siguiente funcion:
TestNewResult() que dentro tiene el dataset que nos dieron en la consigna, pero orientado a mi resolucion y testeo distintos casos.

En el archivo "main.go", dentro de la funcion main(), creo una cadena y un Result. Luego muestro por consola ese resultado.
