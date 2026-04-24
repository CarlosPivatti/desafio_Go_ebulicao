package main //Declarar pacote principal//

import "fmt" //Importar pacote fmt para formatação de entrada e saída//

const ebolicaoK = -273.15 //Declarar variavel constante ebolicao Kelvin do tipo float64 e atribuir valor//

func main() { //Função principal//

	tempKelvin := -(ebolicaoK) + 100     //variavel do valor tempKelvin do tipo float64//
	tempCelsius := (tempKelvin) - 273.15 //variavel do valor tempCelsius do tipo float64//

	// Apareça e resultado em tela//
	fmt.Println("Temperatura em Kelvin = ", tempKelvin, "para embulição da água")
	fmt.Println("Temperatura em Celsius= ", tempCelsius, "para embulição da água")
	fmt.Printf("Temperatura em Kelvin: %g e Temperatura em Celsius: %g", tempKelvin, tempCelsius, "ambos para embulição da água")
}
