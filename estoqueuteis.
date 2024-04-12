
// funcao para adicionar

func (p Produtos) AddProdutos() {
	Produto = append(Produto, p)
}

// para listar produtos
func ListarProdutos() {
	for _, produto := range Produto {
		fmt.Println(produto.ID, produto.Nome, produto.Quantidade, produto.Preco)
	}
}

func AtualizarQuantidaProduto(id int, quantidade int) {
	for i, produto := range Produto {
		if produto.ID == id {
			Produto[i].Quantidade = quantidade
			return
		}
	}
}

