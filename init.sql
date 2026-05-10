-- Tabela de Produtos
CREATE TABLE IF NOT EXISTS produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    preco NUMERIC(10, 2) NOT NULL,
    estoque INT NOT NULL DEFAULT 0
);

-- Tabela de Pedidos (Relaciona Cliente e Produto)
CREATE TABLE IF NOT EXISTS pedidos (
    id SERIAL PRIMARY KEY,
    cliente_id INT REFERENCES clientes(id),
    produto_id INT REFERENCES produtos(id),
    quantidade INT NOT NULL,
    total NUMERIC(10, 2) NOT NULL,
    data_pedido TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Dados iniciais para teste
INSERT INTO produtos (nome, preco, estoque) VALUES ('Café Espresso', 5.50, 20);
INSERT INTO produtos (nome, preco, estoque) VALUES ('Pão de Queijo', 4.50, 15);