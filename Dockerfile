FROM golang:alpine

WORKDIR /account_service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/run

EXPOSE 13961

CMD ["./main"]