FROM golang:1.12.11-alpine3.9 AS builder
LABEL maintainer="Marco Ollivier <marcopollivier@gmail.com>"
WORKDIR /go/src/github.com/marcopollivier/app/
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o whoamigo ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/marcopollivier/app/whoamigo .
EXPOSE 8081
CMD ["./whoamigo"]