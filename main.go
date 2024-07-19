package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name   string
	Weight float64
	Height float64
}

func (p *Person) CalculateIMC() float64 {
	return p.Weight / math.Pow(p.Height, 2)
}

func (p *Person) GetIMCCategory() string {
	imc := p.CalculateIMC()
	switch {
	case imc < 18.5:
		return "Bajo peso"
	case imc < 25:
		return "Peso normal"
	case imc < 30:
		return "Sobrepeso"
	default:
		return "Obesidad"
	}
}

func main() {
	var name string
	var weight, height float64

	fmt.Println("¡Bienvenido al Calculador de IMC en Go!")

	fmt.Print("Introduce tu nombre: ")
	fmt.Scanln(&name)

	fmt.Print("Introduce tu peso en kg: ")
	fmt.Scanln(&weight)

	fmt.Print("Introduce tu altura en metros: ")
	fmt.Scanln(&height)

	person := Person{Name: name, Weight: weight, Height: height}

	imc := person.CalculateIMC()
	category := person.GetIMCCategory()

	fmt.Printf("\n¡Hola %s! Aquí están tus resultados:\n", person.Name)
	fmt.Printf("Tu IMC es: %.2f\n", imc)
	fmt.Printf("Categoría: %s\n", category)

	switch category {
	case "Bajo peso":
		fmt.Println("¡Come un poco más! Estás más flaco que un lápiz.")
	case "Peso normal":
		fmt.Println("¡Perfecto! Mantienes un equilibrio digno de un maestro Jedi.")
	case "Sobrepeso":
		fmt.Println("Un poco de ejercicio no te vendría mal. ¡A mover ese cuerpo!")
	case "Obesidad":
		fmt.Println("Es hora de ponerse en marcha. ¡Tú puedes hacerlo!")
	}
}
