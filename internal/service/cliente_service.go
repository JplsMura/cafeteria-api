package service

import "cafeteria-api/internal/domain"

type ClienteService struct {
	repo domain.ClienteRepository
}

func NewClienteService(repo domain.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) Criar(c *domain.Cliente) error {
	if err := c.Validate(); err != nil {
		return err
	}
	return s.repo.Create(c)
}

func (s *ClienteService) ListarTodos() ([]domain.Cliente, error) {
	return s.repo.GetAll()
}

func (s *ClienteService) Deletar(id int) error {
	return s.repo.Delete(id)
}

func (s *ClienteService) Atualizar(c *domain.Cliente) error {
	if err := c.Validate(); err != nil {
		return err
	}
	return s.repo.Update(c)
}
