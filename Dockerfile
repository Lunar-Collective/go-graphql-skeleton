FROM golang:1.12

WORKDIR /root/app

COPY . .


RUN go get github.com/cosmtrek/air

RUN go build ./cmd/server.go

EXPOSE 80

CMD ./server