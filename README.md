## Roda aplicação dev
```bash
go run cmd/app/main.go
```

## 📁 Estrutura do Repositório
  - cmd/app # Ponto de entrada do backend
  - internal/database # Queries geradas via sqlc
  - internal/domain # Entidades (DDD)
  - internal/handler # Handlers e controladores HTTP
  - internal/service # Logica e manipulação de dados
  - test/ # Testes unitários e de integração
  - schema.sql # Esquema SQL das tabelas
  - queries.sql # Comandos SQL
