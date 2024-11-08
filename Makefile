LOCAL_BIN:=$(CURDIR)/bin


create-migration:
		migrate create -ext sql -dir migrate -seq $(NAME)

migrate-up:
		go run main.go migrate up

migrate-down:
		go run main.go migrate down

up:
	docker compose up -d

lint:
	./bin/golangci-lint run

get-deps:
		GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
