FROM golang:latest

WORKDIR /app

ADD . /app/

RUN go mod tidy
RUN go mod verify

EXPOSE 3000

CMD ["go","run", "main.go"]