FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./

COPY static ./static
COPY templates ./templates
COPY downloads ./downloads

RUN go build -o cv .

EXPOSE 8080

CMD ["./cv"]