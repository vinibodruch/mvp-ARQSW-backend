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

## Executar com Docker

> Os comandos abaixo devem ser executados a partir do repositório do **front-end**, que contém o `docker-compose.yml`.

```bash
# Na pasta mvp-ARQSW-frontend
docker compose up --build
```

Para rodar apenas o back-end isolado:

```bash
# Na pasta mvp-ARQSW-backend
docker build -t movie-watchlist-api .
docker run -p 8080:8080 \
  -e DATABASE_URL="postgres://user:pass@host:5432/watchlist" \
  -e OMDB_API_KEY="sua_chave_aqui" \
  movie-watchlist-api
```
