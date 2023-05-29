package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	Jogardnv := true
	Matriz := [][]int{}

	for Jogardnv {
		aleatorio := rand.Intn(100) + 1
		fmt.Println("Tente adivinhar um número até 100\n")

		var tentativas []int
		acertou := false

		for !acertou {
			var chutes int
			fmt.Print("Digite um número: ")
			fmt.Scanln(&chutes)
			tentativas = append(tentativas, chutes)

			if chutes == aleatorio {
				acertou = true
				fmt.Printf("\nParabéns, você acertou o número depois de %d tentativas!\n\n", len(tentativas))
			} else if chutes > aleatorio {
				fmt.Println("O número é menor...")
			} else {
				fmt.Println("O número é maior...")
			}
		}

		Matriz = append(Matriz, tentativas)

		var JogardnvInput string
		fmt.Print("Deseja jogar novamente? (s/n): ")
		fmt.Scanln(&JogardnvInput)

		if JogardnvInput != "s" {
			Jogardnv = false
		}

		fmt.Println()
	}

	fmt.Println("Número de tentativas em todas as jogadas:")
	for i, tentativas := range Matriz {
		fmt.Printf("Jogada %d: %v\n", i+1, tentativas)
	}
}
