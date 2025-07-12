FROM golang:1.22-alphine

WORKDIR /app

RUN apk add --no-cache git mariadb-cleint tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o snippetbox ./cmd/web

EXPOSE 4000

ENV DB_DSN=web:123456@tcp(mariadb:3306)/snippetbox?parseTime=true

CMD ["./snippetbox"]