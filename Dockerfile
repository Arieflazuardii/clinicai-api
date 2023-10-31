FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app.exe .

EXPOSE 8080

CMD ["/app.exe"]
