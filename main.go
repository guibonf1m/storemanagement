package main

import (
	"fmt"
	"go/types"
	"os"
	"storemanagement/entity"
	"storemanagement/utils"
)

func main() {

	for {
		fmt.Println("1 - Adicionar novo produto")
		fmt.Println("2 - Remover produto por ID")
		fmt.Println("3 - Buscar produto por nome")
		fmt.Println("4 - Listar todos os produtos")
		fmt.Println("5 - Sair \n")
		fmt.Print("Escolha uma opção: ")

		var opcao string
		fmt.Scanln(&opcao)

		switch opcao {

		case "1":

			var Nome string
			var Quantidade int
			var Preco float64

			id := utils.GerarIdAleatorio()
			fmt.Println("O seu ID aleatário: ", id)

			fmt.Println("Qual o nome do produto para cadastro: ")
			fmt.Scanln(&Nome)

			fmt.Println("Qual a quantidade: ")
			fmt.Scanln(&Quantidade)

			fmt.Println("Qual o preço do produto: ")
			fmt.Scanln(&Preco)

			novoProduto := entity.Produto{
				ID:         id,
				Nome:       Nome,
				Quantidade: Quantidade,
				Preco:      Preco,
			}

		case "2":

		case "3":

		case "4":

		case "5":
			fmt.Println("Saindo do programa...")
			os.Exit(5)

		default:
			fmt.Println("Opção inválida, tente novamente!")

		}

	}
}
