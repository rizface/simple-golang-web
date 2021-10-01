FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main main.go

CMD ["/app/main"]