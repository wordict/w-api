# WORDICT
Wordict is a dictionary app that allow user to store their words with meanings dynamically.

```
📦 wordict
 ┣ 📂 .dev
 ┃ ┣ local.yaml
 ┃ ┗ prod.yaml
 ┣ 📂 cmd
 ┃ ┗ 📂 wordict
 ┃ ┃ ┗ main.go
 ┣ 📂 internal
 ┃ ┣ 📂 config
 ┃ ┃ ┣ config.go
 ┃ ┃ ┗ db.go
 ┃ ┣ 📂 core
 ┃ ┃ ┣ 📂 logger
 ┃ ┃ ┃ ┣ logger.go
 ┃ ┃ ┃ ┗ zap_logger_adapter.go
 ┃ ┃ ┣ 📂 postgres
 ┃ ┃ ┃ ┣ mock_postgres.go
 ┃ ┃ ┃ ┣ postgres.go
 ┃ ┃ ┃ ┗ postgres_test.go
 ┃ ┃ ┗ 📂 server
 ┃ ┃ ┃ ┗ server.go
 ┃ ┣ 📂 entity
 ┃ ┃ ┗ user.go
 ┃ ┣ 📂 handler
 ┃ ┃ ┣ handler.go
 ┃ ┃ ┣ handler_test.go
 ┃ ┃ ┗ mock_handler.go
 ┃ ┣ 📂 repository
 ┃ ┃ ┣ repository.go
 ┃ ┃ ┗ repository_test.go
 ┃ ┗ 📂 service
 ┃ ┃ ┣ service.go
 ┃ ┃ ┗ service_test.go
 ┣ 📂 seed
 ┃ ┣ 📂 cmd
 ┃ ┃ ┗ main.go
 ┃ ┣ create_seed.sql
 ┃ ┣ drop_seed.sql
 ┃ ┣ readme.md
 ┃ ┗ seed.go
 ┣ .gitignore
 ┣ Makefile
 ┣ README.md
 ┣ docker-compose.yml
 ┣ go.mod
 ┗ go.sum
 ```