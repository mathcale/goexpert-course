<!-- markdownlint-disable MD007 MD031 MD034 -->
# Go Expert Labs - OpenTelemetry challenge

Instrumentação de serviços com OpenTelemetry e Zipkin

## Executando localmente (dev)

### Requisitos

1. Clone o repositório;
2. Execute o comando `make env` para criar os arquivos de variáveis de ambiente;
3. Edite os novos arquivo `.env` e `.env.docker` e insira sua chave de acesso à API do [WeatherAPI](https://www.weatherapi.com/) à variável `WEATHER_API_KEY`;

### Via Docker

1. Execute o comando `docker compose up` para realizar o build dos containers e iniciar as aplicações nas portas declaradas no arquivo `.env.docker` (8000 e 8001, por padrão);

### Via Go

1. Execute `docker compose up -d collector zipkin` para iniciar o serviço de coleta de traces e o Zipkin;
2. Execute o comando `make run-input` e `make run-orchestrator` para iniciar a aplicação em modo de desenvolvimento;

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
