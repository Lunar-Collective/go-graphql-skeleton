FROM golang:1.12

WORKDIR /root/app

COPY . .


RUN go get -u github.com/cosmtrek/air

RUN go build ./cmd/server.go

EXPOSE 80

CMD go run server