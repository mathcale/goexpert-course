<!-- markdownlint-disable MD007 MD031 MD034 -->
# Go Expert Labs - OpenTelemetry challenge

Instrumentação de serviços com OpenTelemetry e Zipkin

## Executando localmente (dev)

### Requisitos

1. Clone o repositório;
2. Execute o comando `cp .env.example .env` para criar o arquivo de variáveis de ambiente;
3. Edite o novo arquivo `.env` e insira sua chave de acesso à API do [WeatherAPI](https://www.weatherapi.com/) à variável `WEATHER_API_KEY`;

### Via Docker

1. Execute o comando `docker compose up api_input api_orchestrator` para realizar o build do container e iniciar as aplicações nas portas declaradas no arquivo `.env` (8000 e 8001, por padrão);

### Testes

Para executar os testes automatizados, execute o comando `make test`.

## Documentação dos endpoints

### Input API

WIP

### Orchestrator API

#### Request

| Endpoint | Descrição                                 | Método |  Parâmetro |
|----------|-------------------------------------------|--------|------------|
| /        | Calcula a temperatura atual em uma cidade | GET    | zipcode    |

#### Response

- Sucesso:
  - **Código:** 200
  - **Body:**
    ```json
    {
      "city": "São Paulo",
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
