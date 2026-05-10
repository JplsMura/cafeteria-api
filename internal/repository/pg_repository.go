package repository

import (
	"cafeteria-api/internal/domain"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(c *domain.Cliente) error {
	query := `INSERT INTO clientes (nome, total_gasto, idade) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, c.Nome, c.TotalGasto, c.Idade).Scan(&c.ID)
}

func (r *PostgresRepository) GetAll() ([]domain.Cliente, error) {
	rows, err := r.db.Query("SELECT id, nome, total_gasto, idade FROM clientes")
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

func (r *PostgresRepository) GetByID(id int) (*domain.Cliente, error) {
	return nil, nil // Implementaremos depois
}

func (r *PostgresRepository) Update(c *domain.Cliente) error {
	query := `UPDATE clientes SET nome = $1, total_gasto = $2, idade = $3 WHERE id = $4`
	_, err := r.db.Exec(query, c.Nome, c.TotalGasto, c.Idade, c.ID)
	return err
}

func (r *PostgresRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM clientes WHERE id = $1", id)
	return err
}
