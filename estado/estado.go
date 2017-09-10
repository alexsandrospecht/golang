package main

import "fmt"

func main() {
	estados := make(map[string]Estado, 6)
	estados["GO"] = Estado{"Goias", 6414123, "Goiânia"}
	estados["SE"] = Estado{"Sergipe", 1214123, "Aracaju"}

	sergipe := estados["SE"]

	saoPaulo, found := estados["SP"]
	if (found) {
		fmt.Println(saoPaulo)
	}

	//delete(estados, "SE")

	fmt.Println(sergipe)
	fmt.Println(estados)

	for sigla, estado  := range estados {
		fmt.Printf("%s (%s) possui %d habitantes \n", estado.nome, sigla, estado.populacao)
	}

}

type Estado struct {
	nome string
	populacao int
	capital string
}
