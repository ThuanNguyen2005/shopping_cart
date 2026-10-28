include .env
export

MIGRATION_DIR=internal/db/Migrations

CONN_STRING=postgresql://$(USERNAME_DB):$(PASSWORD_DB)@$(HOST):$(PORT)/$(NAME_DB)?sslmode=$(SSLMODE)

export_db:
	docker exec -i postgres-db pg_dump -U root -d master-golang >.backupdb-master-golang.sql

import_db:
	docker exec -i postgres-db psql -U root -d master-golang < .backupdb-master-golang.sql

server: 
	go run d:\go\30.project_shopping_cart\cmd\api\main.go

sqlc:
	sqlc generate 

migrate_create:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(NAME)

migrate_up:
	migrate -path $(MIGRATION_DIR) -database "$(CONN_STRING)" up

migrate_down:
	migrate -path $(MIGRATION_DIR) -database "$(CONN_STRING)" down 1

migrate_force:
	migrate -path $(MIGRATION_DIR) -database $(CONN_STRING) force $(VERSION)

migrate_drop:
	migrate -path $(MIGRATION_DIR) -database $(CONN_STRING) drop

migrate-goto:
	migrate -path $(MIGRATION_DIR) -database $(CONN_STRING) goto $(VERSION)

migrate-down-n:
	migrate -path $(MIGRATION_DIR) -database $(CONN_STRING) down $(N)

.PHONY: export_db export_db migrate_create server
