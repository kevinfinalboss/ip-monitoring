FROM golang:1.16-alpine3.13

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

HEALTHCHECK --interval=60s --timeout=5s --start-period=5s --retries=3 CMD [ "wget", "-q", "http://localhost:8080/healthcheck" ]

CMD ["/app/main"]
