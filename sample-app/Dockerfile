FROM golang:1.21-alpine as builder

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod tidy
COPY main.go .
RUN go build -o sample-app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/sample-app .
EXPOSE 8888
CMD ["./sample-app"]
