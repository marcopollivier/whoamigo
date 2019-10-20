# Who am I, Go? 

Simple HTTP project written in [Go](https://golang.org/) 

See other _**Who am I?**_ options
- [Who am I, Node?](https://github.com/marcopollivier/whoaminode)
- [Who am I, Java?]() -- in progress
- [Who am I, Clojure?]() -- in progress

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
