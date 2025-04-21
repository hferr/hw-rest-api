FROM golang:1.24.1-alpine

WORKDIR /hw-rest-api

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -a -o ./bin ./cmd

CMD ["/hw-rest-api/api"]
EXPOSE 8080
