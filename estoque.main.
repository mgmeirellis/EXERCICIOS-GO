package main

import (
	"fmt"
	"main/uteis"
)

func main() {

	for {
		uteis.Menu()
		fmt.Println("Escolha sua opcao")
		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var produto uteis.Produtos
			fmt.Println("Digite o ID do produto")
			fmt.Scan(&produto.ID)
			fmt.Println("Digite o Nome do produto")
			fmt.Scan(&produto.Nome)

			fmt.Println("Digite a quantidade do produto")
			fmt.Scan(&produto.Quantidade)

			fmt.Println("Digite o preço do produto ")
			fmt.Scan(&produto.Preco)

			produto.AddProdutos()
			break

		case 2:
			var id int
			var quantidade int
			fmt.Println("Digite o Id para atualizar quantidade: ")
			fmt.Scan(&id)
			fmt.Println("Nova Quantidade: ")
			fmt.Scan(&quantidade)
			uteis.AtualizarQuantidaProduto(id, quantidade)
			break

		case 3:
			uteis.ListarProdutos()

			break
		}
		if opcao == 4 {
			break
		}

	}
