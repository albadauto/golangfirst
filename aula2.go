package main

import "fmt"


type Carro struct{
	Name string

}

func (c Carro) andar (){
	fmt.Println(c.Name, "Andou")
}

func main() {

	carro := Carro{
		Name: "BMW",
	}

	carro.andar()

	resultado := func(x ...int) func() int {
		res := 0
		for _, v := range x { // A parte branca no for é pq não irá usar a key do array
			res += v
		}
		return func() int {
			return res * res

		}
	}
	
	fmt.Println(resultado(54, 54 ,54 ,54)())
}

func soma(a int, b int) (result int) {
	result = a + b
	return
}

func somaTudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}

	return resultado
}
