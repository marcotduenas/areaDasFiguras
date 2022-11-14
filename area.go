package area

import "math"

// Pi é a relação entre circunferência e seu diâmetro
const Pi = 3.1416

//Circulo retorna a área da circunferência
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Retângulo retorna a área do retângulo
func Retangulo(base, altura float64) float64 {
	return base * altura
}

//Privada
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
