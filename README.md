# Cafeteria API ☕

Uma API REST robusta e escalável desenvolvida em **Go**, projetada para o gerenciamento eficiente de uma cafeteria. Este projeto aplica os princípios de **Clean Architecture**, garantindo um código desacoplado, fácil de testar e manter.

---

## 🚀 Diferenciais do Projeto

* **Arquitetura Limpa**: Separação clara entre Domínio, Casos de Uso (Service) e Infraestrutura.
* **Go 1.22+ Native Routing**: Utilização do novo `ServeMux` com suporte a verbos HTTP e wildcards diretamente no roteador padrão.
* **Resiliência**: Validações de regras de negócio integradas ao domínio.
* **Observabilidade**: Middleware de Logging customizado para monitoramento de latência e tráfego.
* **Containerização**: Ambiente totalmente configurado com Docker e Docker Compose.

## 🛠️ Tecnologias Utilizadas

* **Linguagem**: [Go 1.22+](https://go.dev/)
* **Banco de Dados**: [PostgreSQL](https://www.postgresql.org/)
* **Orquestração**: [Docker](https://www.docker.com/) & Docker Compose
* **Variáveis de Ambiente**: `godotenv`
* **Driver DB**: `lib/pq`

## 🏗️ Estrutura de Pastas

```text
.
├── cmd/server/          # Ponto de entrada (Main)
├── internal/
│   ├── database/       # Driver e Singleton de conexão
│   ├── domain/         # Entidades e Interfaces (Core)
│   ├── handler/        # Camada de Transporte (HTTP)
│   ├── middleware/     # Filtros e Logs
│   ├── repository/     # Implementação SQL (Data Access)
│   └── service/        # Lógica de Negócio (Usecases)
├── sql/                # Scripts de Migração (init.sql)
└── docker-compose.yml  # Infraestrutura local
```

## ⚙️ Como Executar

### 1. Clonar e Configurar
```bash
git clone https://github.com/JplsMura/cafeteria-api.git
cd cafeteria-api
cp .env.example .env
```
*Certifique-se de preencher os dados no `.env` conforme o seu ambiente.*

### 2. Subir Infraestrutura (Docker)
```bash
docker-compose up -d
```

### 3. Rodar a API
```bash
go run cmd/server/main.go
```
A API estará disponível em: `http://localhost:8080`

## 📌 Endpoints (Clientes)

| Método | Endpoint | Descrição | Status Code |
| :--- | :--- | :--- | :--- |
| `GET` | `/clientes` | Lista todos os clientes | 200 |
| `POST` | `/clientes` | Cadastra novo cliente | 201 |
| `PUT` | `/clientes/{id}` | Atualiza dados cadastrais | 200 |
| `DELETE` | `/clientes/{id}` | Remove cliente do sistema | 204 |

## 🛡️ Regras de Domínio
- **Validação Automática**: Bloqueio de nomes vazios e idades fora do range (0-130 anos).
- **Segurança de Dados**: Uso de Prepared Statements para evitar SQL Injection.
- **Middleware de Log**: Registro automático de Método, Path, Origem e Tempo de Resposta no console.

## 🗺️ Roadmap de Evolução
- [ ] Testes Unitários (Table-driven tests)
- [ ] Encerramento elegante (Graceful Shutdown)
- [ ] Módulo de Produtos e Pedidos
- [ ] Documentação com Swagger/OpenAPI

<<<<<<< HEAD
---
Desenvolvido por [João Pedro Lima Santos](https://github.com/JplsMura)
=======
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
>>>>>>> 53d522a962c4ab649c0de54a34445d3573981082
