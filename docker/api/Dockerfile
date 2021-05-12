FROM golang:1.16

ENV GO111MODULE=on

WORKDIR ../../../go-chain

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]