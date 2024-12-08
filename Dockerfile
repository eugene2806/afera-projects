FROM golang:1.23.3-alpine AS builder

COPY . /afera-projects/

WORKDIR /afera-projects/

RUN go mod download

RUN go build -o ./bin/rest_server ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /afera-projects/bin/rest_server .

COPY --from=builder /afera-projects/.env .

COPY --from=builder /afera-projects/migrate ./migrate

ENTRYPOINT ["/bin/sh", "-c", "./rest_server migrate up & ./rest_server rest"]