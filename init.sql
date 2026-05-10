CREATE TABLE IF NOT EXISTS clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    total_gasto NUMERIC(10, 2) DEFAULT 0.00,
    idade INT NOT NULL
);

-- Um dado de teste para não começarmos vazios
INSERT INTO clientes (nome, total_gasto, idade) VALUES ('Zezinho', 25.06, 30);