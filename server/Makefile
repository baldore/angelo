sqlc:
	sqlc generate

migrate-create:
	migrate create -dir=sql/migration -ext=sql -seq $(NAME)

reset-dev-database:
	psql -d angelo -f sql/migration/000001_setup_tables.down.sql
	psql -d angelo -f sql/migration/000001_setup_tables.up.sql
	psql -d angelo -f seed.sql

.PHONY: sqlc migrate-create reset-dev-database
