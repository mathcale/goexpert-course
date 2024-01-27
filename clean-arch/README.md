# Desafio 3 - Clean Arch

## Como executar

1. Inicie os containers com o comando `docker-compose up -d`;
2. Crie um arquivo `.env` na raiz do projeto com base no `.env.example` com o comando `cp .env.example .env`;
3. Instale as dependências com o comando `make setup`;
4. Execute as migrations com o comando `migrate -path ./migrations -database "mysql://root:root@tcp(localhost:3306)/orders?query" -verbose up`;
5. Execute o projeto com o comando `make run`;

## Mapeamento de serviços

|  **Serviço**   | **Porta** |
| -----------    | --------- |
| Banco de dados | 3306      |
| API            | 8000      |
| GraphQL        | 8080      |
| gRPC           | 50051     |
