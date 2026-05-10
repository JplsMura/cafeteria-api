package repository

import (
	"cafeteria-api/internal/domain"
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, c *domain.Cliente) error {
	query := `INSERT INTO clientes (nome, total_gasto, idade) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRowContext(ctx, query, c.Nome, c.TotalGasto, c.Idade).Scan(&c.ID)
}

func (r *PostgresRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.Cliente, error) {
	query := `SELECT id, nome, total_gasto, idade FROM clientes LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []domain.Cliente
	for rows.Next() {
		var c domain.Cliente
		if err := rows.Scan(&c.ID, &c.Nome, &c.TotalGasto, &c.Idade); err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return clientes, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int) (*domain.Cliente, error) {
	query := `SELECT id, nome, total_gasto, idade FROM clientes WHERE id = $1`
	var c domain.Cliente
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Nome, &c.TotalGasto, &c.Idade)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *PostgresRepository) Update(ctx context.Context, c *domain.Cliente) error {
	query := `UPDATE clientes SET nome = $1, total_gasto = $2, idade = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, c.Nome, c.TotalGasto, c.Idade, c.ID)
	return err
}

func (r *PostgresRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM clientes WHERE id = $1", id)
	return err
}