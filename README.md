## Pré-requisitos

- **Docker e Docker Compose** instalados
- **Go** (versão 1.19 ou superior) instalado

---

## Configuração

1. **Crie um arquivo `.env`**

   Antes de iniciar a aplicação, crie um arquivo `.env` na pasta `cmd/ordersystem` com as seguintes configurações:

   ```env
   DB_DRIVER=mysql
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=root
   DB_NAME=orders
   WEB_SERVER_PORT=:8000
   GRPC_SERVER_PORT=50051
   GRAPHQL_SERVER_PORT=8080
   ```

   **Nota:** Ajuste as variáveis conforme necessário.

---

## Passos para execução

1. **Inicie os serviços necessários com Docker Compose**

   Na raiz do projeto, execute:

   ```bash
   docker-compose up -d
   ```

   Isso iniciará o banco de dados MySQL e outros serviços definidos no `docker-compose.yml`.

2. **Execute a aplicação**

   Ainda na pasta `cmd/ordersystem`, inicie a aplicação com:

   ```bash
   go run main.go wire_gen.go
   ```

---

## Portas

A aplicação responde nas seguintes portas:

- **Web Server (REST):** `http://localhost:8000`
- **gRPC Server:** `localhost:50051`
- **GraphQL Playground:** `http://localhost:8080`

---

## Endpoints disponíveis

- **REST API:**

  - `POST /order` - Criação de um novo pedido
  - `GET /orders` - Listagem de todos os pedidos

- **GraphQL Playground:** Navegue até `http://localhost:8080` para explorar a API GraphQL.

- **gRPC Server:** Utilize uma ferramenta como o **Evans CLI** ou o **BloomRPC** para interagir com o servidor gRPC.

---

## Observações

- Certifique-se de que o MySQL está rodando antes de iniciar a aplicação.
- O sistema criará automaticamente a tabela `orders` no banco de dados na primeira execução.
