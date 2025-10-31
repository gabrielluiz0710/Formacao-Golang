package main

import "fmt"

const ebulicaoF = 212.0

func main() {
	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9

	fmt.Println("Temperatura de ebulição da agua em ºF = ", tempF)
	fmt.Println("Temperatura de ebulição da agua em ºC = ", tempC)

	fmt.Printf("\n\nTemperatura de ebulição da agua em ºF = %g\nTemperatura de ebulição da agua em ºC = %g", tempF, tempC)

	x := 1
	fmt.Printf("\n\nTipo x: %T", x)

	y := 0.2
	c := float64(x) * y
	fmt.Printf("\n\nTipo c: %T", c)
}
