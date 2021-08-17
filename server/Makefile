sqlc:
	sqlc generate

migrate-create:
	migrate create -dir=sql/migration -ext=sql -seq $(NAME)

.PHONY: sqlc migrate-create
