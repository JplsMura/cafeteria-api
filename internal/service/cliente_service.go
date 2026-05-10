package service

import (
	"cafeteria-api/internal/domain"
	"context"
)

type ClienteService struct {
	repo domain.ClienteRepository
	notif *NotificacaoService
}

func NewClienteService(repo domain.ClienteRepository, notif *NotificacaoService) *ClienteService {
	return &ClienteService{
		repo: repo,
		notif: notif,
	}
}

func (s *ClienteService) Criar(ctx context.Context, c *domain.Cliente) error {
	if err := c.Validate(); err != nil {
		return err
	}
	
	err := s.repo.Create(ctx, c)
	if err != nil {
		return err
	}

	// Disparamos o processo em background sem travar a resposta da API
	// Usamos context.Background() ou um novo contexto aqui para garantir que 
	// a tarefa continue mesmo se a requisição HTTP original terminar.
	go s.notif.EnviarBoasVindas(c.Nome)

	return nil
}

func (s *ClienteService) Listar(ctx context.Context, limit, offset int) ([]domain.Cliente, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *ClienteService) Deletar(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *ClienteService) Atualizar(ctx context.Context, c *domain.Cliente) error {
	if err := c.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, c)
}
