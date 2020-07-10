package main

import "math"

// Forma interface de uma forma retangular ou circulo
type Forma interface {
	Area() float64
}

// Retangulo estrutura para um retangulo
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Area calcula a área de um retângulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Circulo estrutura para um circulo
type Circulo struct {
	Raio float64
}

// Area calcula a área de um circulo
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

// Triangulo estutura de um triangulo
type Triangulo struct {
	Base   float64
	Altura float64
}

// Area calcula a área de um triangulo
func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

// Perimetro recebe largura, altura e retorna o perímetro
func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}
