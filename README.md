# Who am I, Go? 

Simple HTTP project written in [Go](https://golang.org/) 

See other _**Who am I?**_ options
- [Who am I, Node?](https://github.com/marcopollivier/whoaminode)
- [Who am I, Java?]() -- in progress
- [Who am I, Clojure?]() -- in progress

## Docker

Download the Docker Image on [Dockerhub](https://hub.docker.com/r/marcopollivier/whoami) 

## API REST

### Response

- **me**: hostname (container ID for example)
- **timestamp**: execution start time
- **duration**: execution duration in microseconds 

```json
{
  "me": "RJ-MB2680.lan",
  "timestamp": "2019-02-12T10:00:00.000Z"
}
```
