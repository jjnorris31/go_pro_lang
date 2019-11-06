package main

import "fmt"

func main() {
	exampleArray := []int{1, 2, 2, 3, 3, 3, 4, 5}
	exampleArray = eraseDuplicates(exampleArray)
	fmt.Printf("%v", exampleArray)
}

// algoritmo para eliminar las repeticiones adyacentes in-place
func eraseDuplicates(ints []int)[]int {

	i := 0
	lon := len(ints)
	
	for j := 1; j < lon; j++ {
		// si 2 elementos adyacentes son iguales entonces
		if ints[i] == ints[j] {
			tempi := i + 1
			tempj := j + 1
			// copia todo el array después de la repetición
			for tempj < lon {
				ints[tempi] = ints[tempj]
				tempi++
				tempj++
			}			
		} else {
			// si no son iguales, simplemente incrementa ambos índices
			i++
		}
	}
	// regresa un slice con los valores duplicados eliminados
	return ints[:i + 1]
 
}