package main

import "fmt"

func main() {
	i := 1

	//Go don't have aritmetic pointer
	//Ponteiro = referencia de memoria
	var p *int = nil

	p = &i //Pegando end da variavel e jogando no ponteiro
	*p++   //Desreferenciando
	i++
	fmt.Println(p, *p, i)
}
