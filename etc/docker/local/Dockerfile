FROM golang:1.21

WORKDIR /go/src/github.com/rinchsan/ringo

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", "./etc/air/.air.toml"]
