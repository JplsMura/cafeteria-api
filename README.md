⚙️ Configuração e Execução

Pré-requisitos
Go 1.22 ou superior instalado.

Docker e Docker Compose instalados.

Passos para Rodar
1. Clone o repositório:

    {git clone [https://github.com/seu-usuario/cafeteria-api.git](https://github.com/seu-usuario/cafeteria-api.git)
cd cafeteria-api}

2. Configure o ambiente:
    Crie um arquivo .env na raiz baseado no .env.example:

    {cp .env.example .env}

3. Suba o banco de dados:
    {docker-compose up -d}

4. Inicie o servidor:
    {go run cmd/server/main.go}

O servidor estará disponível em http://localhost:8080.

📌 Endpoints da API

Clientes

Método|Endpoint|Descrição

GET     | /clientes        | Lista todos os clientes cadastrados
POST    | /clientes        | Cadastra um novo cliente (Valida nome/idade)
PUT     | /clientes/{id}   | Atualiza dados de um cliente existente
DELETE  | /clientes/{id}   | Remove um cliente do sistema

🛡️ Validações e Segurança
Domínio Protegido: O sistema impede o cadastro de clientes com nomes vazios ou idades inconsistentes (negativas ou acima de 130 anos).

Logging Middleware: Todas as operações são logadas no terminal, informando o Método, Rota, IP e o Tempo de Execução.

Variáveis de Ambiente: Credenciais de banco de dados não são expostas no código, sendo carregadas via .env.

🗺️ Roadmap de Evolução
[ ] Implementação de Testes Unitários com testing do Go.

[ ] Implementação de Graceful Shutdown.

[ ] CRUD de Produtos e Pedidos.

[ ] Documentação interativa com Swagger/OpenAPI.

Gerando o .env.example para completar o pacote
env_example = 
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASS=password
DB_NAME=cafeteira
