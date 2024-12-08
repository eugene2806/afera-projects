LOCAL_BIN:=$(CURDIR)/bin


start-server:
		go run main.go rest

create-migration:
		migrate create -ext sql -dir migrate -seq $(NAME)

migrate-up:
		go run main.go migrate up

migrate-down:
		go run main.go migrate down

up:
	docker compose up -d

down:
	docker compose down

lint:
	./bin/golangci-lint run

get-deps:
		GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

build:
	go build -o ./bin/rest_server ./main.go

docker-build:
	docker build -t cr.selcloud.ru/docker-star/test-server:v0.0.1 .

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/rest_linux ./main.go

copy-to-server:
	scp ./bin/rest_linux ./.env root@176.114.91.39:

docker-build-and-push:
		docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/docker-star/test-server:v0.0.1 .
		docker login -u token -p CRgAAAAAMUbqdgoSLDACQnjL-0PFP7crBXkYrXED cr.selcloud.ru/docker-star
		docker push cr.selcloud.ru/docker-star/test-server:v0.0.1

#docker-build-and-push:
#		docker buildx build --no-cache --platform linux/amd64 -t <REGESTRY>/test-server:v0.0.1 .
#		docker login -u <USERNAME> -p <PASSWORD> <REGESTRY>
#		docker push <REGESTRY>/test-server:v0.0.1
