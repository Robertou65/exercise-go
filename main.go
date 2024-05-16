package main

import (
	"fmt"
	"strings"
)

func main() {
	// EJERCICIO: encontrar cuantas veces se repite una palabra en una cadena

	word := "Hola, mi nombre es robert. Mi nombre completo es Robert Hernandez (Robertou65)."

	word = strings.ToLower(word) // pasando toda la cadena a minusculas

	symbols := "!#$%&'()*+,-./_" // simbolos que se van a eliminar

	// bucle para recorer toda la cadena
	for i := 0; i < len(word); i++ {
		if strings.Contains(symbols, string(word[i])) { // condicional que analiza si la cadena contiene alguno de los simbolos
			aux1 := word[:i]   // toma desde el principio hasta una posicion anterior a la actual
			aux2 := word[i+1:] // toma desde una posicion posterior hasta el final
			word = aux1 + aux2 // se unen a la variable original

			if i == len(word)-1 { // si hay un simbolo al final este condicional se ejecutara
				aux1 = word[:i] // toma desde el principio hasta una posicion anterior a la actual
				word = aux1
			}
		}
	}

	// separa la cadena por los espacios y lo guarda en una arreglo
	var arrayWords []string = strings.Split(word, " ")
	mapWords := make(map[string]int) // mapa que guardara las palabras y la veces que se repitieron

	for i := 0; i < len(arrayWords); i++ {
		_, is := mapWords[arrayWords[i]] // guarda true si la palabra esta en el mapa
		if is {                          // si esta entonses solo suma 1
			mapWords[arrayWords[i]] += 1
		} else { // caso contrario la crea y asigna 1
			mapWords[arrayWords[i]] = 1
		}
	}

	// imprime la palabra y las veces que aparece
	for i, k := range mapWords {
		// i para la key y k para el valor
		fmt.Println(i, "aparece", k, "veces")
	}
}
