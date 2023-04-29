package main

import (
	"fmt"
	"guia2/queue"
	"guia2/stack"
)

func main() {
	/* fmt.Println(InvertirCadena("Hola mundo"))
	fmt.Println(Palindromo("sosos"))
	fmt.Println(Balanceada("[[{}[()]()]()]")) */
	q1 := queue.Queue{}
	q2 := q1
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)
	q1.Enqueue(4)
	q1.Enqueue(5)

	q2.Enqueue(7)
	q2.Enqueue(8)
	q2.Enqueue(9)

	fmt.Println(UnirColas(q1, q2))
}

/* func InvertirCadena(cadena string) string {
	bytes := []byte(cadena)

	for i := 0; i < len(bytes)/2; i++ {
		j := len(bytes) - 1 - i
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	cadena = string(bytes)
	return cadena
} */

// Recibe una cadena por parámetro y la devuelve invertida
func InvertirCadena(cadena string) string {
	s := stack.Stack{}
	invertido := ""

	for i := 0; i < len(cadena); i++ {
		s.Push(string(cadena[i]))
	}
	for !s.IsEmpty() {
		aux, _ := s.Pop()
		invertido += aux.(string)
	}

	return invertido
}

/* func Palindromo(cadena string) bool {
	bytes := []byte(cadena)

	for i := 0; i < len(bytes)/2; i++ {
		j := len(bytes) - i - 1
		if bytes[i] != bytes[j] {
			return false
		}
	}
	return true
} */

// Recibe una cadena por parámetro y evalua si es un palíndromo
func Palindromo(cadena string) bool {

	return InvertirCadena(cadena) == cadena
}

// Recibe una cadena de caracteres por parámetro que puede variar entre "{},[],()" y evalua si dicha cadena está balanceada
func Balanceada(cadena string) bool {
	s := stack.Stack{}

	for i := 0; i < len(cadena); i++ {
		aux := string(cadena[i])

		if aux == "[" || aux == "(" || aux == "{" {
			s.Push(aux)
		} else {
			aux2, _ := s.Pop()
			switch aux {
			case "}":
				if aux2 != "{" {
					return false
				}
			case "]":
				if aux2 != "[" {
					return false
				}
			case ")":
				if aux2 != "(" {
					return false
				}
			default:
			}
		}
	}
	return true
}

// Recibe 2 colas por parametros q1 y q2 y devuelve una nueva cola resultante de
// concatenar q2 al final de q1
func UnirColas(q1, q2 queue.Queue) (a queue.Queue) {
	a = q1
	//for while
	for !q2.IsEmpty() {
		aux, _ := q2.Dequeue()
		a.Enqueue(aux)
	}
	return
}
