# Movie Watchlist — Back-end (Go)

API REST desenvolvida em Go com Gin e GORM para gerenciar uma lista pessoal de filmes.

## Tecnologias

- **Go 1.22** + **Gin** (roteamento)
- **GORM** + **PostgreSQL** (banco de dados)
- **Docker** com multi-stage build

## Variáveis de Ambiente

| Variável | Descrição | Exemplo |
|---|---|---|
| `DATABASE_URL` | DSN do PostgreSQL | `postgres://user:pass@db:5432/watchlist` |
| `OMDB_API_KEY` | Chave da OMDb API | `abc123` |
| `PORT` | Porta do servidor (padrão: 8080) | `8080` |

## OMDb API

A API externa utilizada é a **OMDb API** (Open Movie Database).

- Site: https://www.omdbapi.com
- **Licença:** dados fornecidos sob licença [Creative Commons Attribution-NonCommercial 4.0](https://creativecommons.org/licenses/by-nc/4.0/) — uso não comercial permitido com atribuição
- **Cadastro:** obrigatório para obter chave de API (plano FREE: 1.000 req/dia, sem custo)
- Rota consumida pelo back-end: `http://www.omdbapi.com/?t={titulo}&apikey={chave}`

O back-end atua como intermediário para não expor a chave da API no front-end.

## Rotas

| Método | Rota | Descrição |
|---|---|---|
| GET | `/api/search?title=` | Busca filme na OMDb |
| GET | `/api/movies` | Lista watchlist (aceita `?watched=true/false`) |
| POST | `/api/movies` | Adiciona filme à watchlist |
| PUT | `/api/movies/:id` | Atualiza `is_watched` e/ou `personal_rating` |
| DELETE | `/api/movies/:id` | Remove filme da watchlist |

## Executar com Docker Compose

O `docker-compose.yml` deste repositório sobe o stack completo (banco, backend e frontend). Os dois repositórios devem estar na mesma pasta pai.

```bash
# Na pasta mvp-ARQSW-backend
docker compose up --build

# Rodar em segundo plano
docker compose up --build -d

# Parar os serviços
docker compose down

# Parar e remover o volume do banco de dados
docker compose down -v
```

> O mesmo resultado é obtido rodando `docker compose up --build` na pasta `mvp-ARQSW-frontend`.

## Containers individuais (sem Docker Compose)

```bash
# 1. Criar rede compartilhada
docker network create watchlist-net

# 2. Banco de dados
docker run -d \
  --name movie_db \
  --network watchlist-net \
  -e POSTGRES_USER=watchlist \
  -e POSTGRES_PASSWORD=watchlist123 \
  -e POSTGRES_DB=watchlist \
  postgres:16-alpine

# 3. Back-end (na pasta mvp-ARQSW-backend)
docker build -t movie-watchlist-api .
docker run -d \
  --name movie_backend \
  --network watchlist-net \
  -e DATABASE_URL="postgres://watchlist:watchlist123@movie_db:5432/watchlist" \
  -e OMDB_API_KEY="sua_chave_aqui" \
  -e PORT=8080 \
  movie-watchlist-api

# 4. Front-end (na pasta mvp-ARQSW-frontend)
docker build -t movie-watchlist-frontend .
docker run -d \
  --name movie_frontend \
  --network watchlist-net \
  -p 3000:80 \
  movie-watchlist-frontend
```

Após subir, acesse: **http://localhost:3000**

Para parar e remover:

```bash
docker rm -f movie_frontend movie_backend movie_db
docker network rm watchlist-net
```
