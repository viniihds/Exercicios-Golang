package main

import(
	"fmt"
)


func main(){
	res := MetroParaMilimetros(3)

	res2 := TempoParaDistancia(8,33.5)

	res3 := AumentoSalarial(1344, 30)

	res4 := PrecoDesconto(49.90, 20)

	res5 := DiasParaSegundos(3)

	res6 := CelsiusParaFahrenheit(18)

	res7 := AluguelCarro(3,142)

	res8 := AluguelAirbnb(6, 183)

	res9 := DiasPerdidosPorFumar(5, 32)
	
	res10 := ContarAlgarismos(19876)

	fmt.Println(res, "\n", res2, "\n", res3, "\n", res4, "\n", res5, "\n", res6, "\n", res7, "\n", res8, "\n", res9, "\n", res10)
}

func MetroParaMilimetros(a int) int{
	// Recebe um valor em Metros(a) e retorna em milimetros
	return a*1000
}

func TempoParaDistancia(a float64, b float64) float64{
	//Recebe a distancia em km(a) e a velocidade em km/h(b) e retorna o tempo para percorrer essa distancia
	return a/b
}

func AumentoSalarial(a float64, b float64) float64{
	//Recebe o salario(a) e a porcentagem do aumento(b) e retorna o novo salário
	return a+(a*(b/100))
}

func PrecoDesconto(a float64, b float64) float64{
	//Recebe o preço do produto(a) e a porcentagem do desconto(b) e retorna o novo valor após o desconto
	return a-(a*(b/100))
}

func DiasParaSegundos(a int) int{
	//Recebe o número de dias(a) e retorna em segundos
	return a*24*60*60
}

func CelsiusParaFahrenheit(a int) int{
	//Recebe a temperatura em graus Celsius e devolve em Fahrenheit
	return ((a*18)/10)+32
}

func AluguelCarro(a float64, b float64) float64{
	//Recebe uma quantidade de dias que o carro foi alugado e a quantidade de quilômetros rodados, e retorna o valor a ser pago. 1 dia: 60 reais mais R$ 0,15 por km rodado.
	return a*60 + (b*0.15)
}

func AluguelAirbnb(a float64, b float64) float64{
	//Recebe uma quantidade de dias que ficou hospedado e o valor da diária, e retorna o valor a ser pago, considerando um acréscimo de R$ 75,00 para limpeza e 5% de taxa de administração sobre o valor do aluguel, sem a taxa de limpeza
	return 75+a*b+(a*b*(0.05))
}

func DiasPerdidosPorFumar(a int, b int) int{
	//Recebe uma quantidade de cigarros fumados por dia(a) e a quantidade de anos que fuma(b), e retorna o total de dias perdidos, sabendo que cada cigarro reduz a vida em 10 minutos.
	return (a*10*365*b)/60/24
}

func ContarAlgarismos(a int) int{
	count := 0
	for a != 0 {
		a /= 10
		count += 1
	}
	return count
}