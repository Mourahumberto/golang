package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "sexta"
	case 6:
		return "sabado"
	case 7:
		return "sabado"
	default:
		return "Número inválido"
	}
	// switch {
	// case numero == 1:
	// 	return "Domingo"
	// case numero == 2:
	// 	return "Segunda"
	// case numero == 3:
	// 	return "terça"
	// case numero == 4:
	// 	return "quarta"
	// case numero == 5:
	// 	return "sexta"
	// case numero == 6:
	// 	return "sabado"
	// case numero == 7:
	// 	return "sabado"
	// case numero > 7:
	// 	return "ta doido?"
	// default:
	// 	return "Número inválido"
	// }
}

func main() {
	dia := diaDaSemana(3)
	fmt.Println(dia)
	fmt.Println(diaDaSemana(3))
}
