package domain

import "errors"

type Cliente struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	TotalGasto float64 `json:"total_gasto"`
	Idade      int     `json:"idade"`
}

// Validate verifica se os dados do cliente são consistentes
func (c *Cliente) Validate() error {
	if c.Nome == "" {
		return errors.New("O nome é obrigatório")
	}
	if c.Idade < 0 {
		return errors.New("A idade não pode ser negativa")
	}
	if c.Idade > 130 {
		return errors.New("Idade inválida")
	}
	return nil
}

type ClienteRepository interface {
	Create(c *Cliente) error
	GetByID(id int) (*Cliente, error)
	GetAll() ([]Cliente, error)
	Update(c *Cliente) error
	Delete(id int) error
}
