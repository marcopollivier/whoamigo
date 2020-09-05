# Who am I, Go? 

Simple HTTP project written in [Go](https://golang.org/) 

See other _**Who am I?**_ options
- [Who am I, Node?](https://github.com/marcopollivier/whoaminode)
- [Who am I, Java?](https://github.com/marcopollivier/whoamijava)
- [Who am I, Clojure?](https://github.com/marcopollivier/whoamiclojure)

## Run locally 

### Execute the server 
```bash 
$ go run cmd/server/main.go
```

### Execute the CURL command 

```bash 
$ curl http://localhost:8081 | json_pp
```

## Docker

### Build Docker image

```bash
$ docker build -t marcopollivier/whoami:go-latest .
```

or [Download the Docker Image](https://hub.docker.com/r/marcopollivier/whoami)

### Push Docker Image

```bash
$ docker push marcopollivier/whoami:go-latest
```

### Run Docker container

```bash
$ docker run --name whoami_go -d -p 8081:8081 marcopollivier/whoami:go-latest
```

## API REST

### Response

- **me**: hostname (container ID for example)
- **timestamp**: execution start time 

```json
{
  "me": "RJ-MB2680.lan",
  "timestamp": "2019-02-12T10:00:00.000Z"
}
```
