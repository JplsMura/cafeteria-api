package domain

import (
	"context"
	"errors"
)

type Produto struct {
	ID      int     `json:"id"`
	Nome    string  `json:"nome"`
	Preco   float64 `json:"preco"`
	Estoque int     `json:"estoque"`
}

func (p *Produto) Validate() error {
	if p.Nome == "" {
		return errors.New("Nome do produto é obrigatório")
	}
	if p.Preco <= 0 {
		return errors.New("Preço do produto deve ser maior que zero")
	}
	return nil
}

// ProdutoRepository define o contrato para persistência de produtos
type ProdutoRepository interface {
	GetByID(ctx context.Context, id int) (*Produto, error)
	UpdateEstoque(ctx context.Context, id int, quantidade int) error
}
