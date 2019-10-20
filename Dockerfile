FROM golang:latest

LABEL maintainer="Marco Ollivier <marcopollivier@gmail.com>"

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o whoamigo ./cmd/server

EXPOSE 8081

CMD ["./whoamigo"]