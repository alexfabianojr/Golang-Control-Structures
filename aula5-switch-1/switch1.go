package main

import "fmt"

func notaParaConceito(n float64) string {
	nota := int(n)
	//no go o break; eh implicito, se quisermos causar o efeito do switch de java sem break usamos a palavra reservada fallthrought
	switch nota {
	case 10:
		return "A"
	case 9, 8:
		return "B"
	case 7, 6:
		return "C"
	case 5, 4, 3, 2, 1, 0:
		return "D"
	default:
		return "Erro"
	}
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(2))
	fmt.Println(notaParaConceito(0))
	fmt.Println(notaParaConceito(7))
	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(6))
	fmt.Println(notaParaConceito(-1))
}
