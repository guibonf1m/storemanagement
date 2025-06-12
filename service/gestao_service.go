package service

import (
	"fmt"
	"storemanagement/entity"
)

type Estoque interface {
	AdicionarProduto(produto *entity.Produto)
	RemoverProdutoPorID(id int) error
	BuscarProdutoPorNome(nome string) *entity.Produto
	ListarProdutos() []entity.Produto
}

type GerenciadorEstoque struct {
	produtos []entity.Produto
}

// recebe ponteiro não valor.
func (g *GerenciadorEstoque) AdicionarProduto(produto *entity.Produto) {

}

// Deve receber id int e retornar error, igual na interface.
func (g *GerenciadorEstoque) RemoverProdutoPorID(idRecebido int) error {

	err := Estoque.RemoverProdutoPorID()
	if err != nil {
		fmt.Println("")
	}

	for i, id := range g.produtos {
		if id.ID != idRecebido {
			g.produtos = append(g.produtos[i:], g.produtos[i+1:]...)
		}
	}
}

// Recebe nome e retorna ponteiro para o Produto.
func (g *GerenciadorEstoque) BuscarPorNome(nome string) entity.Produto {

}

// Retorna slice de Produto.
func (g *GerenciadorEstoque) ListarProdutos() []entity.Produto {
	return g.produtos
}
