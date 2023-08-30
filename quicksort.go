/*
	Q U I C K SORT

	A ideia básica deste algoritmo é:
	1) eleger um elemento da lista como pivô e removê-lo da lista;
	2) particionar a lista em duas listas distintas: uma contendo elementos menores que o pivô, e outra contendo elementos maiores;
	3) ordenar as duas listas recursivamente;
	4) retornar a combinação da lista ordenada de elementos menores, o próprio pivô, e a lista ordenada de elementos maiores.

*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido!\n", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}

	fmt.Println(quicksort(numeros))

}

func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indice := len(n) / 2
	pivo := n[indice]

	n = append(n[:indice], n[indice+1:]...)

	menores, maiores := particionar(n, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)

}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}
