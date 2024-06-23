FROM golang:latest AS builder

WORKDIR /viia

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /src/vvia

COPY --from=builder /viia/main .

CMD ["./main"]