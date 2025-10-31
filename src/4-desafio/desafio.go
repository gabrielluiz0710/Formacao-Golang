package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	tempK := ebulicaoK
	tempC := tempK - 273

	fmt.Printf("Temperatura de ebuilição em ºK: %g\nTemperatura de ebulição em ºC: %g", tempK, tempC)
}
