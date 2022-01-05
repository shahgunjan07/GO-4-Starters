package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs","daffy","taz"}

	// Print type and lenght of slice
	fmt.Printf("loons = %v and type of loons = %T \n",loons,loons)
	fmt.Println("-----")

	// Print length of slice
	fmt.Println(len(loons))
	fmt.Println("-----")

	// Print slices
	fmt.Println(loons[0:])
	fmt.Println("-----")

	// Slice value range : Print indices of a slice
	for i := range loons {
		fmt.Println(i)
	}

	// Print value and range indices
	fmt.Println("-----")
	for i, name := range loons {
		fmt.Printf("loons = %s at indice = %d \n",name,i)
	}

	fmt.Println("-----")

	// Print value and ignore range indice
	for _, name := range loons {
		fmt.Printf("loon val = %s \n",name)
	}
}