<!-- markdownlint-disable MD007 MD031 MD034 -->
# Go Expert Labs - Cloud Run challenge

Aplicação web em Go que receba um CEP, identifica a cidade e retorna o clima atual em Celsius, Fahrenheit e Kelvin.

## Executando em PROD (Cloud Run)

A aplicação está disponível para acesso no serviço Google Cloud Run e pode ser acessado seguindo os seguintes parâmetros:

- **Endpoint:** https://goexpert-lab-cloudrun-challenge-mj2kexbmja-uc.a.run.app
- **Método:** GET
- **Query Params:**
  - **zipcode:** CEP a ser consultado

Exemplo de requisição com `curl`:

```bash
curl -X GET https://goexpert-lab-cloudrun-challenge-mj2kexbmja-uc.a.run.app/\?zipcode\=01001000
```

## Executando localmente (dev)

### Requisitos

1. Clone o repositório;
2. Execute o comando `cp .env.example .env` para criar o arquivo de variáveis de ambiente;
3. Edite o novo arquivo `.env` e insira sua chave de acesso à API do [WeatherAPI](https://www.weatherapi.com/) à variável `WEATHER_API_KEY`;

### Via Docker

1. Execute o comando `docker compose up api_dev` para realizar o build do container e iniciar a aplicação na porta declarada no arquivo `.env`;

### Via Makefile

1. Execute o comando `make setup` para instalar as dependências do projeto;
2. Execute o comando `make run` para executar o projeto, que subirá um servidor HTTP na porta 8000;

### Testes

Para executar os testes automatizados, execute o comando `make test`.

## Documentação do endpoint

### Request

| Endpoint | Descrição                                 | Método |  Parâmetro |
|----------|-------------------------------------------|--------|------------|
| /        | Calcula a temperatura atual em uma cidade | GET    | zipcode    |

### Response

- Sucesso:
  - **Código:** 200
  - **Body:**
    ```json
    {
      "temp_C": 23.0,
      "temp_F": 73.4,
      "temp_K": 296.15
    }
    ```

- CEP não encontrado:
    - **Código:** 404
    - **Body:**
      ```json
      {
        "message": "zipcode not found"
      }
      ```

- CEP inválido:
    - **Código:** 422
    - **Body:**
      ```json
      {
        "message": "invalid zipcode"
      }
      ```
