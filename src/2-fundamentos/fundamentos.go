package main

import "fmt"

type Pessoa struct {
	nome      string
	idade     int8
	sobrenome string
}

func (p1 Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s %s e tenho %d anos.", p1.nome, p1.sobrenome, p1.idade)
}

type Stringer interface {
	String() string
}

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("2.5 + 2.5 =", 2.5+2.5)

	fmt.Println(len("String"))
	fmt.Println("Hello, " + "World!")
	fmt.Println("String"[2])

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var nome string = "Batata"
	var idade int = 5
	var altura float32 = 1.75
	var estudante bool = true

	fmt.Println("\n\nOlá sr(a). ", nome, "\nIdade: ", idade, "\nAltura: ", altura, "\nEstudante: ", estudante)

	var nome2 = "Batata"
	var idade2 = 5
	var altura2 = 1.75
	var estudante2 = true

	fmt.Println("\n\nOlá sr(a). ", nome2, "\nIdade: ", idade2, "\nAltura: ", altura2, "\nEstudante: ", estudante2)

	var x [5]int
	x[4] = 13
	fmt.Println(x)

	p1 := Pessoa{"Batata", 5, "Frita"}
	p2 := Pessoa{nome: "Cenoura", idade: 3, sobrenome: "Cozida"}
	p3 := Pessoa{nome: "Alface", sobrenome: "Crua"}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	Print(p1)

}

func Print(p1 Stringer) {
	fmt.Println(p1.String())
}
