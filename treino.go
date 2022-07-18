package main

import "fmt"

type Operacoes struct {
	operacao string
}

func main() {
	calc := Operacoes{
		operacao: "soma",
	}

	fmt.Println(calc.calculaTudo(5, 5, 5, 5, 10000))
	
}

func (op Operacoes) calculaTudo(x ...int) int {
	result := 0
	for _, v := range x {
		if op.operacao == "soma" {
			result += v
		} else if op.operacao == "sub" {
			result -= v
		} else if op.operacao == "mult" {
			result *= v
		}
	}

	return result
}