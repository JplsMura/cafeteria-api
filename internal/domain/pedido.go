package domain

import (
	"context"
	"time"
)

type Pedido struct {
	ID         int       `json:"id"`
	ClienteID  int       `json:"cliente_id"`
	ProdutoID  int       `json:"produto_id"`
	Quantidade int       `json:"quantidade"`
	Total      float64   `json:"total"`
	Data       time.Time `json:"data"`
}

// PedidoRepository define o contrato para persistência de pedidos
type PedidoRepository interface {
	Create(ctx context.Context, p *Pedido) error
}
