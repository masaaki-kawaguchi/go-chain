FROM golang:1.16

#ENV GO111MODULE=on

WORKDIR /go/src

COPY ../../main.go /go/src/main.go

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build .

RUN go get github.com/pilu/fresh

CMD ["go", "run", "main.go", "fresh"]