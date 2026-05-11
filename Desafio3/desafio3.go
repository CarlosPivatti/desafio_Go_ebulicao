package main

import (
	"fmt"
	"time"
)

func pingar(canal chan string) { // palavra reservada "chan" indica que a função recebe um canal como argumento
	for i := 0; ; i++ {
		canal <- "ping" // operador de canal "<-" para enviar o valor "ping" para o canal
	}
}
func pongar(canal chan string) {
	for i := 0; ; i++ {
		canal <- "pong" // operador de canal "<-" para enviar o valor "pong" para o canal
	}
}

func imprimir(canal chan string) {
	for {
		msg := <-canal // operador de canal "<-" para receber um valor do canal e armazená-lo na variável "msg"
		fmt.Println(msg)
		time.Sleep(time.Second * 5) // pausa de 1 segundo para evitar que a saída seja muito rápida
	}
}
func main() {
	canal := make(chan string) // função "make" para criar um canal do tipo string
	go pingar(canal)
	go pongar(canal)
	go imprimir(canal)
	var input string
	fmt.Scanln(&input) // espera por uma entrada do usuário para evitar que o programa termine imediatamente
}
