# run test coverage
test:
	go test -v -cover ./...
	
# run linter run it to avoid waste actions minuts
lint:
	golangci-lint run

# start the application dev
start-dev:
	docker-compose up

# develop start
start:
	go run cmd/main.go

init_db:
	docker run --name leal -e POSTGRES_DB=db-pos -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=root -p 5434:5432 -d postgres