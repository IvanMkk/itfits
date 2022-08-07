server-run:
	go run src/server/main.go

server-migration-create:
	migrate create -ext sql -dir db/migrations -seq create_users_table

server-migrate:

# create
# migrate -source src/server/src/schemas/init.go -database postgres://localhost:5432/database up 2
	docker run -v /home/omega/Documents/itfits/src/server/db/migrations:/migrations \
		--network host migrate/migrate \
		-path=/migrations/ \
		-database $POSTGRES \
		up

server-test:
	go test -v ./...
