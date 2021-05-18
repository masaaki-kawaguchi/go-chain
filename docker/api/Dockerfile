FROM golang:1.16

ENV GO111MODULE=on

WORKDIR go/src/github.com/masaaki-kawaguchi/go-chain

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build .

RUN go get github.com/pilu/fresh

RUN go run ./main.go

EXPOSE 8080

CMD ["fresh"]