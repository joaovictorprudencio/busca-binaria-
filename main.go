package main 


import "fmt"

func main()  {

	
	
    lista := []int{1,2,3,4,5,6,7,8,9,10,11}

	alto := len(lista) -1
	chute:= 10
	baixo := 0

	for baixo <= alto {
        meio:= (baixo + alto) / 2

	    if lista[meio] < chute {
			baixo = meio + 1
		} else if lista[meio] > chute  {
			alto = meio -1 
		} else {
			fmt.Printf("valor %d encontrado",chute)
			return

		}
	
	}
 
}



