FROM golang:1.23.3-alpine AS builder

COPY . /afera-projects/

WORKDIR /afera-projects/

RUN go mod download

RUN go build -o ./bin/rest_server ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /afera-projects/bin/rest_server .

COPY --from=builder /afera-projects/.env .

ENTRYPOINT ["./rest_server"]

CMD ["rest"]